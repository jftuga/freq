/*
freq.go
-John Taylor

Display the frequency of each line in a file or from STDIN

To compile:
go build -ldflags="-s -w" freq.go

MIT License; Copyright (c) 2019 John Taylor
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main


import (
    "bufio"
    "fmt"
    "testing"
)

func TestFreq(t *testing.T) {
    data1 := []string {
        "total 28",
        "-rw-rw-r-- 1 jftuga jftuga  1068 Sep  5 15:41 LICENSE",
        "-rw-rw-r-- 1 jftuga jftuga   101 Sep  5 15:41 Makefile",
        "-rw-rw-r-- 1 jftuga jftuga  1938 Sep  5 15:41 README.md",
        "-rw-rw-r-- 1 jftuga jftuga 11354 Sep  5 15:41 freq.go",
        "-rw-rw-r-- 1 jftuga jftuga  1310 Sep  5 16:29 freq_test.go",
        "-rw-rw-r-- 1 jftuga jftuga     0 Sep  5 16:37 testfile",
    }

    data2 := []string {
        "usr","usr","usr","usr","usr","usr",
        "sys","sys","sys","sys","sys","sys","sys","sys","sys","sys","sys",
        "system","system","system","system","system",
        "Usr", "Usr", "Usr",
        "SySTem", "SySTem",
        "User","User","User","User","User","User","User","User","User","User","User","User",
    }

    data3 := []string { "1.1.1.1", "4.2.2.1", "8.8.8.8", "9.9.9.9", }


    fmt.Println(len(data1))
    fmt.Println(len(data2))
    fmt.Println(len(data3))
    fmt.Println()

    var input *bufio.Scanner
    input = bufio.NewScanner(data1)

    for input.Scan() {
        fmt.Println("input:", input.Text())
    }
}

