NAME=padavan_exporter
BINDIR=bin
VERSION=$(shell git describe --tags || echo "unknownversion")
GOBUILD=CGO_ENABLED=0 go build
CMDPATH=./cmd/padavan_exporter

all: linux-amd64 darwin-amd64 freebsd-amd64 windows-amd64 # Most used

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@-$(VERSION) $(CMDPATH)

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@-$(VERSION) $(CMDPATH)

freebsd-amd64:
	GOARCH=amd64 GOOS=freebsd $(GOBUILD) -o $(BINDIR)/$(NAME)-$@-$(VERSION) $(CMDPATH)

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@-$(VERSION).exe $(CMDPATH)

clean:
	rm $(BINDIR)/*
