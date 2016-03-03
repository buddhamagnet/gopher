package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	inFile  io.ReadCloser
	outFile io.ReadWriteCloser
)

func fileOpen(filename string) {
	var err error
	if filename != "" {
		if inFile, err = os.Open(filename); err != nil {
			log.Fatal(err)
		}
	}
}

func fileCreate(filename string) {
	var err error
	if filename != "" {
		if outFile, err = os.Create(filename); err != nil {
			log.Fatal(err)
		}
	}
}

func closeFiles() {
	inFile.Close()
	outFile.Close()
}

func main() {
	inFileName, outFileName, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inFile, outFile = os.Stdin, os.Stdout
	fileOpen(inFileName)
	fileCreate(outFileName)
	defer closeFiles()

	if err = americanize(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}

func filenamesFromCommandLine() (inFilename, outFilename string, err error) {
	return "sample.txt", "output.txt", nil
}

func americanize(inFile io.Reader, outFile io.Writer) error {
	return nil
}
