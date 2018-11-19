# freq

Display the line frequency of each line in a file or from STDIN

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

R:\freq>freq test.txt
      4 d
      3 c
      2 b
      1 a
```
