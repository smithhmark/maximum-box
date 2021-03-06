package main

import (
	"image"
	"image/png"
	"image/color"
	"os"
	"log"
	"flag"
	"fmt"
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
	log.Printf("seting %dth largest square%v to colorvalue:%d", ith, *ss, cv)
	f.PutSSquare(ss, cv)
}

func main() {
	var width = flag.Int("width", 1024, "the width of the field")
	var height = flag.Int("height", 1024, "the height of the field")
	var pattern = flag.String("pattern", "DS", "Pattern {DS (DecreasingSpectrum) | C (Concentric) | RS (RandomSquares) | RP (RandomPoints}")
	var largest = flag.Int("largest", 512, "side length of largest square in pattern")
	var smallest = flag.Int("smallest", 32, "side length of smallest square in pattern")
	var number = flag.Int("number", 10, "number of squares to put into the pattern")
	var opath = flag.String("path", "image.png", "path to place output image")
	var output = flag.String("output", "image", "Outout {image | text}")
	var prob = flag.Float64("prob", .5, "Probably of a pixel being active when building a randompoint pattern")

	flag.Parse()

	f := NewField(*width,*height, 0)
	//f.PutSquare( 10,10, 80, 1)
	//f.PutSquare( 20,20, 60, 1)
	switch *pattern {
	case "DS" :
		DecreasingSpectrum(&f, *number, *largest, *smallest)
	case "C":
		Concentric(&f, *number)
	case "RS":
		for ii := 0; ii < *number; ii++ {
			f.RandomSquare()
		}
	case "RP":
		f.RandomPoints(*prob)
	default:
		Concentric(&f, 40)
	}

	sqs := BruteN(&f,len(colors))
	//sqs := BruteN(&f,5)
	switch *output {
	case "text":
		fmt.Printf("Found: %d", len(sqs))
		fmt.Printf("\tX\tY\tSide\n")
		for _, vv := range sqs {
			fmt.Printf("\t%d\t%d\t%d\n", vv.X, vv.Y, vv.Side)
		}
	case "image":
		sf := makeStandardizedField(&f)
		for ii, vv := range sqs {
			highlightSquare(&sf, vv, ii)
		}
		makeImageFile(&sf, *opath)
	}

}
