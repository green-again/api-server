# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.16.9-buster AS build

# Change working dir -> for locating configrations
RUN mkdir -p /app
WORKDIR /app

# COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

FROM build AS dev
# Copy the local package files to the container's workspace.
COPY . .

# Build the outyet command inside the container.
RUN go install ./cmd/server

# Run the server command by default when the container starts.
CMD /go/bin/server
