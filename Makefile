GOCMD:=$(shell which go)
GOLINT:=$(shell which golint)
GOIMPORT:=$(shell which goimports)
GOFMT:=$(shell which gofmt)
GOBUILD:=$(GOCMD) build
GOINSTALL:=$(GOCMD) install
GOCLEAN:=$(GOCMD) clean
GOTEST:=$(GOCMD) test
GOMODTIDY:=$(GOCMD) mod tidy
GOGET:=$(GOCMD) get
GOLIST:=$(GOCMD) list
GOVET:=$(GOCMD) vet
GOPATH:=$(shell $(GOCMD) env GOPATH)
u := $(if $(update),-u)

BINARY_NAME:= godoc
PACKAGES:=$(shell $(GOLIST) github.com/mitsuaaki/godoc github.com/mitsuaaki/godoc/cmd/godoc github.com/mitsuaaki/godoc/app)
GOFILES:=$(shell find . -name "*.go" -type f)

all: test build

.PHONY: build
build: deps
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/godoc

.PHONY: install
install: deps
	$(GOINSTALL) ./cmd/godoc

.PHONY: test
test:
	echo "mode: count" > .github/coverage.out
	for PKG in $(PACKAGES); do \
		$(GOCMD) test -v -covermode=count -coverprofile=profile.out $$PKG > tmp.out; \
		cat tmp.out; \
		if grep -q "^--- FAIL" tmp.out; then \
			rm tmp.out; \
			exit 1; \
		elif grep -q "build failed" tmp.out; then \
			rm tmp.out; \
			exit; \
		fi; \
		if [ -f profile.out ]; then \
			cat profile.out | grep -v "mode:" >> .github/coverage.out; \
			rm profile.out; \
		fi; \
	done

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: deps
deps:
	$(GOMODTIDY)

.PHONY: vet
vet: deps
	$(GOVET) $(PACKAGES)

.PHONY: fmt
fmt:
	$(GOFMT) -s -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: view-covered
view-covered:
	$(GOTEST) -coverprofile=.github/cover.out $(TARGET)
	$(GOCMD) tool cover -html=.github/cover.out