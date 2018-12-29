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

func outputActual(unique []Line, count int, lineEnding string) {
    for i := 0; i <= count; i++ {
        fmt.Printf("%7d\t%s%s", unique[i].count, unique[i].data, lineEnding)
    }
}

func outputPercentage(unique []Line, count int, total float32, lineEnding string) {
    var percentage float32
    for i := 0; i <= count; i++ {
        percentage = 100 * (float32(unique[i].count) / total)
        fmt.Printf("%7.1f\t%s%s", percentage, unique[i].data, lineEnding)
    }
}

func main() {
    argsAscend := flag.Bool("a", false, "output results in ascending order")
    argsLower := flag.Bool("l", false, "convert to lowercase first")
    argsFirst := flag.Int("n", 0, "only output the first N results")
    argsPercent := flag.Bool("p", false, "output using percentages")
    argsVersion := flag.Bool("v", false, "display version and then exit")
    flag.Parse()

    if *argsVersion {
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

    // read input line-by-line to populate 'tbl' hashtable
    tbl := make(map[string]uint32)
    if *argsLower {
        for input.Scan() {
            tbl[strings.ToLower(input.Text())]++
        }
    } else {
        for input.Scan() {
            tbl[input.Text()]++
        }
    }

    // 'unique' is used for sorting
    var unique []Line
    for data, count := range tbl {
        unique = append(unique, Line{data, count})
    }

    // 'total' is used for the percentage divisor
    var total uint32;
    if *argsPercent {
        for _, count := range tbl {
            total += count
        }
    }

    // run an in-place sort of 'unique'
    if *argsAscend {
        sortAscending(unique)
    } else {
        sortDescending(unique)
    }

    lineEnding := "\n"
    if "windows" == runtime.GOOS {
        lineEnding = "\r\n"
    }

    // 'displayCount' is the number of entries to output
    displayCount := len(unique) - 1
    if *argsFirst > 0 {
        displayCount = *argsFirst - 1
    }
    if displayCount >= len(unique) {
        displayCount = len(unique) - 1
    }

    // display the results to STDOUT
    if *argsPercent {
        outputPercentage(unique, displayCount, float32(total), lineEnding)
    } else {
        outputActual(unique, displayCount, lineEnding)
    }
}

