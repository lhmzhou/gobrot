package main

import (
	"errors"
	"flag"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/lhmzhou/gobrot/gobrot"
)

var (
	output  = flag.String("out", "mandelbrot.jpeg", "name of the output image file")
	height  = flag.Int("h", 2048, "height of the output image in pixels")
	width   = flag.Int("w", 2048, "width of the output image in pixels")
	mode    = flag.String("mode", "seq", "mode: seq, px, row, workers")
	workers = flag.Int("workers", 1, "number of workers to use")
)

func main() {
	flag.Parse()

	// open a new file
	f, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}

	// create the image
	img := gobrot.Create(*height, *width, gobrot.Mode(*mode), *workers)

	// and encoding it
	fmt := filepath.Ext(*output)
	switch fmt {
	case ".png":
		err = png.Encode(f, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(f, img, nil)
	case ".gif":
		err = gif.Encode(f, img, nil)
	default:
		err = errors.New("unkwnown format " + fmt)
	}
	// unless you can't
	if err != nil {
		log.Fatal(err)
	}
}
