# Dynamic Shuffle Check


## Benchmark output
```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/crhntr/alg/dynamicshuffle
BenchmarkIsShuffleIterative6-4                	 5000000	       292 ns/op
BenchmarkIsShuffleIterative12-4               	 2000000	       604 ns/op
BenchmarkIsShuffleIterative120-4              	  100000	     23002 ns/op
BenchmarkIsShuffleIterative240-4              	   20000	     78593 ns/op
BenchmarkIsShuffleIterative1200-4             	    1000	   1804684 ns/op
BenchmarkIsShuffleIterative2400-4             	     200	   7102188 ns/op
BenchmarkIsShuffleIterative12000-4            	      10	 175828255 ns/op
BenchmarkIsShuffleIterative24000-4            	       2	 768624032 ns/op
BenchmarkIsShuffleIterative120000-4           	       1	25299902135 ns/op
BenchmarkIsShuffleRecursive6-4                	20000000	        60.3 ns/op
BenchmarkIsShuffleRecursive12-4               	10000000	       101 ns/op
BenchmarkIsShuffleRecursive120-4              	 1000000	      1544 ns/op
BenchmarkIsShuffleRecursive240-4              	  500000	      3419 ns/op
BenchmarkIsShuffleRecursive1200-4             	  100000	     20276 ns/op
BenchmarkIsShuffleRecursive2400-4             	   30000	     46765 ns/op
BenchmarkIsShuffleRecursive12000-4            	    5000	    260004 ns/op
BenchmarkIsShuffleRecursive24000-4            	    2000	    587121 ns/op
BenchmarkIsShuffleRecursive120000-4           	       1	2220374576 ns/op
BenchmarkIsShuffleIterativeWorstCase6-4       	 5000000	       278 ns/op
BenchmarkIsShuffleIterativeWorstCase12-4      	 3000000	       555 ns/op
BenchmarkIsShuffleIterativeWorstCase120-4     	  100000	     18828 ns/op
BenchmarkIsShuffleIterativeWorstCase240-4     	   20000	     67038 ns/op
BenchmarkIsShuffleIterativeWorstCase1200-4    	    1000	   1527046 ns/op
BenchmarkIsShuffleIterativeWorstCase2400-4    	     200	   6030995 ns/op
BenchmarkIsShuffleIterativeWorstCase12000-4   	      10	 151653147 ns/op
BenchmarkIsShuffleRecursiveWorstCase6-4       	50000000	        40.3 ns/op
BenchmarkIsShuffleRecursiveWorstCase12-4      	20000000	        68.6 ns/op
BenchmarkIsShuffleRecursiveWorstCase120-4     	 3000000	       603 ns/op
BenchmarkIsShuffleRecursiveWorstCase240-4     	 1000000	      1131 ns/op
BenchmarkIsShuffleRecursiveWorstCase1200-4    	  200000	      5978 ns/op
BenchmarkIsShuffleRecursiveWorstCase2400-4    	  100000	     12271 ns/op
BenchmarkIsShuffleRecursiveWorstCase12000-4   	   20000	     77998 ns/op
PASS
ok  	github.com/crhntr/alg/dynamicshuffle	84.304s
```

## Visual Output for Iterative (table) Solution
```
$ go test -v -run .*Iterative.*
=== RUN   TestIsShuffleIterative_000111
w: "010101"

   111
  *---
0 **--
0 -**-
0 --**

--- PASS: TestIsShuffleIterative_000111 (0.01s)
=== RUN   TestIsShuffleIterative_011011
w: "001111"

   011
  **--
0 ****
1 -***
1 -***

--- PASS: TestIsShuffleIterative_011011 (0.01s)
=== RUN   TestIsShuffleIterative_01
--- PASS: TestIsShuffleIterative_01 (0.00s)
=== RUN   TestIsShuffleIterative_10
--- PASS: TestIsShuffleIterative_10 (0.00s)
=== RUN   TestIsShuffleIterative_100
--- PASS: TestIsShuffleIterative_100 (0.00s)
=== RUN   TestIsShuffleIterative_aaaaccebbbbdde
w: "aabbababcdcdee"

   bbbbdde
  *-------
a *-------
a ***-----
a --**----
a ---**---
c ----**--
c -----***
e ------**

--- PASS: TestIsShuffleIterative_aaaaccebbbbdde (0.01s)
=== RUN   TestIsShuffleIterative_ab
w: "aa"

   b
  *-
a *-

--- PASS: TestIsShuffleIterative_ab (0.01s)
=== RUN   TestIsShuffleIterative_aaabbbabcabc
w: "abcabc"

   bbb
  *---
a **--
a ----
a ----

--- PASS: TestIsShuffleIterative_aaabbbabcabc (0.01s)
PASS
ok  	github.com/crhntr/alg/dynamicshuffle	0.071s
```


## Visual Output for Recursive Solution
```
$ go test -v -run .*Recursive.*
=== RUN   TestIsShuffleRecursive_000111
w: "010101", u: "000", v: "111"
w: "01010", u: "000", v: "11"
w: "0101", u: "00", v: "11"
w: "010", u: "00", v: "1"
w: "01", u: "0", v: "1"
w: "0", u: "0", v: ""

--- PASS: TestIsShuffleRecursive_000111 (0.00s)
=== RUN   TestIsShuffleRecursive_011011
w: "001111", u: "011", v: "011"
w: "00111", u: "011", v: "01"
w: "0011", u: "011", v: "0"
w: "001", u: "01", v: "0"
w: "00", u: "0", v: "0"
w: "0", u: "0", v: ""

--- PASS: TestIsShuffleRecursive_011011 (0.00s)
=== RUN   TestIsShuffleRecursive_101101
w: "110101", u: "101", v: "101"
w: "11010", u: "101", v: "10"
w: "1101", u: "101", v: "1"
w: "110", u: "101", v: ""
w: "110", u: "10", v: "1"
w: "11", u: "1", v: "1"
w: "1", u: "1", v: ""

--- PASS: TestIsShuffleRecursive_101101 (0.00s)
=== RUN   TestIsShuffleRecursive_0
w: "0", u: "", v: "0"

--- PASS: TestIsShuffleRecursive_0 (0.00s)
=== RUN   TestIsShuffleRecursive_1
w: "1", u: "1", v: ""

--- PASS: TestIsShuffleRecursive_1 (0.00s)
=== RUN   TestIsShuffleRecursive_BadCase
--- PASS: TestIsShuffleRecursive_BadCase (0.00s)
=== RUN   TestIsShuffleRecursive_aaaaccebbbbdde
w: "aabbababcdcdee", u: "aaaacce", v: "bbbbdde"
w: "aabbababcdcde", u: "aaaacce", v: "bbbbdd"
w: "aabbababcdcd", u: "aaaacc", v: "bbbbdd"
w: "aabbababcdc", u: "aaaacc", v: "bbbbd"
w: "aabbababcd", u: "aaaac", v: "bbbbd"
w: "aabbababc", u: "aaaac", v: "bbbb"
w: "aabbabab", u: "aaaa", v: "bbbb"
w: "aabbaba", u: "aaaa", v: "bbb"
w: "aabbab", u: "aaa", v: "bbb"
w: "aabba", u: "aaa", v: "bb"
w: "aabb", u: "aa", v: "bb"
w: "aab", u: "aa", v: "b"
w: "aa", u: "aa", v: ""

--- PASS: TestIsShuffleRecursive_aaaaccebbbbdde (0.00s)
=== RUN   TestIsShuffleRecursive_aaaaccqbbbbdde
w: "aabbababcdcdee", u: "aaaaccq", v: "bbbbdde"
w: "aabbababcdcde", u: "aaaaccq", v: "bbbbdd"

--- PASS: TestIsShuffleRecursive_aaaaccqbbbbdde (0.00s)
=== RUN   TestIsShuffleRecursive_aaabbbabcabc
w: "abcabc", u: "aaa", v: "bbb"

--- PASS: TestIsShuffleRecursive_aaabbbabcabc (0.00s)
PASS
ok  	github.com/crhntr/alg/dynamicshuffle	0.013s
```
