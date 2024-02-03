package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"runtime"
)

type DrawioFile struct {
	XMLName  xml.Name  `xml:"mxfile"`
	Modified string    `xml:"modified,attr"`
	Pages    int       `xml:"pages,attr"`
	Diagrams []Diagram `xml:"diagram"`
}

type Diagram struct {
	XMLName xml.Name `xml:"diagram"`
	Name    string   `xml:"name,attr"`
	ID      string   `xml:"id,attr"`
}

// NewDrawioFile はfileNameを開きDrawioFile構造体にマッピングする
func NewDrawioFile(fileName string) (*DrawioFile, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	xmlData, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var drawioFile DrawioFile
	err = xml.Unmarshal(xmlData, &drawioFile)
	if err != nil {
		return nil, err
	}

	return &drawioFile, nil
}

// getDrawioExecutablePath はOSに応じてdrawioの実行コマンドパスを切り替える
func getDrawioExecutablePath() (string, error) {
	if runtime.GOOS == "linux" {
		return "drawio", nil
	} else if runtime.GOOS == "windows" {
		return "draw.io", nil
	}
	return "", fmt.Errorf("this software don't support %s", runtime.GOOS)
}
