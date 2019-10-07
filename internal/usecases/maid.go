package usecases

import (
	"github.com/dimitriin/image-service/internal/domain"
)

// ImageCleaner cleans image from objects and cuts borders.
type ImageCleaner interface {
	Clean(image domain.CleanableImage) (domain.CleanableImage, error)
}

// Maid is ImageCleaner implementation.
type Maid struct {
	detector ObjectDetector
	replacer ObjectReplacer
	cutter   ImageCutter
}

// NewMaid is constructor for Maid.
func NewMaid(detector ObjectDetector, replacer ObjectReplacer, cutter ImageCutter) *Maid {
	return &Maid{detector: detector, replacer: replacer, cutter: cutter}
}

// Clean removes text, borders from image. Removed text is filled with black rectangle.
func (m Maid) Clean(img domain.CleanableImage) (domain.CleanableImage, error) {
	objects, err := m.detector.Detect(img)
	if err != nil {
		return img, err
	}

	img, err = m.replacer.Replace(img, objects)
	if err != nil {
		return img, err
	}

	return m.cutter.Cut(img)
}


