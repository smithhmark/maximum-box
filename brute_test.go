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
	input.PutSquare(1,1, 5, 1)
	//fmt.Print(input.Stringify())
	max := Brute(&input)
	if max == nil {
		t.Fatalf("Should have found a square")
	}
	//fmt.Printf("%v\n", max)
	if max.Side != 5 || max.X != 1 || max.Y != 1 {
		t.Fatalf("Failed to find the expected square at 1,1 side=5")
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
	input.PutSquare(1,1, 5, 1)
	input.PutSquare(3,3, 6, 1)
	//fmt.Print(input.Stringify())
	max := Brute(&input)
	if max == nil {
		t.Fatalf("Should have found a square")
	}
	//fmt.Printf("%v\n", max)
	if max.Side != 6 || max.X != 3 || max.Y != 3 {
		t.Fatalf("Failed to find the expected square at 3,3 side=6")
	}
}

