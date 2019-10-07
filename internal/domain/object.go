package domain

import (
	"errors"
	"google.golang.org/genproto/googleapis/cloud/vision/v1"
	"image"
)

// CleanableObject represents interface of the object to be cleaned on CleanableImage image.
type CleanableObject interface {
	Bounds() image.Rectangle
}

// Object is entity which implements CleanableObject interface.
type Object struct {
	bound vision.BoundingPoly
}

// Constructor for Object.
func NewObject(bound vision.BoundingPoly) (Object, error) {
	// Expected only rectangles with 4 vertices.
	if len(bound.Vertices) != 4 {
		return Object{}, errors.New("rectangle expected")
	}
	return Object{bound: bound}, nil
}
// Bounds get object bounds as rectangle points.
func (o Object) Bounds() image.Rectangle {
	minX := o.bound.Vertices[0].X
	minY := o.bound.Vertices[0].Y

	maxX := o.bound.Vertices[0].X
	maxY := o.bound.Vertices[0].Y

	for i := 1 ; i < len(o.bound.Vertices) ; i++ {
		if minX > o.bound.Vertices[i].X {
			minX = o.bound.Vertices[i].X
		}

		if maxX < o.bound.Vertices[i].X {
			maxX = o.bound.Vertices[i].X
		}

		if minY > o.bound.Vertices[i].Y {
			minY = o.bound.Vertices[i].Y
		}

		if maxY < o.bound.Vertices[i].Y {
			maxY = o.bound.Vertices[i].Y
		}
	}

	return image.Rect(int(minX), int(minY), int(maxX), int(maxY))
}

