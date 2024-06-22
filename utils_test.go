package main

import (
	"testing"
)

func TestIsExistDir(t *testing.T) {
	tests := []struct {
		name     string
		dirName  string
		expected bool
	}{
		{name: "testdata/->存在する", dirName: "testdata", expected: true},
		{name: "./->存在する", dirName: "./", expected: true},
		{name: "存在しないディレクトリの指定->存在しない", dirName: "non_existing_directory", expected: false},
		{name: "存在するファイルの指定->存在しない", dirName: "main.go", expected: false},
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
