package oscmd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecFunc(t *testing.T) {
	tests := []struct{
		description string

		cmdIsNotFound cmdIsNotFound
		run run

		cmdName string
		args []string
		outErr error
	}{
		{
			description: "should return cmd not found error",
			cmdIsNotFound: func(cmdName string) bool {return true},
			cmdName: "ls",
			outErr: cmdNotFound,
		},
		{
			description: "should fail to exec cmd",
			cmdIsNotFound: func(cmdName string) bool {return false},
			run: func(cmdName string, args ...string) error {
				return errors.New("failed to run")
			},
			cmdName: "ls",
			outErr: errors.New("failed to run"),

		},
	}

	for _, test := range tests {
		exec := newExecFunc(test.cmdIsNotFound, test.run)
		err := exec(test.cmdName, test.args...)
		assert.Equal(t, test.outErr, err)
	}
}

/*
func TestCommandExistsInLinux(t *testing.T) {
	t.Skip()
	tests := []struct {
		cmd string
		args []string
		err error
	}{
		{
			cmd: "toto",
			err: errors.New("command not found in path"),
		},
		{
			cmd: "ls",
			err: nil,
		},
	}

	for _, test := range tests {
		execFunc := NewExecFunc()
		err := execFunc(test.cmd)
		if test.err == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.err.Error(), err.Error())
		}
	}
}
 */