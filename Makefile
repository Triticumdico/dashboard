SHELL = /bin/zsh
GO_BINARY := $(shell which go)
GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
MIN_GO_MAJOR_VERSION = 1
MIN_GO_MINOR_VERSION = 17
MAIN_PACKAGE = ./src/app/backend
RELEASE_VERSION = v1.0.0
SERVE_BINARY = .tmp/serve/dashboard

.PHONY: clean
clean:
	rm -rf .tmp

.PHONY: build-backend
build-backend: ensure-go
	CGO_ENABLED=0 go build -ldflags "-X $(MAIN_PACKAGE)/client.Version=$(RELEASE_VERSION)" -gcflags="all=-N -l" -o $(SERVE_BINARY) $(MAIN_PACKAGE)

.PHONY: run-backend
run-backend:
	$(SERVE_BINARY)

.PHONY: serve-backend
serve-backend: build-backend run-backend


.PHONY: ensure-go
ensure-go:
ifndef GO_BINARY
	$(error "Cannot find go binary")
endif
	@if [ $(GO_MAJOR_VERSION) -gt $(MIN_GO_MAJOR_VERSION) ]; then \
		exit 0 ;\
  elif [ $(GO_MAJOR_VERSION) -lt $(MIN_GO_MAJOR_VERSION) ]; then \
		exit 1; \
  elif [ $(GO_MINOR_VERSION) -lt $(MIN_GO_MINOR_VERSION) ] ; then \
		exit 1; \
  fi