package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalBytes(t *testing.T) {
	counter := &WriteCounter{}
	aString := "A string with 22 bytes"
	byteLength, err := counter.Write([]byte(aString))
	assert.NoError(t, err)
	assert.Equal(t, 22, byteLength)
	assert.Equal(t, uint64(byteLength), counter.Total)

	otherString := "other 14 bytes"
	moreByteLength, err := counter.Write([]byte(otherString))
	assert.NoError(t, err)
	assert.Equal(t, 14, moreByteLength)
	assert.Equal(t, uint64(14+22), counter.Total)
}
