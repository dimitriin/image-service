package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/dimitriin/image-service/pkg/imagepb/v1"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Path to image
	input := os.Getenv("IMG_INPUT")

	// Path to cleared image2
	output := os.Getenv("IMG_OUTPUT")

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	c := pb.NewImageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	content, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalf("unable to open file with image: %v", err)
	}

	// Execute image clear request
	r, err := c.Clear(ctx, &pb.ClearRequest{
		Image: &pb.Image{Content:content},
	})
	if err != nil {
		log.Fatalf("could not clear image: %v", err)
	}

	// Save cleared image
	err = ioutil.WriteFile(output, r.Image.Content, 0644)
	if err != nil {
		log.Fatalf("could not clear image: %v", err)
	}
}