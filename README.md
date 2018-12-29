# freq

Display the line frequency of each line in a file or from STDIN

The [Releases Page](https://github.com/jftuga/freq/releases) contains binaries for Windows, MacOS, Linux and FreeBSD.

Usage:
```
Usage of freq:

  -N int
    	only output the last N results, useful with -a
  -a	output results in ascending order
  -l	convert to lowercase first
  -n int
    	only output the first N results
  -p	output using percentages
  -v	display version and then exit
```

Examples:

```
R:\freq>type con > test.txt
d
b
a
b
c
c
d
c
d
d
^Z

R:\freq>type test.txt | freq.exe
      4 d
      3 c
      2 b
      1 a

jftuga@linux:~/go/src/github.com/jftuga/freq$ freq -n 2 test.txt
      4 d
      3 c
```
