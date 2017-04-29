// Package extension provides ...
package extension

// image extension
type Extension string

// format enum
var (
	PNG  Extension = "png"
	BMP  Extension = "bmp"
	IFF  Extension = "iff"
	TIFF Extension = "tiff"
	PNF  Extension = "pnf"
	GIF  Extension = "gif"
	JPG  Extension = "jpg"
	JPEG Extension = "jpeg"
	MNG  Extension = "mng"
	PSD  Extension = "psd"
	SAI  Extension = "sai"
	UFO  Extension = "ufo"
	XCF  Extension = "xcf"
	PCX  Extension = "pcx"
	PPM  Extension = "ppm"
	WEBP Extension = "webp"

	Extensions = map[string]Extension{
		"png":  PNG,
		"bmp":  BMP,
		"iff":  IFF,
		"tiff": TIFF,
		"pnf":  PNF,
		"gif":  GIF,
		"jpg":  JPG,
		"jpeg": JPEG,
		"mng":  MNG,
		"psd":  PSD,
		"sai":  SAI,
		"ufo":  UFO,
		"xcf":  XCF,
		"pcx":  PCX,
		"ppm":  PPM,
		"webp": WEBP,
	}
)
