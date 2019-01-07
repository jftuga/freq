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
    "net"
    "os"
    "regexp"
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
func sortInput(unique []Line, ascending bool) {
    sort.Slice(unique, func(i, j int) bool {
        if unique[i].count > unique[j].count {
            return !ascending
        }
        if unique[i].count < unique[j].count {
            return ascending
        }
        // when multiple lines have the same count, then alphabetize these lines
        return unique[i].data < unique[j].data
    })
}

var dnsCache map[string]string
func dnsLookup(ip string) string {
    // 'cached' should never be used since all output is now unique
    // so this may be removed in the future
    cached := dnsCache[ip]
    if len(cached) > 0 {
        return cached
    }
    addresses, err := net.LookupAddr(ip)
    if err != nil {
        return ip
    }

    if len(addresses) == 0 {
        return ip
    }
    resolved := strings.TrimSuffix(addresses[0],".")
    resolved = strings.ToLower(resolved)
    dnsCache[ip] = resolved
    return resolved
}

func output(unique []Line, start int, count int, total float32, lineEnding string, usePercentage bool, dnsResolve bool) {
    if start > 0 {
        start = count - start + 1
    }
    if start < 0 {
        start = 0
    }

    dnsRE := regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$`)

    if usePercentage {
        var percentage float32
        for i := start; i <= count; i++ {
            percentage = 100 * (float32(unique[i].count) / total)
            if( dnsResolve ) {
                if( dnsRE.MatchString(unique[i].data) ) {
                    unique[i].data = dnsLookup(unique[i].data)
                }
            }
            fmt.Printf("%7.1f\t%s%s", percentage, unique[i].data, lineEnding)
         }
     } else {
        for i := start; i <= count; i++ {
            if( dnsResolve ) {
                if( dnsRE.MatchString(unique[i].data) ) {
                    unique[i].data = dnsLookup(unique[i].data)
                }
            }
            fmt.Printf("%7d\t%s%s", unique[i].count, unique[i].data, lineEnding)
        }
     }
}

func ReadInput(input *bufio.Scanner, convertToLower bool) map[string]uint32 {
    tbl := make(map[string]uint32)
    if convertToLower {
        for input.Scan() {
            tbl[strings.ToLower(input.Text())]++
        }
    } else {
        for input.Scan() {
            tbl[input.Text()]++
        }
    }
    return tbl
}

func main() {
    argsAscend := flag.Bool("a", false, "output results in ascending order")
    argsLower := flag.Bool("l", false, "convert to lowercase first")
    argsFirst := flag.Int("n", 0, "only output the first N results")
    argsLast := flag.Int("N", 0, "only output the last N results, useful with -a")
    argsPercent := flag.Bool("p", false, "output using percentages")
    argsResolve := flag.Bool("d", false, "if line only contains IP address, resolve to hostname")
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

    dnsCache = make(map[string]string)
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
    tbl := ReadInput(input, *argsLower)

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

    // output the results to STDOUT
    output(unique, start, displayCount, float32(total), lineEnding, *argsPercent, *argsResolve)
}

