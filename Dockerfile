FROM golang:1.13.1
COPY . /app
WORKDIR /app
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/bin/image-service .
CMD ["./image-service"]

EXPOSE 50051