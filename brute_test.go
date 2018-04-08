package main

import "testing"
//import "fmt"

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

