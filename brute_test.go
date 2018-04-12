package main

import "testing"
//import "fmt"

import "container/heap"

func TestBruteEmpty(t *testing.T) {
	input := NewField(10,10, 0)

	max := Brute(&input)
	if max != nil {
		t.Fatalf("Should have found no solution")
	}
}

func TestBruteSingle(t *testing.T) {
	input := NewField(10,10, 0)
	sq := SimpleSquare{1,1,5}
	input.PutSSquare(&sq, 1)
	//fmt.Print(input.Stringify())
	max := Brute(&input)
	if max == nil {
		t.Fatalf("Should have found a square")
	}
	//fmt.Printf("%v\n", max)
	if *max != sq {
		t.Fatalf("Failed to find the expected square:%v", sq)
	}
}

func TestBruteAbutEnd(t *testing.T) {
	input := NewField(10,10, 0)
	input.PutSquare(5,5, 5, 1)
	//fmt.Print(input.Stringify())
	max := Brute(&input)
	if max == nil {
		t.Fatalf("Should have found a square")
	}
	//fmt.Printf("%v\n", max)
	if max.Side != 5 || max.X != 5 || max.Y != 5 {
		t.Fatalf("Failed to find the expected square at 5,5 side=5")
	}
}

func TestBruteIntersecting(t *testing.T) {
	input := NewField(10,10, 0)
	small := SimpleSquare{1,1, 5}
	large := SimpleSquare{3,3, 6}
	input.PutSSquare(&small, 1)
	input.PutSSquare(&large, 1)
	//fmt.Print(input.Stringify())
	max := Brute(&input)
	if max == nil {
		t.Fatalf("Should have found a square")
	}
	//fmt.Printf("%v\n", max)
	if *max != large {
		t.Fatalf("Failed to find the expected square: %v", large)
	}
}

func TestConcentricSquares(t *testing.T) {
	f := NewField(100,100, 0)
	big := SimpleSquare{ 10,10, 80}
	small := SimpleSquare{20,20, 60}
	f.PutSSquare(&big, 1)
	f.PutSSquare(&small, 1)

	sq := Brute(&f)

	if *sq != big {
		t.Fatalf("Should have found %v", big)
	}
}

func noTestBruteN1(t *testing.T) {
	input := NewField(10,10, 0)
	small := SimpleSquare{1,1, 5}
	large := SimpleSquare{3,3, 6}
	input.PutSSquare(&small, 1)
	input.PutSSquare(&large, 1)

	res := BruteN(&input, 1)
	if len(res) != 1 || *res[0] != large {
		t.Fatalf("failed to find largest square")
	}
}

func TestHeap2ItemsInOrder(t *testing.T) {
	h := make(OrderedSquares, 0, 2)
	small := SimpleSquare{1,1, 5}
	large := SimpleSquare{3,3, 6}
	heap.Push(&h, &small)
	heap.Push(&h, &large)

	tmp := heap.Pop(&h).(*SimpleSquare)
	if *tmp != small {
		t.Fatalf("First pop'ed item should be %v, got %v", small, tmp)
	}
	tmp = heap.Pop(&h).(*SimpleSquare)
	if *tmp != large {
		t.Fatalf("Second pop'ed item should be %v, got %v", large, tmp)
	}
}

func TestHeap(t *testing.T) {
	h := make(OrderedSquares, 0, 5)

	if cap(h) != 5 || len(h) != 0 {
		t.Fatalf("Failed to build heap correctly")
	}

	tiny := SimpleSquare{2,2, 3}
	small := SimpleSquare{1,1, 5}
	medium := SimpleSquare{3,3, 6}
	large := SimpleSquare{4,4, 10}

	heap.Push(&h, &large)
	if cap(h) != 5 || len(h) != 1 {
		t.Fatalf("Failed to build heap.Push correctly")
	}
	if *h[0] != large {
		t.Fatalf("Min item in heap should be %v", large)
	}

	heap.Push(&h, &small)
	if *h[0] != small {
		t.Fatalf("Min item in heap should be %v", small)
	}

	heap.Push(&h, &medium)
	if *h[0] != small {
		t.Fatalf("Min item in heap should be %v", small)
	}
	heap.Push(&h, &tiny)
	if *h[0] != tiny {
		t.Fatalf("Min item in heap should be %v", tiny)
	}

	tmp := heap.Pop(&h).(*SimpleSquare)
	if *tmp != tiny {
		t.Fatalf("First pop'ed item should be %v, got %v", tiny, tmp)
	}
	tmp = heap.Pop(&h).(*SimpleSquare)
	if *tmp != small {
		t.Fatalf("Second pop'ed item should be %v, got %v", small, tmp)
	}
	tmp = heap.Pop(&h).(*SimpleSquare)
	if *tmp != medium {
		t.Fatalf("Third pop'ed item should be %v, got %v", medium, tmp)
	}
	tmp = heap.Pop(&h).(*SimpleSquare)
	if *tmp != large {
		t.Fatalf("Forth pop'ed item should be %v, got %v", large, tmp)
	}
}

func TestBruteN2(t *testing.T) {
	input := NewField(10,10, 0)
	small := SimpleSquare{1,1, 5}
	large := SimpleSquare{3,3, 6}
	input.PutSSquare(&small, 1)
	input.PutSSquare(&large, 1)

	res := BruteN(&input, 2)
	if len(res) != 2 {
		t.Fatalf("expected to find two squares")
	}
	if *res[0] != large {
		//for _, v := range res {
		//	fmt.Printf("growl:%v\n", *v)
		//}
		t.Fatalf("failed to find %v, instead:%v", large, *res[0])
	}
	if *res[1] != small {
		t.Fatalf("failed to find %v, instead:%v", small, *res[1])
	}
}

func TestBruteNSpectrum(t *testing.T) {
	f := NewField(1000,1000, 0)
	DecreasingSpectrum(&f, 8, 300, 50)

	sqs := BruteN(&f, 5)

	if len(sqs) != 5 {
		
		t.Fatalf("Didn't find all the squares, found:%v", sqs)
	}
	for ii, vv := range sqs {
		if ii > 0 && vv.Side > sqs[0].Side {
			for gg:=0;gg<len(sqs);gg++ {
				t.Log(sqs[gg])
			}
			t.Fatalf("heap property violated at %v, v=%v", ii, vv)
		}
	}
}

var result *SimpleSquare

func benchBruteRandSq(size,number int, b *testing.B){
	f := NewField(size,size, 0)
	for ii := 0; ii < number; ii++ {
		f.RandomSquare()
	}
	var r *SimpleSquare

	b.ResetTimer()
	for ii := 0 ; ii < b.N ; ii++ {
		r = Brute(&f)
	}
	result = r
}
func benchBruteRand(size int, prob float64, b *testing.B){
	f := NewField(size,size, 0)
	f.RandomPoints(prob)
	var r *SimpleSquare

	b.ResetTimer()
	for ii := 0 ; ii < b.N ; ii++ {
		r = Brute(&f)
	}
	result = r
}

func BenchmarkBruteRandomPts90(b *testing.B) {benchBruteRand(1024,.9,b) }
func BenchmarkBruteRandomPts88(b *testing.B) {benchBruteRand(1024,.88,b) }
func BenchmarkBruteRandomPts85(b *testing.B) {benchBruteRand(1024,.85,b) }
func BenchmarkBruteRandomPts80(b *testing.B) {benchBruteRand(1024,.80,b) }

func BenchmarkBruteRandomPts80_256(b *testing.B) {benchBruteRand(256,.80,b) }
func BenchmarkBruteRandomPts80_512(b *testing.B) {benchBruteRand(512,.80,b) }
func BenchmarkBruteRandomPts80_1024(b *testing.B) {benchBruteRand(1024,.80,b) }
func BenchmarkBruteRandomPts80_2028(b *testing.B) {benchBruteRand(2028,.80,b) }

func BenchmarkBruteRandomSqs80_64(b *testing.B) {benchBruteRand(64,80,b) }
func BenchmarkBruteRandomSqs80_128(b *testing.B) {benchBruteRand(128,80,b) }
func BenchmarkBruteRandomSqs80_256(b *testing.B) {benchBruteRand(256,80,b) }
func BenchmarkBruteRandomSqs80_512(b *testing.B) {benchBruteRand(512,80,b) }
//func BenchmarkBruteRandomSqs80_1024(b *testing.B) {benchBruteRand(1024,80,b) }
//func BenchmarkBruteRandomSqs80_2028(b *testing.B) {benchBruteRand(2028,80,b) }
