package images

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/pkg/errors"

	"github.com/nfnt/resize"
)

// ResizeImage takes in the source path of an image, resizes it to the dimensions requested, and copies it to the destination
func ResizeImage(srcPath string, destPath string, format string, width uint, height uint) error {

	fmt.Println("Resizing image")
	file, err := os.Open(srcPath)
	if err != nil {
		return errors.Wrap(err, "cannot open original image")
	}

	if format == "png" {
		img, err := png.Decode(file)
		if err != nil {
			return errors.Wrapf(err, "%s is not a PNG file", srcPath)
		}
		file.Close()
		m := resize.Resize(width, height, img, resize.Lanczos3)
		out, err := os.Create(destPath)
		if err != nil {
			return errors.Wrapf(err, "cannot create destination path %s", destPath)
		}
		defer out.Close()
		png.Encode(out, m)
	} else {
		img, err := jpeg.Decode(file)
		if err != nil {
			return errors.Wrapf(err, "%s is not a JPEG file", srcPath)
		}
		file.Close()
		m := resize.Resize(width, height, img, resize.Lanczos3)
		out, err := os.Create(destPath)
		if err != nil {
			return errors.Wrapf(err, "cannot create destination path %s", destPath)
		}
		defer out.Close()
		jpeg.Encode(out, m, nil)
	}
	return nil
}
