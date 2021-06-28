package oscmd_test

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/audrenbdb/oscmd"
	"github.com/stretchr/testify/assert"
)

func TestRmkDirInUserHomeDirOnLinux(t *testing.T) {
	//integration test done 28/06 11:00 UTC
	t.Skip()
	if runtime.GOOS != "linux" {
		return
	}

	homeDir, err := os.UserHomeDir()
	if assert.NoError(t, err) {
		tests := []struct {
			dirPath string
			outFullPath string
			outErr error
		}{
			{
				dirPath: "test",
				outFullPath: path.Join(homeDir, "test/"),
				outErr:  nil,
			},
			{
				dirPath: "/opt/test",
				outFullPath: path.Join(homeDir, "/opt/test/"),
				outErr:  nil,
			},
		}

		for _, test := range tests {
			rmkDirFunc := oscmd.NewRmkDirInUserHomeDirFunc()
			fullPath, err := rmkDirFunc(test.dirPath)
			if test.outErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, test.outFullPath, fullPath)
			}
		}
	}



}
