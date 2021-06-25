package oscmd_test

import (
	"runtime"
	"testing"

	"github.com/audrenbdb/oscmd"
	"github.com/stretchr/testify/assert"
)

func TestRmkDirInUserHomeDirOnLinux(t *testing.T) {
	//integration test done 25/06 11:57 UTC
	t.Skip()
	if runtime.GOOS != "linux" {
		return
	}

	tests := []struct {
		dirPath string

		outErr error
	}{
		{
			dirPath: "test",
			outErr:  nil,
		},
		{
			dirPath: "/opt/test",
			outErr:  nil,
		},
	}

	for _, test := range tests {
		rmkDirFunc := oscmd.NewRmkDirInUserHomeDirFunc()
		err := rmkDirFunc(test.dirPath)
		if test.outErr != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}

}
