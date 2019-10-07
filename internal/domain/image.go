package domain

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/jpeg"
)

// CleanableImage represents interface for image entity to clean.
type CleanableImage interface {
	image.Image
	Content() ([]byte, error)
}

// Image is entity which implements CleanableImage interface.
type Image struct {
	img image.Image
}

func (c *Image) ColorModel() color.Model {
	return c.img.ColorModel()
}

func (c *Image) Bounds() image.Rectangle {
	return c.img.Bounds()
}

func (c *Image) At(x, y int) color.Color {
	return c.img.At(x, y)
}

func (c *Image) Content() ([]byte, error) {
	buf := bytes.Buffer{}
	err := jpeg.Encode(&buf, c.img, nil)
	return buf.Bytes(), err
}

// NewImage constructor for image from []bytes array.
func NewImage(content []byte) (*Image, error) {
	img, format, err := image.Decode(bytes.NewReader(content))

	if err != nil {
		return nil, err
	}

	if format != "jpeg" {
		return nil, errors.New("only jpeg format is supported")
	}

	return &Image{img}, nil
}

// NewImageFromStd constructor for Image from standard image.Image object.
func NewImageFromStd(img image.Image) (*Image, error) {
	buf := &bytes.Buffer{}
	err := jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, err
	}

	return NewImage(buf.Bytes())
}