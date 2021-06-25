package oscmd_test

import (
	"errors"
	"github.com/audrenbdb/oscmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandExistsInLinux(t *testing.T) {
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
		execFunc := oscmd.NewExecFunc()
		err := execFunc(test.cmd)
		if test.err == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.err.Error(), err.Error())
		}

	}
}