package main

import (
	"image"
	"image/png"
	"image/color"
	"os"
	"log"
)

var colors = []color.Color {
	color.NRGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	},
	color.NRGBA{
		R: 255,
		G: 128,
		B: 0,
		A: 255,
	},
	color.NRGBA{
		R: 255,
		G: 255,
		B: 0,
		A: 255,
	},
	color.NRGBA{
		R: 128,
		G: 255,
		B: 0,
		A: 255,
	},
	color.NRGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	},
	color.NRGBA{
		R: 0,
		G: 255,
		B: 128,
		A: 255,
	},
	color.NRGBA{
		R: 0,
		G: 255,
		B: 255,
		A: 255,
	},
	color.NRGBA{
		R: 0,
		G: 128,
		B: 255,
		A: 255,
	},
	color.NRGBA{
		R: 0,
		G: 0, 
		B: 255,
		A: 255,
	},
	color.NRGBA{
		R: 128,
		G: 0, 
		B: 255,
		A: 255,
	},
	color.NRGBA{
		R: 255,
		G: 0, 
		B: 255,
		A: 255,
	},
	color.NRGBA{
		R: 255,
		G: 0, 
		B: 128,
		A: 255,
	},
}

func getColor(cv uint8) color.Color {
	if tmp := 255 - int(cv); tmp > len(colors) {
		return colors[len(colors)-1]
	} else {
		return colors[tmp]
	}
}

func makeImageFile(f *Field, path string) {
	img := image.NewNRGBA(image.Rect(0, 0, f.XDim, f.YDim))

	for y := 0; y < f.YDim; y++ {
		for x := 0; x < f.XDim; x++ {
			v := f.Get(x,y)
			if v == 0 {
				img.Set(x, y, color.Black)
			} else if v == 1 {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, getColor(v))
			}
		}
	}

	fd, err := os.Create(path)
	if err != nil {
		log.Fatal("Couldn't create file:", path)
	}

	if err := png.Encode(fd, img); err != nil {
		fd.Close()
		log.Fatal(err)
	}

	if err := fd.Close(); err != nil {
		log.Fatal(err)
	}
}

func makeStandardizedField(f *Field) (nf Field){
	nf = NewField(f.XDim, f.YDim, 0)

	for yy := 0 ; yy < f.YDim; yy++ {
		for xx := 0 ; xx < f.XDim; xx++ {
			if tmp := f.Get(xx, yy); tmp != f.Fill {
				nf.Set(xx,yy, 1)
			} else {
				nf.Set(xx,yy, 0)
			}
		}
	}
	return
}

func highlightSquare(f *Field, ss *SimpleSquare, ith int) {
	if ith > len(colors) {
		log.Fatal("Can't set square to be the ", ith, "largest, max is:", len(colors))
	}

	cv := 255 - uint8(ith)
	log.Printf("seting %dth largest square to colorvalue:%d", ith, cv)
	f.PutSSquare(ss, cv)
}

func main() {
	f := NewField(1000,1000, 0)
	//f.PutSquare( 10,10, 80, 1)
	//f.PutSquare( 20,20, 60, 1)
	//Concentric(&f, 4)
	DecreasingSpectrum(&f, 8, 300, 50)

	sqs := BruteN(&f,len(colors))
	//sqs := BruteN(&f,5)
	log.Printf("Found: %d", len(sqs))
	for _, vv := range sqs {
		log.Printf("\t%v", vv)
	}

	sf := makeStandardizedField(&f)
	for ii, vv := range sqs {
		highlightSquare(&sf, vv, ii)
	}

	makeImageFile(&sf, "test.png")
}
