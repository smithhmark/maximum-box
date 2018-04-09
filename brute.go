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
	for Py := 0 ; Py < f.YDim ; Py++ {
		for Px := 0 ; Px < f.YDim ; Px++ {
			if f.Get(Px, Py) == 1 {
				largestS := largestSquareAt(f, Px,Py)
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

func Brute(f *Field) *SimpleSquare {
	var maxS *SimpleSquare
	for Py := 0 ; Py < f.YDim ; Py++ {
		for Px := 0 ; Px < f.YDim ; Px++ {
			if f.Get(Px, Py) == 1 {
				largestS := largestSquareAt(f, Px,Py)
				if largestS.Side > 0 {
					if maxS == nil {
						maxS = &largestS
					} else if largestS.Side > maxS.Side {
						maxS = &largestS
					} 
				}
			}
		}
	}
	return maxS
}

func largestSquareAt(f *Field, px, py int) SimpleSquare {
	ii := 1
	xCandidates := sl.NewStack()
	//log.Printf("Walking horizontally looking for possible side lengths")
	for px+ii <f.XDim && f.Get(px+ii, py) == 1 && py+1 < f.YDim{
		if f.Get(px+ii, py+1) == 1 {
			xCandidates.Push(ii)
		}
		ii++
	}

	ii = 1
	//log.Printf("Walking vertically looking for possible side lengths")
	yCandidates := sl.NewStack()
	for py+ii < f.YDim && f.Get(px, py+ii) == 1 && px+1 < f.XDim{
		if f.Get(px+1, py+ii) == 1 {
			yCandidates.Push(ii)
		}
		ii++
	}

	candidates := intersectCandidates(xCandidates, yCandidates)
	//log.Printf("candidate sizes:%v", candidates)

	for _, side := range candidates {
		if completesSquare(f, px,py, side){
			//log.Printf("found a square:%v:%v:%v", px,py,side)
			return SimpleSquare{px,py,side+1}
		}
	}
	//log.Printf("Did not find a square")
	return SimpleSquare{0,0,0}
}

func completesSquare(f *Field, px,py, side int) bool {
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

func intersectCandidates(s1, s2 *sl.Stack) []int {
	merged := make([]int, 0, 2)
	if s1.Size() == 0 || s2.Size() == 0 {
		return merged
	}
	v1 := pop(s1)
	v2 := pop(s2)

	for v1 > 0 && v2 > 0 {
		if v1 > v2 {
			v1 = pop(s1)
		} else if v2 > v1 {
			v2 = pop(s2)
		} else {
			if len(merged) == cap(merged) {
				newmerged := make([]int, len(merged), 2*cap(merged))
				copy(newmerged, merged)
				merged = newmerged
			}
			merged = append(merged, v1)
			v1 = pop(s1)
			v2 = pop(s2)
		}
	}
	return merged
}



