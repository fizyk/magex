package file

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type fileSuite struct {
	suite.Suite
}

func TestFileSuite(t *testing.T) {
	suite.Run(t, new(fileSuite))
}

func (s *fileSuite) TestExists() {
	s.True(Exists("exists.go"))
}

func (s *fileSuite) TestDoesNotExists() {
	s.False(Exists("does_not_exists.go"))
}
