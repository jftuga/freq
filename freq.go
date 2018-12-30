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

var version string

// Slices are passed by reference
func sortInput(unique []Line, descending bool) {
    sort.Slice(unique, func(i, j int) bool {
        if unique[i].count > unique[j].count {
            return descending
        }
        if unique[i].count < unique[j].count {
            return !descending
        }
        // when multiple lines have the same count, then alphabetize these lines
        return unique[i].data < unique[j].data
    })
}

func outputActual(unique []Line, start int, count int, lineEnding string) {
    if start > 0 {
        start = count - start + 1
    }
    if start < 0 {
        start = 0
    }
    for i := start; i <= count; i++ {
        fmt.Printf("%7d\t%s%s", unique[i].count, unique[i].data, lineEnding)
    }
}

func outputPercentage(unique []Line, start int, count int, total float32, lineEnding string) {
    if start > 0 {
        start = count - start + 1
    }
    if start < 0 {
        start = 0
    }

    var percentage float32
    for i := start; i <= count; i++ {
        percentage = 100 * (float32(unique[i].count) / total)
        fmt.Printf("%7.1f\t%s%s", percentage, unique[i].data, lineEnding)
    }
}

func main() {
    argsAscend := flag.Bool("a", false, "output results in ascending order")
    argsLower := flag.Bool("l", false, "convert to lowercase first")
    argsFirst := flag.Int("n", 0, "only output the first N results")
    argsLast := flag.Int("N", 0, "only output the last N results, useful with -a")
    argsPercent := flag.Bool("p", false, "output using percentages")
    argsVersion := flag.Bool("v", false, "display version and then exit")
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "\n%s %s, display the frequency of each line in a file or from STDIN.\n\n", os.Args[0], version)
        fmt.Fprintf(os.Stderr, "Usage for %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

    flag.Parse()
    if *argsVersion {
        fmt.Println("version:", version)
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
    sortInput(unique, *argsAscend)

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

    // determine where to start the results output
    start := 0
    if *argsLast > 0 {
        start = *argsLast
    }

    // display the results to STDOUT
    if *argsPercent {
        outputPercentage(unique, start, displayCount, float32(total), lineEnding)
    } else {
        outputActual(unique, start, displayCount, lineEnding)
    }
}

