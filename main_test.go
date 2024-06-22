package main

import (
	"flag"
	"testing"
)

// CLIのオプションと引数のテスト
func Testオプション(t *testing.T) {
	filePath := "testdata/valid.drawio"
	filePath2 := "testdata/valid2"
	tests := []struct {
		name string
		args []string
		want int
	}{
		// 正常系
		{name: "-oオプションなし->正常終了", args: []string{filePath}, want: 0},
		{name: "-oオプションあり->正常終了", args: []string{"-o", "./testdata", filePath}, want: 0},
		{name: "複数ファイル指定->正常終了", args: []string{"-o", "./testdata", filePath, filePath2}, want: 0},
		// 異常系
		{name: "引数指定なし->異常終了", args: nil, want: 1},
		{name: "ファイル指定なし->異常終了", args: []string{"-o", "."}, want: 1},
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

func Testファイル読み込み(t *testing.T) {
	nonExtensionfile := "testdata/valid2"
	nonExistentFile := "testdata/non-existent.drawio"
	invalidFile := "testdata/invalid.drawio"
	tests := []struct {
		name string
		args []string
		want int
	}{
		// 正常系
		{name: "拡張子なしのファイル指定 -> 正常終了", args: []string{"-o", "./testdata", nonExtensionfile}, want: 0},
		// 異常系
		{name: "存在しないファイル指定 -> 異常終了", args: []string{"-o", "./testdata", nonExistentFile}, want: 1},
		{name: "無効なファイル指定 -> 異常終了", args: []string{"-o", "./testdata", invalidFile}, want: 1},
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
