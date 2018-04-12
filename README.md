# maximum-box
The maximum box task from "Cracking the Coding Interview"

## Introduction
The Maximum Box problem is this: given a grid of "pixels" where each pixel can have one of two values, on or off, find the maximum square such that all four sides are composed of "on" pixels. 

The initial version of this challenge would be to assume that all squares are aligned vertically allowing the solution to only look for horizontal and vertical sides.

An interesting extension is to allow the sides to be at multiple angles.

## Approaches
### Clever
I plan on tackling this in (at least conceptual) stages.
 1. Using a line-sweep to find line segments
 1. Use a segment intersection to find corners.
 1. assemble rectangles out of doubly intersecting corners
 1. filter out none-squares
 1. find the largest square
### Brute Force
In the case where squares are aligned with the grid, one could:
 1. iterate through the entire grid looking for "on" pixels, call this P
 1. On pixels are potentially the first (eg upper left corner) pixel in a square so:
    1. Iterate right while the pixels are on, making a note (pushing the offset from P onto a stack) of pixels with an on "down" neighbor
    1. Iterate down while the pixels are on, making a note (pushing the offset from P onto a stack) of pixels with an on "right" neighbor
    1. for each offset in common in both sets of notes, called O:
       1. search down from the (Px+O,Py) for O "on" pixels
       1. search right from the (Px, Py+O) for O "on" pixels
       1. if both of the above exist, then we have found the largest square with key corner P, save P and O, possibly in a max heap, or just keep them if O is larger than previous max
       1. break out of inner most loop(over the intersections of left and right notes.
#### Optimizations
of the following, only the precomputation of pixel runs impacts asymtotic runtime.
 * the scan through all pixels could use the current max Square to terminate sweeps early:
   - if the currentMax.Side > XDim - px, don't bother checking if there is a square, and just jump forward to the next row.
   - Similarly, if the currentMax.Side > YDim - py, just stop and return currentMax. 
 * when sweeping out from P, there is no reason to continue sweeping out in the second dimention once the first off cell is found
 * could pass the current max down to where we identify candidates, and refuse to evalulate any with side < maxSquare.Side 
 * a __substancial__ source of potentially repeated work is the sweeps to check whether a candidate square is complete. If the horizonal and vertical runs were pre-computeded, determining if a square was compete could be done by looking up the vertical run below the top right corner was larger than the candidate side length. a similar lookup calculation applies to the lower left corner. That removes 2 O(N) loops from the inner.
 
 ## Benchmarking
 The addition of some benchmarks demonstrates that the runtime grows much faster than the size of the inputs.
 ```shell
 $ go test -bench=Sqs -run=XXX
goos: darwin
goarch: amd64
pkg: github.com/smithhmark/maximum-box
BenchmarkBruteRandomSqs80_64-4    	     100	  19374205 ns/op
BenchmarkBruteRandomSqs80_128-4   	      10	 153761364 ns/op
BenchmarkBruteRandomSqs80_256-4   	       1	1230243028 ns/op
BenchmarkBruteRandomSqs80_512-4   	       1	10106609898 ns/op
PASS
ok  	github.com/smithhmark/maximum-box	15.011s
```
Each doubling of the input size results in approximately an 8X increase in duration.

 ## Profiling
```shell
 $ go test -bench=Sqs80_512 -run=XXX -cpuprofile=brute.prof
 
 $ go tool pprof brute_test.go brute.prof
```

### First round

Apparentely, The use of the singly linked list thrashes memory, and so there is a lot of allocation and deallocation. The largest chunk of time in largestSquareAt is spent by Push, and that time comes from allocation new Elements.
