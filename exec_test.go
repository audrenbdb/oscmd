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
		run           cmdRun

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
				return errors.New("failed to cmdRun")
			},
			cmdName: "ls",
			outErr: errors.New("failed to cmdRun"),

		},
	}

	for _, test := range tests {
		exec := newExecFunc(test.cmdIsNotFound, test.run)
		err := exec(test.cmdName, test.args...)
		assert.Equal(t, test.outErr, err)
	}
}

func TestStartFunc(t *testing.T) {
	tests := []struct{
		description string

		cmdIsNotFound cmdIsNotFound
		run           cmdRun

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
				return errors.New("failed to cmdRun")
			},
			cmdName: "ls",
			outErr: errors.New("failed to cmdRun"),

		},
	}

	for _, test := range tests {
		exec := newStartFunc(test.cmdIsNotFound, test.run)
		err := exec(test.cmdName, test.args...)
		assert.Equal(t, test.outErr, err)
	}
}