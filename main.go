package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
)

func main() {
	var fileName = flag.String("f", "hoge.drawio", "specify drawio filename")
	var outDirName = flag.String("o", ".", "specify output directory")
	flag.Parse()
	if *fileName == "hoge.drawio" {
		log.Fatal("ファイル名を指定してください")
	}
	if !isExistDir(*outDirName) {
		log.Fatal("そのようなディレクトリは存在しません")
	}

	drawioPath, err := chooseDrawioPath()
	if err != nil {
		log.Fatal(err)
	}

	drawioFile, err := NewDrawioFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("drawio page size: ", drawioFile.Pages)

	baseFileName := getFileNameWithoutExt(*fileName)
	for i := 0; i < drawioFile.Pages; i++ {
		pageNumber := strconv.Itoa(i)
		outFileName := filepath.Join(*outDirName, fmt.Sprintf("%s-%d.png", baseFileName, i+1))

		cmd := exec.Command(drawioPath, "-x", "-f", "png", "-o", outFileName, "-p", pageNumber, *fileName)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("exported %s -> %s\n", drawioFile.Diagrams[i].Name, outFileName)
	}

}
