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
	splitted_command := strings.Split(e.Command," ")
	if splitted_command[0] == "sudo" {
		splitted_command = splitted_command[1:]
	}
	cmd := exec.Command(splitted_command[0],splitted_command[1:]...)
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

