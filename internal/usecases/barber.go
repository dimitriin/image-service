package usecases

import (
	"github.com/dimitriin/image-service/internal/domain"
	"github.com/nxshock/colorcrop"
	"image"
	"image/color"
	"image/draw"
)

// ImageCutter cuts borders around image.
type ImageCutter interface {
	Cut(img domain.CleanableImage) (domain.CleanableImage, error)
}

// Barber is ImageCutter realization which uses https://github.com/nxshock/colorcrop library
type Barber struct {
	color color.RGBA
	threshold float64
}

// NewBarber is constructor for Barber struct
func NewBarber(color color.RGBA, threshold float64) *Barber {
	return &Barber{color: color, threshold: threshold}
}

// Cut remove border around image using https://github.com/nxshock/colorcrop library
func (b Barber) Cut(img domain.CleanableImage) (domain.CleanableImage, error) {

	drawableImg := image.NewRGBA(img.Bounds())
	draw.Draw(drawableImg, img.Bounds(), img, image.Point{}, draw.Src)

	croppedImage := colorcrop.Crop(drawableImg, b.color, b.threshold)
	return domain.NewImageFromStd(croppedImage)
}

