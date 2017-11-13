# Dynamic Shuffle Check

```
crhntr@crhntr ~/src/github.com/crhntr/alg/dynamicshuffle [master]$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/crhntr/alg/dynamicshuffle
BenchmarkIsShuffleRecursive6-4        	30000000	        58.5 ns/op
BenchmarkIsShuffleRecursive12-4       	20000000	       113 ns/op
BenchmarkIsShuffleRecursive120-4      	 1000000	      1527 ns/op
BenchmarkIsShuffleRecursive240-4      	  500000	      3458 ns/op
BenchmarkIsShuffleRecursive1200-4     	  100000	     19133 ns/op
BenchmarkIsShuffleRecursive2400-4     	   30000	     45242 ns/op
BenchmarkIsShuffleRecursive12000-4    	    5000	    260416 ns/op
BenchmarkIsShuffleRecursive24000-4    	    2000	    591285 ns/op
BenchmarkIsShuffleRecursive120000-4   	       1	2195223362 ns/op
PASS
ok  	github.com/crhntr/alg/dynamicshuffle	17.593s
```
