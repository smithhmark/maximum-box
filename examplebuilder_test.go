package main

import "testing"

import "fmt"

func TestField(t *testing.T) {
	f := NewField(5,6, 0)

	for ii := 0; ii < 5 ; ii++ {
		for jj := 0; jj < 6 ; jj++ {
			if f.Field[jj][ii] != 0 {
				t.Fatalf("failed to initialize Field correctly")
			}
			if f.Get(ii, jj) != 0 {
				t.Fatalf("Get did something strange")
			}
		}
	}
}

func TestSquare(t *testing.T) {
	f := NewField(5,6, 0)
	f.PutSquare(0,0, 5, 1)
	for ii := 0 ; ii < 5 ; ii++ {
		sum := f.Get(ii,0) + f.Get(ii, 4)
		if sum != 2 {
			t.Fatalf("Failed to put fills where desired")
		}
		sum = f.Get(0,ii) + f.Get(4,ii)
		if sum != 2 {
			t.Fatalf("Failed to put fills where desired")
		}
		if f.Get(ii,5) != 0 {
			fmt.Print(f.Field[5])
			t.Fatalf("shouldn't have filled in bottom row")
		}
	}

	g := NewField(5,6, 0)
	s := SimpleSquare{0,0, 5}
	g.PutSSquare(&s, 1)
	for row := 0; row < g.YDim ; row++ {
		for col := 0 ; col < g.XDim ; col++ {
			if f.Field[row][col] != g.Field[row][col] {
				t.Fatalf("Discovered difference at (%d,%d)", col,row)
			}
		}
	}
	if g.CheckSSquare(&s, 1) != true {
		t.Fatalf("CheckSSquare failed to find what we know is there")
	}
}

func TestConcentric(t *testing.T) {
	f := NewField(100, 100, 0)
	sq1 := SimpleSquare{10,10, 80}
	sq2 := SimpleSquare{20,20, 60}
	sq3 := SimpleSquare{30,30, 40}
	sq4 := SimpleSquare{40,40, 20}

	Concentric(&f, 4)

	if !f.CheckSSquare(&sq1, 1) {
		t.Fatalf("Concentric didn't place %v", sq1)
	}
	if !f.CheckSSquare(&sq2, 1) {
		t.Fatalf("Concentric didn't place %v", sq2)
	}
	if !f.CheckSSquare(&sq3, 1) {
		t.Fatalf("Concentric didn't place %v", sq3)
	}
	if !f.CheckSSquare(&sq4, 1) {
		t.Fatalf("Concentric didn't place %v", sq4)
	}
}
