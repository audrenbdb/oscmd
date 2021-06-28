//Package oscmd provides a wrapper around os commands to ease unit testing
package oscmd

import (
	"os"
	"path"
)

//RmkDirInUserHomeDir and returns its full path
type RmkDirInUserHomeDir = func(dirPath string) (string, error)

func NewRmkDirInUserHomeDirFunc() RmkDirInUserHomeDir {
	return func(dirPath string) (string, error) {
		fullPath, err := getDirFullPathFromUserHomeDir(dirPath)
		if err != nil {
			return "", err
		}
		return fullPath, rmkDir(fullPath)
	}
}

func rmkDir(fullPath string) error {
	err := os.RemoveAll(fullPath)
	if err != nil {
		return err
	}
	return os.MkdirAll(fullPath, os.ModePerm)
}

func getDirFullPathFromUserHomeDir(dirPath string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, dirPath), nil
}