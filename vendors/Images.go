package vendors

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// ResizeImage ..
func ResizeImage(imageName string, outputName string, path string, format string) {
	file, err := os.Open(path + imageName)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg or png into image.Image
	var img image.Image
	if format == "png" {
		img, err = png.Decode(file)
	} else {
		img, err = jpeg.Decode(file)
	}
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(600, 600, img, resize.Lanczos2)

	out, err := os.Create(path + outputName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	if format == "png" {
		png.Encode(out, m)
	} else {
		jpeg.Encode(out, m, nil)
	}

	// Delete Old Image
	e := os.Remove(path + imageName)
	if e != nil {
		log.Fatal(e)
	}
}

func guessImageFormat(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	return
}
