FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy all the source code, including the subdirectories, but only the .go files
COPY . ./

RUN ["ls", "-la"]

# Build
# CGO_ENABLED=0 is required to build a static binary so it doesnt depend on any shared libraries
RUN CGO_ENABLED=0 GOOS=linux go build -o /hestia-api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 9000

# Run
CMD ["/hestia-api"]