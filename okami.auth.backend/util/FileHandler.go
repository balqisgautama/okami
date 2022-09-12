package util

import (
	"os"
	"strings"
)

func MoveFile(oldPath string, newPath string) error {
	newPath = strings.Replace(newPath, "\\", "/", -1)
	newPathSplit := strings.Split(newPath, "/")
	newPathMkdir := strings.Join(newPathSplit[0:len(newPathSplit)-1], "/")
	_ = os.MkdirAll(newPathMkdir, 0770)
	return os.Rename(oldPath, newPath)
}
