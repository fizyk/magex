package command

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

// PipeCommands runs first command and pipes it's output to second one
func PipeCommands(mainCommand *exec.Cmd, pipeToCommand *exec.Cmd) (stdout bytes.Buffer, stderr bytes.Buffer) {
	fmt.Printf("%s | %s \n", mainCommand.String(), pipeToCommand.String())
	read, write := io.Pipe()
	// Main command will write
	mainCommand.Stdout = write

	// pipeCommand will read
	pipeToCommand.Stdin = read
	pipeToCommand.Stdout = &stdout
	pipeToCommand.Stderr = &stderr
	mainCommand.Start()
	pipeToCommand.Start()
	// Wait for the mainCommand to finish
	mainCommand.Wait()
	// Close write io
	write.Close()
	// Wait for the pipeCommand to finish
	pipeToCommand.Wait()
	return
}
