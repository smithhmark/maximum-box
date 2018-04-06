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
