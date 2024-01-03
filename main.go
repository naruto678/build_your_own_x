package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	byteCount := flag.Bool("c", false, "parse byte counts")
	lineCount := flag.Bool("l", false, "get line counts")
	wordCount := flag.Bool("w", false, "get word counts")
	charCount := flag.Bool("m", false, "get byte count")
	flag.Parse()
	tails := flag.Args()

	content := []byte{}
	fileName := ""

	if len(tails) == 0 {
		ioContent, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		content = ioContent
		// no file name is specified . try reading from stdin
	} else {

		fileName = tails[0]
		fileContent, err := os.ReadFile(fileName)
		content = fileContent
		if err != nil {
			panic(err)
		}
	}

	defaultOpt := true
	if *byteCount {
		result := CountBytes(content)
		fmt.Fprintf(os.Stdout, "%d %s\n", result, fileName)
		defaultOpt = false
	}
	if *lineCount {
		result := CountLines(content)
		fmt.Fprintf(os.Stdout, "%d %s\n", result, fileName)
		defaultOpt = false
	}

	if *wordCount {
		result := CountWords(content)
		fmt.Fprintf(os.Stdout, "%d %s\n", result, fileName)
		defaultOpt = false
	}
	if *charCount {
		result := CountChars(content)
		fmt.Fprintf(os.Stdout, "%d %s\n", result, fileName)
		defaultOpt = false
	}

	if defaultOpt {
		byteCount := CountBytes(content)
		lineCount := CountLines(content)
		wordCount := CountWords(content)
		fmt.Fprintf(os.Stdout, "  %d %d %d %s\n", lineCount, wordCount, byteCount, fileName)

	}

}

func CountBytes(content []byte) int {
	return len(content)

}
func CountLines(content []byte) int {
	newLine := byte('\n')
	count := 0
	for _, val := range content {
		if val == newLine {
			count++
		}
	}
	return count

}
func CountWords(content []byte) int {
	count := 0
	strContent := string(content)
	for _, line := range strings.Split(strContent, "\n") {
		count += len(strings.Fields(line))
	}
	return count

}
func CountChars(content []byte) int {
	strContent := string(content)
	return len(strContent)
}
