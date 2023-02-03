/*
Wasp API

REST API for the Wasp node

API version: 0.4.0-alpha.8-14-gbb23763d3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
	"time"
)

// checks if the ConsensusWorkflowMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsensusWorkflowMetrics{}

// ConsensusWorkflowMetrics struct for ConsensusWorkflowMetrics
type ConsensusWorkflowMetrics struct {
	// Shows current state index of the consensus
	CurrentStateIndex *uint32 `json:"currentStateIndex,omitempty"`
	// Shows if batch proposal is sent out in current consensus iteration
	FlagBatchProposalSent bool `json:"flagBatchProposalSent"`
	// Shows if consensus on batch is reached and known in current consensus iteration
	FlagConsensusBatchKnown bool `json:"flagConsensusBatchKnown"`
	// Shows if consensus algorithm is still not completed in current consensus iteration
	FlagInProgress bool `json:"flagInProgress"`
	// Shows if state output is received in current consensus iteration
	FlagStateReceived bool `json:"flagStateReceived"`
	// Shows if consensus on transaction is reached in current consensus iteration
	FlagTransactionFinalized bool `json:"flagTransactionFinalized"`
	// Shows if transaction is posted to L1 in current consensus iteration
	FlagTransactionPosted bool `json:"flagTransactionPosted"`
	// Shows if L1 reported that it has seen the transaction of current consensus iteration
	FlagTransactionSeen bool `json:"flagTransactionSeen"`
	// Shows if virtual machine has returned its results in current consensus iteration
	FlagVMResultSigned bool `json:"flagVMResultSigned"`
	// Shows if virtual machine is started in current consensus iteration
	FlagVMStarted bool `json:"flagVMStarted"`
	// Shows when batch proposal was last sent out in current consensus iteration
	TimeBatchProposalSent time.Time `json:"timeBatchProposalSent"`
	// Shows when algorithm was last completed in current consensus iteration
	TimeCompleted time.Time `json:"timeCompleted"`
	// Shows when ACS results of consensus on batch was last received in current consensus iteration
	TimeConsensusBatchKnown time.Time `json:"timeConsensusBatchKnown"`
	// Shows when algorithm last noted that all the data for consensus on transaction had been received in current consensus iteration
	TimeTransactionFinalized time.Time `json:"timeTransactionFinalized"`
	// Shows when transaction was last posted to L1 in current consensus iteration
	TimeTransactionPosted time.Time `json:"timeTransactionPosted"`
	// Shows when algorithm last noted that transaction had been seen by L1 in current consensus iteration
	TimeTransactionSeen time.Time `json:"timeTransactionSeen"`
	// Shows when virtual machine results were last received and signed in current consensus iteration
	TimeVMResultSigned time.Time `json:"timeVMResultSigned"`
	// Shows when virtual machine was last started in current consensus iteration
	TimeVMStarted time.Time `json:"timeVMStarted"`
}

// NewConsensusWorkflowMetrics instantiates a new ConsensusWorkflowMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsensusWorkflowMetrics(flagBatchProposalSent bool, flagConsensusBatchKnown bool, flagInProgress bool, flagStateReceived bool, flagTransactionFinalized bool, flagTransactionPosted bool, flagTransactionSeen bool, flagVMResultSigned bool, flagVMStarted bool, timeBatchProposalSent time.Time, timeCompleted time.Time, timeConsensusBatchKnown time.Time, timeTransactionFinalized time.Time, timeTransactionPosted time.Time, timeTransactionSeen time.Time, timeVMResultSigned time.Time, timeVMStarted time.Time) *ConsensusWorkflowMetrics {
	this := ConsensusWorkflowMetrics{}
	this.FlagBatchProposalSent = flagBatchProposalSent
	this.FlagConsensusBatchKnown = flagConsensusBatchKnown
	this.FlagInProgress = flagInProgress
	this.FlagStateReceived = flagStateReceived
	this.FlagTransactionFinalized = flagTransactionFinalized
	this.FlagTransactionPosted = flagTransactionPosted
	this.FlagTransactionSeen = flagTransactionSeen
	this.FlagVMResultSigned = flagVMResultSigned
	this.FlagVMStarted = flagVMStarted
	this.TimeBatchProposalSent = timeBatchProposalSent
	this.TimeCompleted = timeCompleted
	this.TimeConsensusBatchKnown = timeConsensusBatchKnown
	this.TimeTransactionFinalized = timeTransactionFinalized
	this.TimeTransactionPosted = timeTransactionPosted
	this.TimeTransactionSeen = timeTransactionSeen
	this.TimeVMResultSigned = timeVMResultSigned
	this.TimeVMStarted = timeVMStarted
	return &this
}

// NewConsensusWorkflowMetricsWithDefaults instantiates a new ConsensusWorkflowMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsensusWorkflowMetricsWithDefaults() *ConsensusWorkflowMetrics {
	this := ConsensusWorkflowMetrics{}
	return &this
}

// GetCurrentStateIndex returns the CurrentStateIndex field value if set, zero value otherwise.
func (o *ConsensusWorkflowMetrics) GetCurrentStateIndex() uint32 {
	if o == nil || isNil(o.CurrentStateIndex) {
		var ret uint32
		return ret
	}
	return *o.CurrentStateIndex
}

// GetCurrentStateIndexOk returns a tuple with the CurrentStateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetCurrentStateIndexOk() (*uint32, bool) {
	if o == nil || isNil(o.CurrentStateIndex) {
		return nil, false
	}
	return o.CurrentStateIndex, true
}

// HasCurrentStateIndex returns a boolean if a field has been set.
func (o *ConsensusWorkflowMetrics) HasCurrentStateIndex() bool {
	if o != nil && !isNil(o.CurrentStateIndex) {
		return true
	}

	return false
}

// SetCurrentStateIndex gets a reference to the given uint32 and assigns it to the CurrentStateIndex field.
func (o *ConsensusWorkflowMetrics) SetCurrentStateIndex(v uint32) {
	o.CurrentStateIndex = &v
}

// GetFlagBatchProposalSent returns the FlagBatchProposalSent field value
func (o *ConsensusWorkflowMetrics) GetFlagBatchProposalSent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagBatchProposalSent
}

// GetFlagBatchProposalSentOk returns a tuple with the FlagBatchProposalSent field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagBatchProposalSentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagBatchProposalSent, true
}

// SetFlagBatchProposalSent sets field value
func (o *ConsensusWorkflowMetrics) SetFlagBatchProposalSent(v bool) {
	o.FlagBatchProposalSent = v
}

// GetFlagConsensusBatchKnown returns the FlagConsensusBatchKnown field value
func (o *ConsensusWorkflowMetrics) GetFlagConsensusBatchKnown() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagConsensusBatchKnown
}

// GetFlagConsensusBatchKnownOk returns a tuple with the FlagConsensusBatchKnown field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagConsensusBatchKnownOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagConsensusBatchKnown, true
}

// SetFlagConsensusBatchKnown sets field value
func (o *ConsensusWorkflowMetrics) SetFlagConsensusBatchKnown(v bool) {
	o.FlagConsensusBatchKnown = v
}

// GetFlagInProgress returns the FlagInProgress field value
func (o *ConsensusWorkflowMetrics) GetFlagInProgress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagInProgress
}

// GetFlagInProgressOk returns a tuple with the FlagInProgress field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagInProgressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagInProgress, true
}

// SetFlagInProgress sets field value
func (o *ConsensusWorkflowMetrics) SetFlagInProgress(v bool) {
	o.FlagInProgress = v
}

// GetFlagStateReceived returns the FlagStateReceived field value
func (o *ConsensusWorkflowMetrics) GetFlagStateReceived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagStateReceived
}

// GetFlagStateReceivedOk returns a tuple with the FlagStateReceived field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagStateReceivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagStateReceived, true
}

// SetFlagStateReceived sets field value
func (o *ConsensusWorkflowMetrics) SetFlagStateReceived(v bool) {
	o.FlagStateReceived = v
}

// GetFlagTransactionFinalized returns the FlagTransactionFinalized field value
func (o *ConsensusWorkflowMetrics) GetFlagTransactionFinalized() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagTransactionFinalized
}

// GetFlagTransactionFinalizedOk returns a tuple with the FlagTransactionFinalized field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagTransactionFinalizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagTransactionFinalized, true
}

// SetFlagTransactionFinalized sets field value
func (o *ConsensusWorkflowMetrics) SetFlagTransactionFinalized(v bool) {
	o.FlagTransactionFinalized = v
}

// GetFlagTransactionPosted returns the FlagTransactionPosted field value
func (o *ConsensusWorkflowMetrics) GetFlagTransactionPosted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagTransactionPosted
}

// GetFlagTransactionPostedOk returns a tuple with the FlagTransactionPosted field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagTransactionPostedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagTransactionPosted, true
}

// SetFlagTransactionPosted sets field value
func (o *ConsensusWorkflowMetrics) SetFlagTransactionPosted(v bool) {
	o.FlagTransactionPosted = v
}

// GetFlagTransactionSeen returns the FlagTransactionSeen field value
func (o *ConsensusWorkflowMetrics) GetFlagTransactionSeen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagTransactionSeen
}

// GetFlagTransactionSeenOk returns a tuple with the FlagTransactionSeen field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagTransactionSeenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagTransactionSeen, true
}

// SetFlagTransactionSeen sets field value
func (o *ConsensusWorkflowMetrics) SetFlagTransactionSeen(v bool) {
	o.FlagTransactionSeen = v
}

// GetFlagVMResultSigned returns the FlagVMResultSigned field value
func (o *ConsensusWorkflowMetrics) GetFlagVMResultSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagVMResultSigned
}

// GetFlagVMResultSignedOk returns a tuple with the FlagVMResultSigned field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagVMResultSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagVMResultSigned, true
}

// SetFlagVMResultSigned sets field value
func (o *ConsensusWorkflowMetrics) SetFlagVMResultSigned(v bool) {
	o.FlagVMResultSigned = v
}

// GetFlagVMStarted returns the FlagVMStarted field value
func (o *ConsensusWorkflowMetrics) GetFlagVMStarted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagVMStarted
}

// GetFlagVMStartedOk returns a tuple with the FlagVMStarted field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetFlagVMStartedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagVMStarted, true
}

// SetFlagVMStarted sets field value
func (o *ConsensusWorkflowMetrics) SetFlagVMStarted(v bool) {
	o.FlagVMStarted = v
}

// GetTimeBatchProposalSent returns the TimeBatchProposalSent field value
func (o *ConsensusWorkflowMetrics) GetTimeBatchProposalSent() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeBatchProposalSent
}

// GetTimeBatchProposalSentOk returns a tuple with the TimeBatchProposalSent field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetTimeBatchProposalSentOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeBatchProposalSent, true
}

// SetTimeBatchProposalSent sets field value
func (o *ConsensusWorkflowMetrics) SetTimeBatchProposalSent(v time.Time) {
	o.TimeBatchProposalSent = v
}

// GetTimeCompleted returns the TimeCompleted field value
func (o *ConsensusWorkflowMetrics) GetTimeCompleted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeCompleted
}

// GetTimeCompletedOk returns a tuple with the TimeCompleted field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetTimeCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeCompleted, true
}

// SetTimeCompleted sets field value
func (o *ConsensusWorkflowMetrics) SetTimeCompleted(v time.Time) {
	o.TimeCompleted = v
}

// GetTimeConsensusBatchKnown returns the TimeConsensusBatchKnown field value
func (o *ConsensusWorkflowMetrics) GetTimeConsensusBatchKnown() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeConsensusBatchKnown
}

// GetTimeConsensusBatchKnownOk returns a tuple with the TimeConsensusBatchKnown field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetTimeConsensusBatchKnownOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeConsensusBatchKnown, true
}

// SetTimeConsensusBatchKnown sets field value
func (o *ConsensusWorkflowMetrics) SetTimeConsensusBatchKnown(v time.Time) {
	o.TimeConsensusBatchKnown = v
}

// GetTimeTransactionFinalized returns the TimeTransactionFinalized field value
func (o *ConsensusWorkflowMetrics) GetTimeTransactionFinalized() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeTransactionFinalized
}

// GetTimeTransactionFinalizedOk returns a tuple with the TimeTransactionFinalized field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetTimeTransactionFinalizedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeTransactionFinalized, true
}

// SetTimeTransactionFinalized sets field value
func (o *ConsensusWorkflowMetrics) SetTimeTransactionFinalized(v time.Time) {
	o.TimeTransactionFinalized = v
}

// GetTimeTransactionPosted returns the TimeTransactionPosted field value
func (o *ConsensusWorkflowMetrics) GetTimeTransactionPosted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeTransactionPosted
}

// GetTimeTransactionPostedOk returns a tuple with the TimeTransactionPosted field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetTimeTransactionPostedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeTransactionPosted, true
}

// SetTimeTransactionPosted sets field value
func (o *ConsensusWorkflowMetrics) SetTimeTransactionPosted(v time.Time) {
	o.TimeTransactionPosted = v
}

// GetTimeTransactionSeen returns the TimeTransactionSeen field value
func (o *ConsensusWorkflowMetrics) GetTimeTransactionSeen() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeTransactionSeen
}

// GetTimeTransactionSeenOk returns a tuple with the TimeTransactionSeen field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetTimeTransactionSeenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeTransactionSeen, true
}

// SetTimeTransactionSeen sets field value
func (o *ConsensusWorkflowMetrics) SetTimeTransactionSeen(v time.Time) {
	o.TimeTransactionSeen = v
}

// GetTimeVMResultSigned returns the TimeVMResultSigned field value
func (o *ConsensusWorkflowMetrics) GetTimeVMResultSigned() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeVMResultSigned
}

// GetTimeVMResultSignedOk returns a tuple with the TimeVMResultSigned field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetTimeVMResultSignedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeVMResultSigned, true
}

// SetTimeVMResultSigned sets field value
func (o *ConsensusWorkflowMetrics) SetTimeVMResultSigned(v time.Time) {
	o.TimeVMResultSigned = v
}

// GetTimeVMStarted returns the TimeVMStarted field value
func (o *ConsensusWorkflowMetrics) GetTimeVMStarted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeVMStarted
}

// GetTimeVMStartedOk returns a tuple with the TimeVMStarted field value
// and a boolean to check if the value has been set.
func (o *ConsensusWorkflowMetrics) GetTimeVMStartedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeVMStarted, true
}

// SetTimeVMStarted sets field value
func (o *ConsensusWorkflowMetrics) SetTimeVMStarted(v time.Time) {
	o.TimeVMStarted = v
}

func (o ConsensusWorkflowMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsensusWorkflowMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CurrentStateIndex) {
		toSerialize["currentStateIndex"] = o.CurrentStateIndex
	}
	toSerialize["flagBatchProposalSent"] = o.FlagBatchProposalSent
	toSerialize["flagConsensusBatchKnown"] = o.FlagConsensusBatchKnown
	toSerialize["flagInProgress"] = o.FlagInProgress
	toSerialize["flagStateReceived"] = o.FlagStateReceived
	toSerialize["flagTransactionFinalized"] = o.FlagTransactionFinalized
	toSerialize["flagTransactionPosted"] = o.FlagTransactionPosted
	toSerialize["flagTransactionSeen"] = o.FlagTransactionSeen
	toSerialize["flagVMResultSigned"] = o.FlagVMResultSigned
	toSerialize["flagVMStarted"] = o.FlagVMStarted
	toSerialize["timeBatchProposalSent"] = o.TimeBatchProposalSent
	toSerialize["timeCompleted"] = o.TimeCompleted
	toSerialize["timeConsensusBatchKnown"] = o.TimeConsensusBatchKnown
	toSerialize["timeTransactionFinalized"] = o.TimeTransactionFinalized
	toSerialize["timeTransactionPosted"] = o.TimeTransactionPosted
	toSerialize["timeTransactionSeen"] = o.TimeTransactionSeen
	toSerialize["timeVMResultSigned"] = o.TimeVMResultSigned
	toSerialize["timeVMStarted"] = o.TimeVMStarted
	return toSerialize, nil
}

type NullableConsensusWorkflowMetrics struct {
	value *ConsensusWorkflowMetrics
	isSet bool
}

func (v NullableConsensusWorkflowMetrics) Get() *ConsensusWorkflowMetrics {
	return v.value
}

func (v *NullableConsensusWorkflowMetrics) Set(val *ConsensusWorkflowMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableConsensusWorkflowMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableConsensusWorkflowMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsensusWorkflowMetrics(val *ConsensusWorkflowMetrics) *NullableConsensusWorkflowMetrics {
	return &NullableConsensusWorkflowMetrics{value: val, isSet: true}
}

func (v NullableConsensusWorkflowMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsensusWorkflowMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


