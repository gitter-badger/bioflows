package process

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
)

type CommandExecutor struct{
	Command        string
	CommandDir string
	InitialCommand string
	PreCommandArgs []string
	buffer         *bytes.Buffer
	errorBuff      *bytes.Buffer
}

func (e *CommandExecutor) Init() {
	e.InitialCommand = "bash"
	e.PreCommandArgs = []string{"-c"}
}

func (e *CommandExecutor) Run() (int, error) {
	e.buffer = &bytes.Buffer{}
	e.errorBuff = &bytes.Buffer{}
	splittedCommand := strings.Split(e.Command," ")
	cmd := exec.Command(e.InitialCommand, strings.Join(e.PreCommandArgs," "),strings.Join(splittedCommand," "))
	cmd.Dir = e.CommandDir
	cmd.Stdout = e.buffer
	cmd.Stderr = e.errorBuff
	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exiterr , ok := err.(*exec.ExitError) ; ok {
			if status , ok := exiterr.Sys().(syscall.WaitStatus); ok {
				exitCode = status.ExitStatus()
			}
		}
	}

	return exitCode , err
}

func (e CommandExecutor) GetOutput() *bytes.Buffer {
	return e.buffer
}

func (e CommandExecutor) GetError() *bytes.Buffer {
	return e.errorBuff
}


