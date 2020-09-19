package process


import (
	"bytes"
	"strings"
	"os/exec"
)

type CommandExecutor struct{
	Command string
	buffer *bytes.Buffer
	errorBuff *bytes.Buffer
}

func (e *CommandExecutor) Run() error {
	e.buffer = &bytes.Buffer{}
	e.errorBuff = &bytes.Buffer{}
	splittedCommand := strings.Split(e.Command," ")
	if splittedCommand[0] == "sudo" {
		splittedCommand = splittedCommand[1:]
	}
	cmd := exec.Command(splittedCommand[0], splittedCommand[1:]...)
	cmd.Stdout = e.buffer
	cmd.Stderr = e.errorBuff
	return cmd.Run()
}

func (e CommandExecutor) GetOutput() *bytes.Buffer {
	return e.buffer
}

func (e CommandExecutor) GetError() *bytes.Buffer {
	return e.errorBuff
}


