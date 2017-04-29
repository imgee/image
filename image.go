// Package image provides ...
package image

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	ext "github.com/imgee/image/extension"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	// "golang.org/x/image/webp"
)

var ErrUnsupported = errors.New("unsupported image type")

// encode image
func Encode(fm ext.Extension, w io.Writer, m image.Image) (err error) {
	switch fm {
	case ext.GIF:
		return gif.Encode(w, m, nil)
	case ext.JPG, ext.JPEG:
		return jpeg.Encode(w, m, nil)
	case ext.PNG:
		return png.Encode(w, m)
	case ext.TIFF:
		return tiff.Encode(w, m, nil)
	case ext.BMP:
		return bmp.Encode(w, m)
	}

	return ErrUnsupported
}
