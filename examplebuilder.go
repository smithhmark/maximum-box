package main

import "log"
import "bytes"
import "fmt"

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

type SimpleSquare struct {
	X,Y, Side int
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
	if x >= f.XDim || y >= f.YDim {
		panic(fmt.Sprintf("Trying to Get at (%d,%d) when bounds <%d,%d>", x,y, f.XDim, f.YDim))
		return // error
	}
	v = f.Field[y][x]
	return
}

func (f *Field) PutSSquare(s SimpleSquare, fill uint8) {
	f.PutSquare(s.X, s.Y, s.Side, fill)
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

func (f *Field) Stringify() string {
	var buffer bytes.Buffer

	for row := 0 ; row < f.YDim ; row++ {
		for col := 0 ; col < f.XDim ; col++ {
			if f.Field[row][col] == 0 {
				buffer.WriteString("0")
			} else {
				buffer.WriteString("1")
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
