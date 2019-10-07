package usecases

import (
	"github.com/dimitriin/image-service/internal/domain"
	"image"
	"image/color"
	"reflect"
	"testing"
)

func getBlackBorderedRectangle(width, height, border int) *image.RGBA {
	black := color.RGBA{0, 0, 0, 1}
	white := color.RGBA{255, 255, 255, 1}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	r := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r.Set(x, y, black)
		}
	}

	for x := border; x < width - border; x++ {
		for y := border; y < height - border; y++ {
			r.Set(x, y, white)
		}
	}
	return r
}

func TestBlackBorders(t *testing.T)  {
	black := color.RGBA{0, 0, 0, 1}
	borderedRectangle := getBlackBorderedRectangle(200, 100, 5)

	borderedImg, err := domain.NewImageFromStd(borderedRectangle)
	if err != nil {
		t.Fatal(err.Error())
	}

	b := NewBarber(black, 0.5)
	cuttedImage, err := b.Cut(borderedImg)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedRectangle := getBlackBorderedRectangle(190, 90, 0)

	expectedImg, err := domain.NewImageFromStd(expectedRectangle)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !reflect.DeepEqual(expectedImg.Bounds(), cuttedImage.Bounds()) {
		t.Fatal("expected image is not equal cutted image")
	}
}
