FROM golang:1.10
# Create the folder structure
WORKDIR /go/src/github.com/go-services/gos-project

# Copy project files
COPY . .

# Install dep
RUN go get -u github.com/golang/dep/cmd/dep

# Install dependencies
RUN dep ensure

# Install reflux to watch the app
RUN go get github.com/cespare/reflex

RUN chmod +x /go/src/github.com/go-services/gos-project/services/user/docker/entrypoint.sh
# Run entrypoint
ENTRYPOINT  /go/src/github.com/go-services/gos-project/services/user/docker/entrypoint.sh