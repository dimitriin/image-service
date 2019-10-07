package usecases

import (
	"bytes"
	"github.com/dimitriin/image-service/internal/domain"
	"image"
	"image/color"
	"testing"
)

func getRectangle(c color.Color) *image.RGBA {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	r := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r.Set(x, y, c)
		}
	}
	return r
}

func TestWhiteRectangle(t *testing.T) {
	black := color.RGBA{0, 0, 0, 1}
	white := color.RGBA{255, 255, 255, 1}

	whiteRectangle := getRectangle(white)
	img, err := domain.NewImageFromStd(whiteRectangle)
	if err != nil {
		t.Fatal(err.Error())
	}

	m := NewMolar(black)
	replacedImg, err := m.Replace(img, []domain.CleanableObject{img.Bounds()})
	if err != nil {
		t.Fatal(err.Error())
	}

	replacedContent, err := replacedImg.Content()
	if err != nil {
		t.Fatal(err.Error())
	}

	blackRectangle := getRectangle(black)
	blackImg, err := domain.NewImageFromStd(blackRectangle)
	if err != nil {
		t.Fatal(err.Error())
	}

	content, err := blackImg.Content()
	if err != nil {
		t.Fatal(err.Error())
	}

	res := bytes.Compare(content, replacedContent)
	if res != 0 {
		t.Fatal("content is not equal")
	}

}
