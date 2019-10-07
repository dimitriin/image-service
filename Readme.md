# Image Service
It's GRPC service with functionality:
* Remove texts on the image with black rectangle;
* Remove black borders around image.  

Protocol description provided [here](./proto/v1/service.proto).

# Build
``
GOOS=linux BINARY_OUTPUT="bin/image-service" make build
``

# Run
``
SERVER_PORT=50051 GOOGLE_APPLICATION_CREDENTIALS=path_to_creds.json make run
``

`GOOGLE_APPLICATION_CREDENTIALS` is path to Google Vision API credentials. 
How to retrieve credentials read [here](https://cloud.google.com/vision/docs/before-you-begin).  

# Client example
Start server and run client example:
``
IMG_INPUT=path_to_some_image.jpg IMG_OUTPUT=path_to_cleared_file.jpg go run cmd/client/main.go
``