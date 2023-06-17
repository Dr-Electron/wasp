// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package state

import (
	"errors"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/iotaledger/hive.go/kvstore"
	"github.com/iotaledger/wasp/packages/kv/buffered"
	"github.com/iotaledger/wasp/packages/metrics"
	"github.com/iotaledger/wasp/packages/trie"
)

// store is the implementation of the Store interface
type store struct {
	// db is the backing key-value store
	db *storeDB

	// stateCache is a cache of immutable state readers by trie root. Reusing the
	// State instances allows to better take advantage of its internal caches.
	stateCache *lru.Cache[trie.Hash, *state]

	metrics metrics.IStateMetrics
}

func NewStore(db kvstore.KVStore) Store {
	return NewStoreWithMetrics(db, metrics.NewEmptyStateMetrics())
}

func NewStoreWithMetrics(db kvstore.KVStore, metrics metrics.IStateMetrics) Store {
	stateCache, err := lru.New[trie.Hash, *state](100)
	if err != nil {
		panic(err)
	}
	return &store{
		db:         &storeDB{db},
		stateCache: stateCache,
		metrics:    metrics,
	}
}

func (s *store) blockByTrieRoot(root trie.Hash) (Block, error) {
	return s.db.readBlock(root)
}

func (s *store) HasTrieRoot(root trie.Hash) bool {
	return s.db.hasBlock(root)
}

func (s *store) BlockByTrieRoot(root trie.Hash) (Block, error) {
	return s.blockByTrieRoot(root)
}

func (s *store) stateByTrieRoot(root trie.Hash) (*state, error) {
	if r, ok := s.stateCache.Get(root); ok {
		return r, nil
	}
	r, err := newState(s.db, root)
	if err != nil {
		return nil, err
	}
	s.stateCache.Add(root, r)
	return r, nil
}

func (s *store) StateByTrieRoot(root trie.Hash) (State, error) {
	return s.stateByTrieRoot(root)
}

func (s *store) NewOriginStateDraft() StateDraft {
	return newOriginStateDraft()
}

func (s *store) NewStateDraft(timestamp time.Time, prevL1Commitment *L1Commitment) (StateDraft, error) {
	prevState, err := s.stateByTrieRoot(prevL1Commitment.TrieRoot())
	if err != nil {
		return nil, err
	}
	return newStateDraft(timestamp, prevL1Commitment, prevState), nil
}

func (s *store) NewEmptyStateDraft(prevL1Commitment *L1Commitment) (StateDraft, error) {
	if prevL1Commitment == nil {
		return nil, errors.New("nil prevL1Commitment")
	}
	prevState, err := s.stateByTrieRoot(prevL1Commitment.TrieRoot())
	if err != nil {
		return nil, err
	}
	return newEmptyStateDraft(prevL1Commitment, prevState), nil
}

func (s *store) extractBlock(d StateDraft) (Block, *buffered.Mutations, trie.CommitStats) {
	buf, bufDB := s.db.buffered()

	var baseTrieRoot trie.Hash
	{
		baseL1Commitment := d.BaseL1Commitment()
		if baseL1Commitment != nil {
			if !s.db.hasBlock(baseL1Commitment.TrieRoot()) {
				panic("cannot commit state: base trie root not found")
			}
			baseTrieRoot = baseL1Commitment.TrieRoot()
		} else {
			baseTrieRoot = bufDB.initTrie()
		}
	}

	// compute state db mutations
	block, stats := func() (Block, trie.CommitStats) {
		trie, err := bufDB.trieUpdatable(baseTrieRoot)
		if err != nil {
			// should not happen
			panic(err)
		}
		for k, v := range d.Mutations().Sets {
			trie.Update([]byte(k), v)
		}
		for k := range d.Mutations().Dels {
			trie.Delete([]byte(k))
		}
		trieRoot, stats := trie.Commit(bufDB.trieStore())
		block := &block{
			trieRoot:             trieRoot,
			mutations:            d.Mutations(),
			previousL1Commitment: d.BaseL1Commitment(),
		}
		bufDB.saveBlock(block)
		return block, stats
	}()

	return block, buf.muts, stats
}

func (s *store) ExtractBlock(d StateDraft) Block {
	block, _, _ := s.extractBlock(d)
	return block
}

func (s *store) Commit(d StateDraft) Block {
	start := time.Now()
	block, muts, stats := s.extractBlock(d)
	s.db.commitToDB(muts)
	s.metrics.BlockCommitted(time.Since(start), stats.CreatedNodes, stats.CreatedValues)
	return block
}

func (s *store) Prune(trieRoot trie.Hash) (trie.PruneStats, error) {
	start := time.Now()
	buf, bufDB := s.db.buffered()
	stats, err := trie.Prune(bufDB.trieStore(), trieRoot)
	if err != nil {
		return trie.PruneStats{}, err
	}
	s.db.pruneBlock(trieRoot)
	s.db.commitToDB(buf.muts)
	s.stateCache.Remove(trieRoot)
	s.metrics.BlockPruned(time.Since(start), stats.DeletedNodes, stats.DeletedValues)
	return stats, nil
}

func (s *store) SetLatest(trieRoot trie.Hash) error {
	_, err := s.BlockByTrieRoot(trieRoot)
	if err != nil {
		return err
	}
	s.db.setLatestTrieRoot(trieRoot)
	return nil
}

func (s *store) LatestBlock() (Block, error) {
	root, err := s.db.latestTrieRoot()
	if err != nil {
		return nil, err
	}
	return s.BlockByTrieRoot(root)
}

func (s *store) LatestBlockIndex() (uint32, error) {
	latestTrieRoot, err := s.LatestTrieRoot()
	if err != nil {
		return 0, err
	}
	state, err := s.StateByTrieRoot(latestTrieRoot)
	if err != nil {
		return 0, err
	}
	return state.BlockIndex(), nil
}

func (s *store) LatestState() (State, error) {
	root, err := s.db.latestTrieRoot()
	if err != nil {
		return nil, err
	}
	return s.StateByTrieRoot(root)
}

func (s *store) LatestTrieRoot() (trie.Hash, error) {
	return s.db.latestTrieRoot()
}
