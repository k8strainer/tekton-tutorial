# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM docker.io/library/golang:1.12 as builder
#FROM docker.io/library/golang:1.20 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/k8strainer/tekton-tutorial
COPY picalc.go picalc_test.go .
#RUN go mod init

# Run unit tests
RUN CGO_ENABLED=0 GOOS=linux go test

# Build the code inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o picalc

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
#FROM docker.io/library/alpine:latest
FROM registry.access.redhat.com/ubi8

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/k8strainer/tekton-tutorial/picalc /picalc

# Service must listen to $PORT environment variable.
# This default value facilitates local development.
ENV PORT 8080

# Run the web service on container startup.
CMD ["/picalc"]
