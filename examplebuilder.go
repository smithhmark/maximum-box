package main

import (
	"log"
	"bytes"
	"fmt"
	"math/rand"
)

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
	Fill uint8
}

func NewField(xdim, ydim int, fill uint8) Field {
	f := make([][]uint8, ydim)
	for row := 0 ; row < ydim ; row++ {
		f[row] = make([]uint8, xdim)
		for ii := 0 ; ii < xdim ; ii++ {
			f[row][ii] = fill
		}
	}
	return Field{xdim, ydim, f, fill}
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

func (f *Field) PutSSquare(s *SimpleSquare, fill uint8) {
	f.PutSquare(s.X, s.Y, s.Side, fill)
}

func (f *Field) RandomPoints(prob float64) {
	for row := 0 ; row < f.YDim ; row++ {
		for col := 0 ; col < f.XDim ; col++ {
			if rand.Float64() < prob{
				f.Set(row, col, 1)
			}
		}
	}
}

func (f *Field) RandomSquare() {
	var side int
	if f.XDim < f.YDim {
		side = rand.Intn(f.XDim)
	} else {
		side = rand.Intn(f.YDim)
	}

	x := rand.Intn(f.XDim - side)
	y := rand.Intn(f.YDim - side)

	f.PutSquare(x,y,side, 1)
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

func (f *Field) CheckSSquare(ss *SimpleSquare, fill uint8) bool {
	for ii := 0 ; ii < ss.Side ; ii++ {
		if  x, y := ss.X+ii, ss.Y;f.Get(x,y) != fill {
			log.Printf("CheckSSquare %v failed @(%d,%d)", *ss, x,y) 
			return false
		}
		if x, y := ss.X+ii, ss.Y + ss.Side - 1; f.Get(x,y) != fill {
			log.Printf("CheckSSquare %v failed @(%d,%d)", *ss, x,y) 
			return false
		}
		if x, y := ss.X, ss.Y+ii; f.Get(x,y) != fill {
			log.Printf("CheckSSquare %v failed @(%d,%d)", *ss, x,y) 
			return false
		}
		if x, y := ss.X+ss.Side-1, ss.Y+ii; f.Get(x,y) != fill {
			log.Printf("CheckSSquare %v failed @(%d,%d)", *ss, x,y) 
			return false
		}
	}
	return true
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


func Concentric(f *Field, num int) {
	var offset int

	centerX := f.XDim / 2
	centerY := f.YDim / 2

	if f.XDim < f.YDim {
		offset = f.XDim / ((1+num) *2)
	} else {
		offset = f.YDim / ((1+num) *2)
	}

	for ii := 1 ; ii <= num ; ii++ {
		x := centerX - offset * ii
		y := centerY - offset * ii
		sq := SimpleSquare{x,y, 2*offset*ii}
		log.Printf("concentric sq:%d at %v", ii, sq)
		f.PutSSquare(&sq, 1)
	}
}

func DecreasingSpectrum(f *Field, num, largest, smallest int) {
	centerXs := make([]int,num)
	xOffset := f.XDim /(num +1)
	centerYs := make([]int,num)
	yOffset := f.YDim /(num +1)
	for ii := 1 ; ii <= num ; ii++ {
		centerXs[ii-1] = xOffset*ii
		centerYs[ii-1] = yOffset*ii
	}

	sizeStep := (largest - smallest) /num
	for ii := 0 ; ii < num ; ii++ {
		sqSz := largest - sizeStep*ii
		x := centerXs[ii] - sqSz /2
		y := centerYs[ii] - sqSz /2
		if x < 0 {
			sqSz += x*2
			x = centerXs[ii] - sqSz /2
			y = centerYs[ii] - sqSz /2
		}
		if y < 0 {
			sqSz += y*2
			x = centerXs[ii] - sqSz /2
			y = centerYs[ii] - sqSz /2
		}
		log.Printf("DecreasingSpectrum placing sq:{%d, %d, %d}", x,y,sqSz)
		f.PutSquare(x,y,sqSz, 1)
	}
}
