package main

import (
	sl "github.com/smithhmark/gomnibus/singlelink"
	"container/heap"
	//"log"
	//"fmt"
)

type OrderedSquares []*SimpleSquare;

func (h OrderedSquares) Len() int {
	return len(h)
}
func (h OrderedSquares) Less(i, j int) bool {
	return h[i].Side < h[j].Side
}
func (h OrderedSquares) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *OrderedSquares) Push(x interface{}) {
	*h = append(*h, x.(*SimpleSquare))
}
func (h *OrderedSquares) Pop() interface{} {
	old := *h
	n := len(old) -1
	x := old[n]
	*h = old[:n]
	return x
}

func addSquare(h *OrderedSquares, sqr *SimpleSquare, n int) {
	if sqr.Side > 0 {
		if h.Len() > 0 {
			//log.Printf("found valid square:%v", *sqr)
			minOfMax := (*h)[0]
			if h.Len() < n {
				//log.Printf("\tnot yet up to N")
				heap.Push(h, sqr)
			} else if sqr.Side > minOfMax.Side {
				//log.Printf("\tlarger than Nth:%v", *(*h)[0])
				heap.Pop(h)
				heap.Push(h, sqr)
			//} else {
				//log.Printf("\tNot big enough")
			}
		} else {
			//log.Printf("found First valid square:%v", *sqr)
			heap.Push(h, sqr)
		}
	}
}

func BruteN(f *Field, n int) []*SimpleSquare {
	maxSqs := make(OrderedSquares, 0, n)
	heap.Init(&maxSqs)
	cfg := config(f, false)
	for Py := 0 ; Py < f.YDim ; Py++ {
		if len(maxSqs) > 0 && Py > f.YDim - maxSqs[0].Side {
			continue
		}
		for Px := 0 ; Px < f.YDim ; Px++ {
			if len(maxSqs) > 0 && Px > f.XDim - maxSqs[0].Side {
				continue
			}
			if f.Get(Px, Py) == 1 {
				largestS := largestSquareAt(f, Px,Py, cfg)
				addSquare(&maxSqs, &largestS, n)
			}
		}
	}

	ret := make([]*SimpleSquare, maxSqs.Len())
	idx := maxSqs.Len() - 1
	for maxSqs.Len() > 0 {
		tmp := heap.Pop(&maxSqs).(*SimpleSquare)
		ret[idx] =  tmp
		//fmt.Printf("Popped:%v, placed at:%d\n", tmp, idx)
		idx--
	}
	return ret
}

func completions(f *Field) ([][]int, [][]int) {
	horts := make([][]int, f.XDim)
	verts := make([][]int, f.XDim)
	for col := 0 ; col < f.XDim ; col++ {
		verts[col] = make([]int,f.YDim)
		horts[col] = make([]int,f.YDim)
	}
	for row := f.YDim -1 ; row >= 0 ; row-- {
		for col := f.XDim -1 ; col >= 0 ; col-- {
			val := f.Get(row, col)
			if val != 0 {
				if col == f.XDim - 1 {
					horts[col][row] = 1
				} else {
					tmp := horts[col+1][row]
					horts[col][row] = tmp + 1
				}
				if row == f.YDim -1 {
					verts[col][row] = 1
				} else {
					tmp := verts[col][row+1]
					verts[col][row] = tmp + 1
				}
			}
		}
	}
	return horts, verts
}

type largestSquareData struct {
	xCandidates []int
	yCandidates []int
	mergedCandidates []int
	horts, verts [][]int
}

func config(f *Field, useLookups bool) *largestSquareData {
	xCandidates := make([]int, f.XDim)
	yCandidates := make([]int, f.YDim)
	var mergedCandidates []int
	if f.XDim > f.YDim {
		mergedCandidates = make([]int, 0, f.XDim)
	} else {
		mergedCandidates = make([]int, 0, f.YDim)
	}
	var horts, verts [][]int
	if useLookups {
		horts, verts = completions(f)
	} else {
		horts, verts = make([][]int,0), make([][]int,0)
	}
	conf := largestSquareData{xCandidates,yCandidates,mergedCandidates, horts,verts}
	return &conf
}

func Brute(f *Field) *SimpleSquare {
	var maxS *SimpleSquare
	//log.Printf("Completed Setup")
	cfg := config(f, false)
	for Py := 0 ; Py < f.YDim ; Py++ {
		if maxS != nil && Py > f.YDim - maxS.Side {
			continue
		}
		for Px := 0 ; Px < f.YDim ; Px++ {
			if maxS != nil && Px > f.XDim - maxS.Side {
				continue
			}
			if f.Get(Px, Py) == 1 {
				largestS := largestSquareAt(f, Px,Py, cfg)
				if largestS.Side > 0 {
					if maxS == nil {
						maxS = &largestS
					} else if largestS.Side > maxS.Side {
						//log.Printf("Found new largest square:%v", largestS)
						maxS = &largestS
					} 
				}
			}
		}
	}
	return maxS
}

func largestSquareAt(f *Field, px, py int, cfg *largestSquareData) SimpleSquare {
	ii := 1
	xCandidates := cfg.xCandidates
	yCandidates := cfg.yCandidates
	mergedCandidates := cfg.mergedCandidates
	horts := cfg.horts
	verts := cfg.verts
	curXCand, curYCand := -1, -1
	//log.Printf("Walking horizontally looking for possible side lengths")
	for px+ii <f.XDim && f.Get(px+ii, py) == 1 && py+1 < f.YDim{
		if f.Get(px+ii, py+1) == 1 {
			curXCand++
			xCandidates[curXCand] = ii
		}
		ii++
	}

	ii = 1
	//log.Printf("Walking vertically looking for possible side lengths")
	for py+ii < f.YDim && f.Get(px, py+ii) == 1 && px+1 < f.XDim{
		if f.Get(px+1, py+ii) == 1 {
			curYCand++
			yCandidates[curYCand] = ii
		}
		ii++
	}

	candidates := intersectCandidates(xCandidates, yCandidates, curXCand, curYCand, mergedCandidates)
	//log.Printf("candidate sizes:%v", candidates)

	for _, side := range candidates {
		if completesSquare(f, px,py, side, horts, verts){
			//log.Printf("found a square:%v:%v:%v", px,py,side)
			return SimpleSquare{px,py,side+1}
		}
	}
	//log.Printf("Did not find a square")
	return SimpleSquare{0,0,0}
}

func completesSquare(f *Field, px,py, side int, horts, verts [][]int) bool {
	if len(horts) > 0 {
		if horts[px][py +side] >= side && verts[px+side][py] >= side {
			return true
		} else {
			return false
		}
	} else {
		for ii := 0 ; ii < side ; ii++ {
			rowVal := f.Get(px+ii, py+side)
			colVal := f.Get(px+side, py+ii)
			sum := rowVal + colVal
			//sum := f.Get(px+ii, py+side-1) + f.Get(px+side-1, py+ii)
			if sum != 2 {
				return false
			}
		}
		return true
	}
}

func pop(s *sl.Stack) int {
	if s.Size() ==0 {
		return -1
	}
	tmp, err := s.Pop()
	if err != nil {
		return -1
	}
	return tmp.(int)
}

func intersectCandidates(s1, s2 []int, cx, cy int, merged []int) []int {
	/*var merged []int
	if len(s1) > len(s2) {
		merged = make([]int, 0, len(s1))
	}else {
		merged = make([]int, 0, len(s1))
	}
	*/
	if cx == -1 || cy == -1 {
		return merged
	}

	var v1 int
	var v2 int
	for cx >= 0 && cy >= 0 {
		v1 = s1[cx]
		v2 = s2[cy]
		if v1 > v2 {
			cx--
		} else if v2 > v1 {
			cy--
		} else {
			merged = append(merged, v1)
			cx--
			cy--
		}
	}
	return merged
}

