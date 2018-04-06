package main

import "log"

var EXAMPLE1 = [][]uint8{
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
}


type Field struct {
	XDim, YDim int
	Field [][]uint8
}
func NewField(xdim, ydim int, fill uint8) Field {
	f := make([][]uint8, ydim)
	for row := 0 ; row < ydim ; row++ {
		f[row] = make([]uint8, xdim)
		for ii := 0 ; ii < xdim ; ii++ {
			f[row][ii] = fill
		}
	}
	return Field{xdim, ydim, f}
}

func (f *Field) Set(x, y int, v uint8) {
	if x >= f.XDim || y >= f.YDim {
		log.Printf("Set(%v,%v)", x, y)
		return // error
	}
	f.Field[y][x] = v
}

func (f *Field) Get(x, y int) (v uint8) {
	if x > f.XDim || y > f.YDim {
		return // error
	}
	v = f.Field[y][x]
	return
}

func (f *Field) PutSquare(x,y, side int, fill uint8) {
	if x + side > f.XDim {
		log.Printf("Square too wide")
		return
	}
	if y + side > f.YDim {
		log.Printf("Square too tall")
		return
	}
	for ii := 0 ; ii < side ; ii++ {
		f.Set(x+ii, y, fill)
		f.Set(x+ii, y + side - 1, fill)
		f.Set(x, y+ii, fill)
		f.Set(x+side-1, y+ii, fill)
	}
}

