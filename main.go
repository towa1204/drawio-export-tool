package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var commandLine = flag.NewFlagSet("drawio-export", flag.ExitOnError)

func run(args []string) int {
	fileName := commandLine.String("f", "", "specify drawio filename")
	outDirName := commandLine.String("o", ".", "specify output directory")
	if err := commandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse flags: %v\n", err)
		return 1
	}

	if *fileName == "" {
		fmt.Fprintln(os.Stderr, "filename must be provided")
		return 1
	}
	if !isExistDir(*outDirName) {
		fmt.Fprintln(os.Stderr, "invalid directory")
		return 1
	}

	drawioPath, err := getDrawioExecutablePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unsupported OS: %v\n", err)
		return 1
	}

	drawioFile, err := NewDrawioFile(*fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse %s: %v\n", *fileName, err)
		return 1
	}

	fmt.Println("drawio page size: ", drawioFile.Pages)

	baseFileName := removeExtension(*fileName)
	for i := 0; i < drawioFile.Pages; i++ {
		pageNumber := strconv.Itoa(i)
		outFileName := filepath.Join(*outDirName, fmt.Sprintf("%s-%d.png", baseFileName, i+1))

		cmd := exec.Command(drawioPath, "-x", "-f", "png", "-o", outFileName, "-p", pageNumber, *fileName)
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[page-number %s] failed to export: %v\n", pageNumber, err)
			return 1
		}

		fmt.Printf("exported %s -> %s\n", drawioFile.Diagrams[i].Name, outFileName)
	}
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
