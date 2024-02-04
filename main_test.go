package main

import (
	"flag"
	"testing"
)

func TestOptionsArgument(t *testing.T) {
	filePath := "testdata/valid.drawio"
	tests := []struct {
		name string
		args []string
		want int
	}{
		// 正常系
		// {name: "-h", args: []string{"-h"}, want: 0},
		// {name: "--help", args: []string{"--help"}, want: 0},
		{name: "-f only", args: []string{"-f", filePath}, want: 0},
		{name: "-o ./testdata", args: []string{"-f", filePath, "-o", "./testdata"}, want: 0},
		// 異常系
		{name: "no options", args: nil, want: 1},
		{name: "no -f options", args: []string{"-o", "."}, want: 1},
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
		{name: "-o ./testdata", args: []string{"-f", nonExtensionfile, "-o", "./testdata"}, want: 0},
		// 異常系
		{name: "non-existent file", args: []string{"-f", nonExistentFile, "-o", "./testdata"}, want: 1},
		{name: "non-existent file", args: []string{"-f", invalidFile, "-o", "./testdata"}, want: 1},
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
