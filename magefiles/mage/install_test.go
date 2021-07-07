package mage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterOK(t *testing.T) {
	assert.True(t, filter("Linux-64bit.tar.gz"))
}

func TestFilterNotOk(t *testing.T) {
	assert.False(t, filter("Anythin, really"))
}
