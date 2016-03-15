package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	inFile          io.ReadCloser
	outFile         io.ReadWriteCloser
	wordRx          *regexp.Regexp
	replacer        func(string) string
	britishAmerican string
)

func init() {
	wordRx = regexp.MustCompile("[A-Za-z]+")
	britishAmerican = "british-american.txt"
}

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
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		return "", "", fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt", filepath.Base(os.Args[0]))
	}
	if len(os.Args) > 1 {
		inFilename = os.Args[1]
		if len(os.Args) > 2 {
			outFilename = os.Args[2]
		}
	}
	if inFilename != "" && inFilename == outFilename {
		log.Fatal("cannot overwrite input file")
	}
	return inFilename, outFilename, nil
}

func americanize(inFile io.Reader, outFile io.Writer) (err error) {
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()
	if replacer, err = makeReplaceFunction(britishAmerican); err != nil {
		return err
	}
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}
		line = wordRx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

func makeReplaceFunction(file string) (func(string) string, error) {
	rawBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	text := string(rawBytes)
	usForBritish := make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}
	return func(word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}
		return word
	}, nil
}
