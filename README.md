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
 * the scan through all pixels could use the current max Square to terminate sweeps early:
   - if the currentMax.Side > XDim - px, don't bother checking if there is a square, and just jump forward to the next row.
   - Similarly, if the currentMax.Side > YDim - py, just stop and return currentMax. 
 * when sweeping out from P, there is no reason to continue sweeping out in the second dimention once the first off cell is found
 * could pass the current max down to where we identify candidates, and refuse to evalulate any with side < maxSquare.Side 
