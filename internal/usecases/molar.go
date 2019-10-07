package usecases

import (
	"github.com/dimitriin/image-service/internal/domain"
	"image"
	"image/color"
	"image/draw"
)

// ObjectReplacer represents interface for replacing object from images.
type ObjectReplacer interface {
	Replace(img domain.CleanableImage, objects []domain.CleanableObject) (domain.CleanableImage, error)
}

// Molar is implementation if ObjectReplacer interface.
type Molar struct {
	color color.RGBA
}

// Constructor for Molar struct.
func NewMolar(color color.RGBA) *Molar {
	return &Molar{color: color}
}

// Replace is method which replace from image provided objects and fill objects areas with black rectangle.
func (m Molar) Replace(img domain.CleanableImage, objects []domain.CleanableObject) (domain.CleanableImage, error) {
	drawableImg := image.NewRGBA(img.Bounds())
	draw.Draw(drawableImg, img.Bounds(), img, image.Point{}, draw.Src)
	for _, object := range objects {
		draw.Draw(drawableImg, object.Bounds(), &image.Uniform{C: m.color}, image.Point{}, draw.Src)
	}
	return domain.NewImageFromStd(drawableImg)
}

