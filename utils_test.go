package main

import (
	"path/filepath"
	"testing"
)

func TestPathJoin(t *testing.T) {
	tests := []struct {
		name     string
		basePath string
		subPath  string
		expected string
	}{
		{".. + subfolder", "..", "subfolder", "../subfolder"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// パスを結合する
			combinedPath := filepath.Join(tt.basePath, tt.subPath)

			if combinedPath != tt.expected {
				t.Errorf("結果が期待される値と一致しません。期待値: %v, 結果: %v", tt.expected, combinedPath)
			}
		})
	}
}

func TestIsExistDir(t *testing.T) {
	tests := []struct {
		name     string
		dirName  string
		expected bool
	}{
		{"ExistingDirectory-01", "testdata", true},
		{"ExistingDirectory-02", "./", true},
		{"NonExistingDirectory", "non_existing_directory", false},
		{"ExistingFile", "testfile.txt", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isExistDir(tt.dirName)

			if result != tt.expected {
				t.Errorf("結果が期待される値と一致しません。期待値: %v, 結果: %v", tt.expected, result)
			}
		})
	}
}
