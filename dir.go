//oscmd provides a wrapper around os commands to ease unit testing
package oscmd

import (
	"os"
	"path"
)

//RmkDirInUserHomeDir remove all folder and subfolder
//And create a new dir with given perm
type RmkDirInUserHomeDir = func(dirPath string) error

func NewRmkDirInUserHomeDirFunc() RmkDirInUserHomeDir {
	return func(dirPath string) error {
		fullPath, err := getDirFullPathFromUserHomeDir(dirPath)
		if err != nil {
			return err
		}
		return rmkDir(fullPath)
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