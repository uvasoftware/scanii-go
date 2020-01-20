# vim:ft=make:

GOCMD = go
GOBUILD = $(GOCMD) build
GOGET = $(GOCMD) get
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test

all: install

test:
	$(GOTEST) -v ./pkg/client/

install:
	$(GOINSTALL) -v ./pkg/client/

clean:
	$(GOCLEAN) -n -i -x

build: 
	$(GOBUILD) -v -race ./pkg/client/
