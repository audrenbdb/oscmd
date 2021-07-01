//Package oscmd provides a wrapper around os commands to ease unit testing
package oscmd

import "os/exec"

//Run runs a command and wait for its execution end
type Run = func(name string, args ...string) error

//Start starts the command and does not wait for it to end
type Start = func(name string, args ...string) error


func NewRunFunc() Run {
	lookPath := exec.LookPath
	run := newCmdRunFunc()
	return newExecFunc(newCmdIsNotFoundFunc(lookPath), run)
}

func NewStartFunc() Start {
	lookPath := exec.LookPath
	start := newCmdStartFunc()
	return newStartFunc(newCmdIsNotFoundFunc(lookPath), start)
}

func newStartFunc(cmdIsNotFound cmdIsNotFound, start cmdStart) Start {
	return func(cmdName string, args ...string) error {
		if cmdIsNotFound(cmdName) {return cmdNotFound}
		return start(cmdName, args...)
	}
}



func newExecFunc(cmdIsNotFound cmdIsNotFound, run cmdRun) Run {
	return func(cmdName string, args ...string) error {
		if cmdIsNotFound(cmdName) {
			return cmdNotFound
		}
		return run(cmdName, args...)
	}
}

type lookPath = func(file string) (string, error)
type cmdIsNotFound = func(cmdName string) bool
type cmdRun = func(cmdName string, args ...string) error
type cmdStart = func(cmdName string, args ...string) error

func newCmdStartFunc() cmdStart {
	return func(cmdName string, args ...string) error {
		cmd := exec.Command(cmdName, args...)
		return cmd.Start()
	}
}

func newCmdRunFunc() cmdRun {
	return func(cmdName string, args ...string) error {
		cmd := exec.Command(cmdName, args...)
		return cmd.Run()
	}
}

func newCmdIsNotFoundFunc(lookPath lookPath) cmdIsNotFound {
	return func(cmdName string) bool {
		_, err := lookPath(cmdName)
		return err != nil
	}
}

