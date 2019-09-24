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
    "bytes"
    "fmt"
    "strings"
    "testing"
)

func validateResult(t *testing.T, candidate string, correct []string) bool {
    array := strings.Split(candidate, "\n")
    for i := 0; i < len(array)-1; i++ {
        trimmed := strings.TrimSpace(array[i])
        trimmed = strings.Replace(trimmed,"\t"," ", 1)
        if( 0 != strings.Compare(trimmed,correct[i]) ) {
            fmt.Printf("FAILURE: %d  %+q  !=  %+q\n", i, trimmed, correct[i])
            t.Fail()
        } else {
            fmt.Printf("     OK: %d  %+q  ==  %+q\n", i, trimmed, correct[i])
        }
    }

    return true
}

func TestFreq(t *testing.T) {
    /*
    data1 := []string {
        "total 28",
        "-rw-rw-r-- 1 jftuga jftuga     1068 Sep  5 15:41 LICENSE",
        "-rw-rw-r-- 1 jftuga jftuga      101 Sep  5 15:41 Makefile",
        "-rwxrwxr-x 1 jftuga jftuga  2457264 Sep  9 04:24 freq"
        "-rw-rw-r-- 1 jftuga jftuga    11354 Sep  5 15:41 freq.go",
        "-rw-rw-r-- 1 jftuga jftuga     1310 Sep  5 16:29 freq_test.go",
        "-rw-rw-r-- 1 jftuga jftuga        6 Sep  5 16:37 testfile",
    }
    */

    data2 := []string {
        "usr","usr","usr","usr","usr","usr",
        "sys","sys","sys","sys","sys","sys","sys","sys","sys","sys","sys",
        "system","system","system","system","system",
        "Usr", "Usr", "Usr",
        "SySTem", "SySTem", "system", "SYSTEM",
        "User","User","User","User","User","User","User","User","User","User","User","User",
    }

    correct2 := []string {
        "12 user", "11 sys", "9 system", "9 usr",
    }

    //data3 := []string { "1.1.1.1", "4.2.2.1", "8.8.8.8", "9.9.9.9", }

    //blobData1 := []byte( strings.Join(data1,"\n") )
    //inputData1 := bufio.NewScanner(bytes.NewReader(blobData1))
    blobData2 := []byte( strings.Join(data2,"\n") )
    inputData2 := bufio.NewScanner(bytes.NewReader(blobData2))
    //blobData3 := []byte( strings.Join(data3,"\n") )
    //inputData3 := bufio.NewScanner(bytes.NewReader(blobData3))


    // func ReadInput(input *bufio.Scanner, convertToLower bool, substringStart int, substringEnd int, matchRegExp string) map[string]uint32
    // func output(unique []Line, start int, count int, total float32, lineEnding string, usePercentage bool, dnsResolve bool, bare bool) {

    // test -l, lowercase
    tbl := make(map[string]uint32)
    tbl = ReadInput(inputData2, true, 0, 0, "")
    fmt.Println("tbl:", tbl)
    unique, total := uniqueAndSort(tbl, false, false)

    result := output(unique, 0, len(unique)-1, float32(total), "\n", false, false, false)
    //fmt.Printf("result\n------\n_%s_\n\n", result)
    //fmt.Print(result)
    validateResult(t, result, correct2)
}

