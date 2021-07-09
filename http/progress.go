package http

import (
	"github.com/cheggaaa/pb/v3"
	"time"
)

const RefreshRate = time.Millisecond * 100

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer
// interface and we can pass this into io.TeeReader() which will report progress on each
// write cycle.
type WriteCounter struct {
	Total int // bytes read so far
	bar   *pb.ProgressBar
}

func NewWriteCounter(total int) *WriteCounter {
	bar := pb.New(total)
	bar.SetRefreshRate(RefreshRate)
	bar.Set(pb.Bytes, true)

	return &WriteCounter{
		bar: bar,
	}
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	current := len(p)
	wc.Total += current
	wc.bar.Write()
	wc.bar.Add(current)
	return current, nil
}

func (wc *WriteCounter) Start() {
	wc.bar.Start()
}

func (wc *WriteCounter) Finish() {
	wc.bar.Finish()
}
