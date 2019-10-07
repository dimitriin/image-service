package usecases

import (
	"bytes"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/dimitriin/image-service/internal/domain"
	"github.com/dimitriin/image-service/internal/infrastructure"
	"github.com/uber-go/zap"
)

// ObjectDetector expected interface for cleanable objects detection on image.
type ObjectDetector interface {
	Detect(img domain.CleanableImage) ([]domain.CleanableObject, error)
}

// GoogleVisionTextDetector is implementation of ObjectDetector using Google Vision API.
type GoogleVisionTextDetector struct {
	client *vision.ImageAnnotatorClient
	ctx context.Context
	logger *zap.SugaredLogger
}

// Constructor for GoogleVisionTextDetector.
func NewGoogleVisionTextDetector(client *vision.ImageAnnotatorClient,
	ctx context.Context, l *zap.SugaredLogger) *GoogleVisionTextDetector {
	return &GoogleVisionTextDetector{client: client, ctx: ctx, logger: l}
}

// Detect calls DetectText method of Google Vision API and get cleanable objects from response.
func (g *GoogleVisionTextDetector) Detect(img domain.CleanableImage) ([]domain.CleanableObject, error) {
	objects := make([]domain.CleanableObject, 0)

	content, err := img.Content()
	if err != nil {
		return objects, err
	}
	id := infrastructure.GetRequestId(content)

	vImage, err := vision.NewImageFromReader(bytes.NewBuffer(content))
	if err != nil {
		return objects, err
	}

	annotations, err := g.client.DetectTexts(g.ctx, vImage, nil, 1000)
	if err != nil {
		return objects, err
	}

	for _, annotation := range annotations {
		if annotation.BoundingPoly != nil {
			obj, err := domain.NewObject(*annotation.BoundingPoly)
			if err != nil {
				return objects, err
			}
			objects = append(objects, obj)
			g.logger.Infof("[%s] detected object: %v", id, obj)
		}
	}

	return objects, nil
}



