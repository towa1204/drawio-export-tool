package main

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func removeExtension(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// isExistDir は与えられた文字列がディレクトリとして存在するか判定する
func isExistDir(dirName string) bool {
	absPath, err := filepath.Abs(dirName)
	if err != nil {
		return false
	}

	fileInfo, err := os.Stat(absPath)

	if !errors.Is(err, fs.ErrNotExist) && err == nil && fileInfo.IsDir() {
		return true
	}

	return false
}
