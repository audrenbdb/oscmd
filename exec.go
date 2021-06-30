//oscmd provides a wrapper around os commands to ease unit testing
package oscmd

import "os/exec"

//Exec creates and run a given command and its args
type Exec = func(name string, args ...string) error


func NewExecFunc() Exec {
	lookPath := exec.LookPath
	run := newRunFunc()
	return newExecFunc(newCmdIsNotFoundFunc(lookPath), run)
}


func newExecFunc(cmdIsNotFound cmdIsNotFound, run run) Exec {
	return func(cmdName string, args ...string) error {
		if cmdIsNotFound(cmdName) {
			return cmdNotFound
		}
		return run(cmdName, args...)
	}
}

type lookPath = func(file string) (string, error)
type cmdIsNotFound = func(cmdName string) bool
type run = func(cmdName string, args ...string) error

func newRunFunc() run {
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

