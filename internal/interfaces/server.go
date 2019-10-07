package interfaces

import (
	"context"
	"errors"
	"github.com/dimitriin/image-service/internal/domain"
	"github.com/dimitriin/image-service/internal/infrastructure"
	"github.com/dimitriin/image-service/internal/usecases"
	pb "github.com/dimitriin/image-service/pkg/imagepb/v1"
	"go.uber.org/zap"
)

// Server provides handlers for GRPC service.
type Server struct {
	cleaner usecases.ImageCleaner
	logger *zap.SugaredLogger
}

// Constructor for Server.
func NewServer(cleaner usecases.ImageCleaner, logger *zap.SugaredLogger) *Server {
	return &Server{cleaner: cleaner, logger: logger}
}

// Clear handler.
func (s *Server) Clear(ctx context.Context, in *pb.ClearRequest) (*pb.ClearResponse, error) {

	if in.Image == nil || in.Image.Content == nil {
		err := errors.New("no image provided")
		s.logger.Errorf("[%s] %s", "nil", err.Error())
		return &pb.ClearResponse{}, err
	}

	id := infrastructure.GetRequestId(in.Image.Content)
	s.logger.Info("[%s] incoming request", id)

	img, err := domain.NewImage(in.Image.Content)
	if err != nil {
		s.logger.Errorf("[%s] %s", id, err.Error())
		return &pb.ClearResponse{}, err
	}

	cleanedImg, err := s.cleaner.Clean(img)
	if err != nil {
		s.logger.Errorf("[%s] %s", id, err.Error())
		return &pb.ClearResponse{}, err
	}

	content, err := cleanedImg.Content()
	if err != nil {
		s.logger.Errorf("[%s] %s", id, err.Error())
		return &pb.ClearResponse{}, err
	}

	resp := &pb.ClearResponse{
		Image: &pb.Image{
			Content: content,
		},
	}
	s.logger.Errorf("[%s] success response", id)

	return resp, nil
}

// Crop handler.
func (s *Server) Crop(ctx context.Context, in *pb.CropRequest) (*pb.CropResponse, error) {
	return &pb.CropResponse{}, errors.New("not implemented")
}