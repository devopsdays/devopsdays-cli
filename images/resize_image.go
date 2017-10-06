package images

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// ResizeImage takes in the source path of an image, resizes it to the dimensions requested, and copies it to the destination
func ResizeImage(srcPath string, destPath string, format string, width uint, height uint) {

	fmt.Println("Resizing image")
	file, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	if format == "png" {
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		m := resize.Resize(width, height, img, resize.Lanczos3)
		out, err := os.Create(destPath)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		png.Encode(out, m)
	} else {
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		m := resize.Resize(width, height, img, resize.Lanczos3)
		out, err := os.Create(destPath)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		jpeg.Encode(out, m, nil)
	}

}
