# Go params
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
LDFLAGS='-s -w -extldflags "-static"'	# Build as static binary & remove symbol info.
BINARY_NAME=twitter

all: build
build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) --ldflags $(LDFLAGS) -i -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/ChimeraCoder/anaconda
