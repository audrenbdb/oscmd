//oscmd provides a wrapper around os commands to ease unit testing
package oscmd

import (
	"os/exec"
)

//Exec creates and run a given command and its args
type Exec = func(name string, args ...string) error

func NewExecFunc() Exec {
	return func(cmdName string, args ...string) error {
		if isNotFound(cmdName) {
			return notFound
		}
		cmd := exec.Command(cmdName, args...)
		return cmd.Run()
	}
}


func isNotFound(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err != nil
}

//error handling
type cmdErr string
func (e cmdErr) Error() string {return string(e)}
const notFound = cmdErr("command not found in path")

