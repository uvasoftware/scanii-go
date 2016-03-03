# vim:ft=make:

GOCMD = go
GOBUILD = $(GOCMD) build
GOGET = $(GOCMD) get
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test

all: install

test:
	$(GOTEST) -v -cover github.com/uvasoftware/scanii-go

install:
	$(GOINSTALL) -v

clean:
	$(GOCLEAN) -n -i -x

build: 
	$(GOBUILD) -v -race 
