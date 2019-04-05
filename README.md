
# Just a playground for training my Golang skills

### Vendor

```
# verifying dependencies
$ go mod verify
# downloading dependencies from go.mod
$ go get
# establishing the local vendor directory
$ go mod vendor
```

### Building

```
# building using the local vendor directory
$ go build -mod vendor -o build/app
```

### Testing

```
# launch all benchmarks, with memory info
$ go test -bench . -benchmem
# or a specific func name pattern
$ go test -bench BranchPrediction -benchmem

# $ go test -bench BranchPrediction -benchmem
# goos: linux
# goarch: amd64
# pkg: github.com/itnelo/go-training
# 				testName-CPUCount			operations	  speed for 1 op     bandwidth       bytes allocs per op   allocs calls per op
# BenchmarkBranchPrediction-8                  	     300	   4336272 ns/op	 	 1934.52 MB/s	       0 B/op	       0 allocs/op
# BenchmarkBranchPredictionBitwise-8           	    2000	    770194 ns/op		10891.54 MB/s	       0 B/op	       0 allocs/op
# BenchmarkParallelBranchPredictionIfElse-8    	    2000	    664332 ns/op		12627.13 MB/s	       1 B/op	       0 allocs/op
# BenchmarkParallelBranchPredictionBitwise-8   	   10000	    182866 ns/op		45872.81 MB/s	       0 B/op	       0 allocs/op
# PASS
# ok  	github.com/itnelo/go-training	6.889s

# restrict parallel benchmarks to 2 processes, short mode
$ go test -bench BranchPrediction -benchmem -cpu 2 -short
```