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

# watch the shared,clients and service folder
ENTRYPOINT  go run services/user/cmd/migration/main.go && reflex -r '^(?:shared|clients|services/user)(\/[a-z_\-\s0-9\.]+)+\.go$'  -s -- sh -c 'go run services/user/cmd/service/main.go' -v
