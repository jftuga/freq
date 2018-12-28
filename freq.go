/*
freq.go
-John Taylor

Display the frequency of each line in a file or from STDIN

To compile:
go build -ldflags="-s -w" freq.go
*/

package main

import (
	"bufio"
    "flag"
	"fmt"
	"os"
	"runtime"
	"sort"
    "strings"
)

type Line struct {
	data  string
	count uint32
}

// Slices are passed by reference
func sortDescending(unique []Line) {
	// when multiple lines have the same count, then alphabetize these lines
	sort.Slice(unique, func(i, j int) bool {
		if unique[i].count > unique[j].count {
			return true
		}
		if unique[i].count < unique[j].count {
			return false
		}
		return unique[i].data < unique[j].data
	})
}

func sortAscending(unique []Line) {
	// when multiple lines have the same count, then alphabetize these lines
	sort.Slice(unique, func(i, j int) bool {
		if unique[i].count < unique[j].count {
			return true
		}
		if unique[i].count > unique[j].count {
			return false
		}
		return unique[i].data < unique[j].data
	})
}

func main() {
    args_ascend := flag.Bool("a", false, "output results in ascending order")
    args_lower := flag.Bool("l", false, "convert to lowercase first")
    args_first := flag.Int("n", 0, "only output the first N results")
    flag.Parse()

	var input *bufio.Scanner
    args := flag.Args()

	if 0 == len(args) { // read from STDIN
		input = bufio.NewScanner(os.Stdin)
	} else { // read from filename
		fname := args[0]
		file, err := os.Open(fname)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		input = bufio.NewScanner(file)
	}

	tbl := make(map[string]uint32)
    if true == *args_lower {
        for input.Scan() {
            tbl[strings.ToLower(input.Text())]++
        }
    } else {
        for input.Scan() {
            tbl[input.Text()]++
        }
    }

	var unique []Line
	for data, count := range tbl {
		unique = append(unique, Line{data, count})
	}

    if true == *args_ascend {
        sortAscending(unique)
    } else {
        sortDescending(unique)
    }

	// Unbelievable but true:
	// Go1.11.1 does not output CRLF on Windows with Println or Printf
	lineEnding := "\n"
	if "windows" == runtime.GOOS {
		lineEnding = "\r\n"
	}
    if *args_first > 0 {
        for i, entry := range unique {
            fmt.Printf("%7d\t%s%s", entry.count, entry.data, lineEnding)
            if i+1 ==  *args_first {
                break
            }
        }
    } else {
        for _, entry := range unique {
            fmt.Printf("%7d\t%s%s", entry.count, entry.data, lineEnding)
        }
   }
}
