package main

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/dimitriin/image-service/internal/interfaces"
	"github.com/dimitriin/image-service/internal/usecases"
	pb "github.com/dimitriin/image-service/pkg/imagepb/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"image/color"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultPort = ":50051"
)

func main() {

	// Init logger
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()
	sugar := logger.Sugar()

	// Get port from ENV.
	port := os.Getenv("SERVER_PORT")
	if len(port) > 0 {
		port = defaultPort
	}
	sugar.Infof("starting server on %s port", port)

	// Creates a Google Vision API client.
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		sugar.Fatalf("failed to create google vision api client: %v", err)
	}
	defer func() {
		_ = client.Close()
	}()

	black := color.RGBA{R: 0, G: 0, B: 0, A: 1}
	// Creates object detector using Google Vision API.
	objDetector := usecases.NewGoogleVisionTextDetector(client, ctx, sugar)
	// Creates object replacer with black rectangle.
	objReplacer := usecases.NewMolar(black)
	// Creates black borders cutter around image.
	borderCutter := usecases.NewBarber(black, 0.5)
	cleaner := usecases.NewMaid(objDetector, objReplacer, borderCutter)

	// Initializes GRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		sugar.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterImageServiceServer(s, interfaces.NewServer(cleaner, sugar))

	// Setup signals processing.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		s.GracefulStop()
		sugar.Info("server was gracefully stopped")
	}()

	// Starts GRPC server.
	if err := s.Serve(lis); err != nil {
		sugar.Fatalf("failed to serve: %v", err)
	}
}

