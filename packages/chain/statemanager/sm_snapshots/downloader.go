package sm_snapshots

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type downloader struct {
	ctx         context.Context
	chunkReader io.ReadCloser
	filePath    string
	fileSize    int
	chunkEnd    int
	chunkSize   int
}

var (
	_ io.Reader     = &downloader{}
	_ io.Closer     = &downloader{}
	_ io.ReadCloser = &downloader{}
)

const defaultChunkSizeConst = 1024

func NewDownloader(
	ctx context.Context,
	filePath string,
	chunkSize ...int,
) (io.ReadCloser, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodHead, filePath, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to make head request to %s: %w", filePath, err)
	}
	head, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to receive header for url %s: %w", filePath, err)
	}
	defer head.Body.Close()

	acceptRanges := head.Header.Get("Accept-Ranges")
	fileSizeStr := head.Header.Get("Content-Length")
	fileSize, err := strconv.Atoi(fileSizeStr)
	result := &downloader{
		ctx:      ctx,
		filePath: filePath,
		fileSize: fileSize,
	}
	if err != nil {
		return nil, fmt.Errorf("failed to convert file length %v to integer: %w", fileSizeStr, err)
	}
	if acceptRanges == "" || strings.ToLower(acceptRanges) == "none" {
		result.chunkSize = 0
	} else {
		if len(chunkSize) > 0 {
			result.chunkSize = chunkSize[0]
		} else {
			result.chunkSize = defaultChunkSizeConst
		}
		result.chunkEnd = 0
	}
	err = result.setReader()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *downloader) setReader() error {
	request, err := http.NewRequestWithContext(d.ctx, http.MethodGet, d.filePath, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to make get request to %s: %w", d.filePath, err)
	}
	chunkPartStr := ""
	if d.chunkSize > 0 {
		start := d.chunkEnd
		end := start + d.chunkSize
		if end > d.fileSize {
			end = d.fileSize
		}
		request.Header.Add("Range", "bytes="+strconv.Itoa(start)+"-"+strconv.Itoa(end-1))
		chunkPartStr = fmt.Sprintf(" byte %v to %v", start, end)
		d.chunkEnd = end
	} else {
		d.chunkEnd = d.fileSize
	}
	chunk, err := http.DefaultClient.Do(request) //nolint:bodyclose// closing is handled differently; linter cannot understand that
	if err != nil {
		return fmt.Errorf("failed to get file%s from %s: %w", chunkPartStr, d.filePath, err)
	}
	d.chunkReader = chunk.Body
	return nil
}

func (d *downloader) Read(b []byte) (int, error) {
	n, err := d.chunkReader.Read(b)
	if err == io.EOF {
		if d.chunkEnd >= d.fileSize {
			return n, err
		}
		d.chunkReader.Close()
		err = d.setReader()
		if err != nil {
			return n, err
		}
		var nn int
		nn, err = d.Read(b[n:])
		return n + nn, err
	}
	return n, err
}

func (d *downloader) Close() error {
	return d.chunkReader.Close()
}
