# Dynamic Shuffle Check

```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/crhntr/alg/dynamicshuffle
BenchmarkIsShuffleIterative6-4        	 5000000	       295 ns/op
BenchmarkIsShuffleIterative12-4       	 2000000	       594 ns/op
BenchmarkIsShuffleIterative120-4      	  100000	     23086 ns/op
BenchmarkIsShuffleIterative240-4      	   20000	     78039 ns/op
BenchmarkIsShuffleIterative1200-4     	    1000	   1781639 ns/op
BenchmarkIsShuffleIterative2400-4     	     200	   7260337 ns/op
BenchmarkIsShuffleIterative12000-4    	      10	 176584035 ns/op
BenchmarkIsShuffleIterative24000-4    	       2	 779736519 ns/op
BenchmarkIsShuffleIterative120000-4   	       1	23367554761 ns/op
BenchmarkIsShuffleRecursive6-4        	20000000	        59.4 ns/op
BenchmarkIsShuffleRecursive12-4       	20000000	       101 ns/op
BenchmarkIsShuffleRecursive120-4      	 1000000	      1525 ns/op
BenchmarkIsShuffleRecursive240-4      	  500000	      3314 ns/op
BenchmarkIsShuffleRecursive1200-4     	  100000	     19658 ns/op
BenchmarkIsShuffleRecursive2400-4     	   30000	     45011 ns/op
BenchmarkIsShuffleRecursive12000-4    	    5000	    254381 ns/op
BenchmarkIsShuffleRecursive24000-4    	    2000	    587108 ns/op
BenchmarkIsShuffleRecursive120000-4   	       1	2168333656 ns/op
PASS
ok  	github.com/crhntr/alg/dynamicshuffle	58.512s
```
