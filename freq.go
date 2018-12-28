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

var BuildTime string

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
    args_percent := flag.Bool("p", false, "output using percentages")
    args_version := flag.Bool("v", false, "display version and then exit")
    flag.Parse()

    if true == *args_version {
        fmt.Println("version:", BuildTime)
        return
    }

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

    var total uint32;
    if true == *args_percent {
        for _, count := range tbl {
            total += count
        }
    }

    if true == *args_ascend {
        sortAscending(unique)
    } else {
        sortDescending(unique)
    }

	lineEnding := "\n"
	if "windows" == runtime.GOOS {
		lineEnding = "\r\n"
	}

    // code redundancy in order to increase speed
    if *args_first > 0 {
        if true == *args_percent {
            for i, entry := range unique {
                var percentage float32
                percentage = float32(entry.count) / float32(total)
                percentage *= 100
                fmt.Printf("%7.1f\t%s%s", percentage, entry.data, lineEnding)
                if i+1 ==  *args_first {
                    break
                }
            }
        } else {
            for i, entry := range unique {
                fmt.Printf("%7d\t%s%s", entry.count, entry.data, lineEnding)
                if i+1 ==  *args_first {
                    break
                }
            }
        }
    } else {
        if true == *args_percent {
            var percentage float32
            for _, entry := range unique {
                percentage = float32(entry.count) / float32(total)
                percentage *= 100
                fmt.Printf("%7.1f\t%s%s", percentage, entry.data, lineEnding)
            }
        } else {
            for _, entry := range unique {
                fmt.Printf("%7d\t%s%s", entry.count, entry.data, lineEnding)
            }
       }
   }
}
