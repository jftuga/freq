# freq

Display the line frequency of each line in a file or from STDIN

The [Releases Page](https://github.com/jftuga/freq/releases) contains binaries for Windows, MacOS, Linux and FreeBSD.

Usage:
```
Usage of freq:

  -N int
    	only output the last N results, useful with -a
  -a	output results in ascending order
  -d  if line only contains IP address, resolve to hostname
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

## Speed

For input greater than a few hundred megs in size, `freq` is faster than:

    sort | uniq -c | sort -nr
    # or
    awk '{a[$0]++}END{for(i in a){print a[i] " " i}}' | sort -nr
    
but slower than something like:

    export LC_ALL=C
    sort -S 8G --parallel=4 -T /mnt/fast_ssd/tmp | uniq -c | sort -n -r -S 8G --parallel=4 -T /mnt/fast_ssd/tmp
    
See also:  https://www.reddit.com/r/commandline/comments/a7hq5n/psa_improving_gnu_sort_speed/
    
