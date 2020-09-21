package process


import (
	"bytes"
	"fmt"
	"strings"
	"os/exec"
)

type CommandExecutor struct{
	Command        string
	InitialCommand string
	PreCommandArgs []string
	buffer         *bytes.Buffer
	errorBuff      *bytes.Buffer
}

func (e *CommandExecutor) Init() {
	e.InitialCommand = "bash"
	e.PreCommandArgs = []string{"-c"}
}

func (e *CommandExecutor) Run() error {
	e.buffer = &bytes.Buffer{}
	e.errorBuff = &bytes.Buffer{}
	splittedCommand := strings.Split(e.Command," ")
	cmd := exec.Command(e.InitialCommand, strings.Join(e.PreCommandArgs," "),strings.Join(splittedCommand," "))
	fmt.Println(cmd.String())
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


