SHELL           = /bin/bash
APP             = k8sinfra-gen
PACKAGE         = github.com/Azure/k8s-infra/hack/generator

include ../../tools.mk

DATE            ?= $(shell date +%FT%T%z)
VERSION         ?= $(shell git rev-list -1 HEAD)
SHORT_VERSION   ?= $(shell git rev-parse --short HEAD)
TREE_STATE      ?= $(if $(shell git status --porcelain),dirty,clean)
GO_DIRS         := $(shell $(GO) list -f '{{.Dir}}' ./...)

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all: header-check fmt test gen-arm

# There is a ci specific target because we want the CI pass to fail if
# the code has not been go fmt-ed, whereas locally we want "make all"
# to just format the code for you
.PHONY: ci
ci: build test-cover gen-arm

.PHONY: gen-arm
gen-arm: build ; $(info $(M) running k8sinfra-gen gen…) @ ## Running gen
	$(Q) ./bin/$(APP) gen-types azure-arm.yaml
	# Run with verbosity 4 to see a stack trace for errors.
	# $(Q) ./bin/$(APP) gen azure-arm.yaml -v 4

.PHONY: build
build: tidy lint ; $(info $(M) building ./bin/$(APP))
	$(Q) $(GO) build -ldflags "-X $(PACKAGE)/pkg/codegen.GitCommit=$(SHORT_VERSION) -X $(PACKAGE)/pkg/codegen.GitTreeState=$(TREE_STATE)" -o ./bin/$(APP)

.PHONY: lint
lint: $(GOLANGCI_LINT) ; $(info $(M) running golangci configured linters…) ## Lint codebase
	$(Q) $(GOLANGCI_LINT) run -v

.PHONY: lint-full
lint-full: $(GOLANGCI_LINT) ; $(info $(M) running golangci configured linters…) ## Run slower linters to detect possible issues
	$(Q) $(GOLANGCI_LINT) run -v --fast=false

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$(Q) $(GOFMT) -l -s -w $(GO_DIRS)

.PHONY: tidy
tidy: ; $(info $(M) running tidy…) @ ## Run tidy
	$Q $(GO) mod tidy

.PHONY: build-debug
build-debug: ; $(info $(M) buiding debug...)
	$Q $(GO)  build -o ./bin/$(APP) -tags debug

.PHONY: test
test: ; $(info $(M) running go test…)
	$(Q) $(GO) test ./... -tags=noexit

.PHONY: test-cover
test-cover: $(GCOV2LCOV) ; $(info $(M) running go test…)
	# NOTE: if you update the 'test-cover' target, also update ./github/workflows/test.yml
	#       for the Windows part of the "test-generator" job.
	$(Q) $(GO) test -tags=noexit -race -covermode atomic -coverprofile=cover.out -coverpkg=./... ./...
	$(Q) $(GCOV2LCOV) -infile cover.out -outfile coverage.lcov

.PHONY: gox
gox: install-tools
	$(Q) $(TOOLS_BIN)/gox -osarch="darwin/amd64 windows/amd64 linux/amd64" -ldflags "-X $(PACKAGE)/cmd.GitCommit=$(VERSION)" -output "./bin/$(SHORT_VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}"
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_darwin_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_darwin_amd64
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_linux_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_linux_amd64
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_windows_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_windows_amd64.exe

.PHONY: diagrams
diagrams: $(patsubst %.dot,%.png,$(wildcard spec/images/*.dot spec/images/*/*.dot))

%.png: %.dot
	$(info $(M) Generating diagram $@ from $< …)
	$(Q) dot $< -Tpng -o $@
