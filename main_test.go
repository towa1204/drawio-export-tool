package main

import (
	"flag"
	"testing"
)

// CLIのオプションと引数のテスト
func TestOptionsArgument(t *testing.T) {
	filePath := "testdata/valid.drawio"
	tests := []struct {
		name string
		args []string
		want int
	}{
		// 正常系
		{name: "no options", args: []string{filePath}, want: 0},
		{name: "-o ./testdata", args: []string{"-o", "./testdata", filePath}, want: 0},
		// 異常系
		{name: "no argment", args: nil, want: 1},
		{name: "no file", args: []string{"-o", "."}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandLine = flag.NewFlagSet("drawio-export", flag.ExitOnError)
			if got := run(tt.args); got != tt.want {
				t.Errorf("%v: run() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestImportFile(t *testing.T) {
	nonExtensionfile := "testdata/valid2"
	nonExistentFile := "testdata/non-existent.drawio"
	invalidFile := "testdata/invalid.drawio"
	tests := []struct {
		name string
		args []string
		want int
	}{
		// 正常系
		{name: "-o ./testdata", args: []string{"-o", "./testdata", nonExtensionfile}, want: 0},
		// 異常系
		{name: "non-existent file", args: []string{"-o", "./testdata", nonExistentFile}, want: 1},
		{name: "non-existent file", args: []string{"-o", "./testdata", invalidFile}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandLine = flag.NewFlagSet("drawio-export", flag.ExitOnError)
			if got := run(tt.args); got != tt.want {
				t.Errorf("%v: run() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkRun(b *testing.B) {
	filePath := "testdata/valid.drawio"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		commandLine = flag.NewFlagSet("drawio-export", flag.ExitOnError)
		run([]string{"-o", "./testdata", filePath})
	}
}
