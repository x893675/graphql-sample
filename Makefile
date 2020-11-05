# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif


generate: gqlgen
	#go run github.com/99designs/gqlgen && go run ./app/models/model_tags/model_tags.go
	$(GQL_GEN)
	go run tools/model-tag/main.go

build: fmt checkfmt

	CGO_ENABLED=0 go build -o dist/graphql-sample -a -ldflags "-w -s" ./cmd


run:
	./dist/graphql-sample


fmt:
	gofmt -s -w ./

checkfmt:
	@echo checking gofmt...
	@res=$$(gofmt -d -e -s $$(find . -type d \( -path ./src/vendor -o -path ./tests \) -prune -o -name '*.go' -print)); \
	if [ -n "$${res}" ]; then \
		echo checking gofmt fail... ; \
		echo "$${res}"; \
		exit 1; \
	fi

gqlgen:
ifeq (, $(shell which gqlgen))
	@{ \
	set -e ;\
	GQL_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$GQL_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get github.com/99designs/gqlgen ;\
	rm -rf $$CLIENT_GEN_TMP_DIR ;\
	}
GQL_GEN=$(GOBIN)/gqlgen
else
GQL_GEN=$(shell which gqlgen)
endif