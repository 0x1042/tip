export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

.PHONY: fmt build vet clean

clean:
	rm -fr bin

vet:
	go vet ./...

fmt:
	go install mvdan.cc/gofumpt@latest
	go mod tidy
	gofumpt -l -w .

build: fmt clean vet
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/tip
