package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

func main() {
	var fileName = flag.String("f", "hoge.drawio", "specify drawio filename")
	flag.Parse()
	if *fileName == "hoge.drawio" {
		log.Fatal("ファイル名を指定してください")
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

	for i := 0; i < drawioFile.Pages; i++ {
		pageNumber := strconv.Itoa(i + 1)
		outputFileName := fmt.Sprintf("%s.png", pageNumber)
		cmd := exec.Command(drawioPath, "-x", "-f", "png", "-o", outputFileName, "-p", pageNumber, *fileName)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("exported %s -> %s\n", drawioFile.Diagrams[i].Name, outputFileName)
	}

}
