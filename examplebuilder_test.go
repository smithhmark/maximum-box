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
}

