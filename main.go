package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"golang.org/x/sync/errgroup"
)

var commandLine = flag.NewFlagSet("drawio-export", flag.ExitOnError)

func run(args []string) int {
	outDirName := commandLine.String("o", ".", "specify output directory")
	skipOption := commandLine.Bool("s", false, "skip export same name file")
	if err := commandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse flags: %v\n", err)
		return 1
	}

	fileNames := commandLine.Args()
	if len(fileNames) == 0 {
		fmt.Fprintln(os.Stderr, "filename must be provided")
		return 1
	}

	if !isExistDir(*outDirName) {
		fmt.Fprintln(os.Stderr, "invalid directory")
		return 1
	}

	drawioPath, err := GetDrawioExecutablePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unsupported OS: %v\n", err)
		return 1
	}

	// ファイルごとに全ページエクスポートを実行
	// 【エラーハンドリングの方針】
	// ・エクスポートに失敗したページがあっても継続する
	// ・失敗したページは標準エラー出力先に出力する
	// ・1つでも失敗したページがあれば、終了ステータスは1にする
	errFlg := false
	for _, fileName := range fileNames {
		fileName := fileName
		err := ExportAllPage(drawioPath, fileName, *outDirName, *skipOption)
		if err != nil {
			errFlg = true
		}
	}

	if errFlg {
		return 1
	}
	return 0
}

// 指定ファイルの全ページをpngエクスポート
func ExportAllPage(drawioPath, fileName, outDirName string, skipOption bool) error {
	drawioFile, err := NewDrawioFile(fileName)
	if err != nil {
		message := fmt.Sprintf("failed to parse %s: %v", fileName, err)
		fmt.Fprintln(os.Stderr, message)
		return fmt.Errorf(message)
	}

	// fmt.Println("drawio page size: ", drawioFile.Pages)

	var eg errgroup.Group
	eg.SetLimit(10)

	baseFileName := removeExtension(fileName)
	for i := 0; i < drawioFile.Pages; i++ {
		i := i
		eg.Go(func() error {
			pageName := drawioFile.Diagrams[i].Name
			pageNumber := strconv.Itoa(i)
			outFileName := filepath.Join(outDirName, fmt.Sprintf("%s-%d.png", baseFileName, i+1))

			if skipOption && isExistFile(outFileName) {
				fmt.Printf("skip export %s %s\n", fileName, pageName)
				return nil
			}

			cmd := exec.Command(drawioPath, "-x", "-f", "png", "-o", outFileName, "-p", pageNumber, fileName)
			err := cmd.Run()
			if err != nil {
				message := fmt.Sprintf("failed to export %s %s: %v", fileName, pageName, err)
				fmt.Fprintln(os.Stderr, message)
				return fmt.Errorf(message)
			}

			// fmt.Printf("exported %s -> %s\n", drawioFile.Diagrams[i].Name, outFileName)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func main() {
	os.Exit(run(os.Args[1:]))
}
