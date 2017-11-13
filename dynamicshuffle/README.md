# Dynamic Shuffle Check

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
