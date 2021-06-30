package command

import (
	"github.com/stretchr/testify/suite"
	"os/exec"
	"testing"
)

type pipeSuite struct {
	suite.Suite
}

func TestPipeSuite(t *testing.T) {
	suite.Run(t, new(pipeSuite))
}

func (s *pipeSuite) Test_Stdout() {
	pipeFromCommand := exec.Command("ls ", "-l")
	pipeToCommand := exec.Command("wc", "-l")
	stdout, stderr := PipeCommands(pipeFromCommand, pipeToCommand)
	s.Empty(stderr.String())
	s.Equal("0\n", stdout.String())
}
