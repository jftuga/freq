/*
freq.go
-John Taylor

Display the line frequency of each line in a file or from STDIN

To compile:
go build -ldflags="-s -w" freq.go
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
)

type Line struct {
	data  string
	count uint32
}

func main() {
	var input *bufio.Scanner
	if 1 == len(os.Args) { // read from STDIN
		input = bufio.NewScanner(os.Stdin)
	} else { // read from filename
		fname := os.Args[1]
		file, err := os.Open(fname)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		input = bufio.NewScanner(file)
	}

	tbl := make(map[string]uint32)
	for input.Scan() {
		tbl[input.Text()]++
	}

	var unique []Line
	for data, count := range tbl {
		unique = append(unique, Line{data, count})
	}

	// when multiple lines have the same count, then
	// alphabetize these lines
	sort.Slice(unique, func(i, j int) bool {
		if unique[i].count > unique[j].count {
			return true
		}
		if unique[i].count < unique[j].count {
			return false
		}
		return unique[i].data < unique[j].data
	})

	// Unbelievable but true:
	// Go1.11.1 does not output CRLF on Windows with Println or Printf
	lineEnding := "\n"
	if "windows" == runtime.GOOS {
		lineEnding = "\r\n"
	}
	for _, entry := range unique {
		fmt.Printf("%7d\t%s%s", entry.count, entry.data, lineEnding)
	}
}
