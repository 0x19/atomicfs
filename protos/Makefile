help:
	@echo "usage: make target"
	@echo "       deps: install required build tools"
	@echo "       list: list protoc libraries"
	@echo "       all-go: build all for golang"
	@echo "       all: build everything"
	@echo "       clean-all: clean everything"
	@echo "       lib-[proto set]: clean and build one lib"
	@echo "       clean-[proto set]: clean one lib"

CORTEX_ROOT=$(shell p=$$(pwd); echo $${p%"/protos"})

# Add new protobuf definitions here to add them to *-all targets
protos=\
blocks

proto_std = -Ithird_party/protobuf -Ithird_party/googleapis -I/usr/local/bin/ -I${HOME}/go/bin

# proto_i provides the additional includes to the protoc call
proto_i =

blocks/v1.go: protoset = blocks
blocks.go: blocks/v1.go
	@true
blocks.clean: blocks/v1.clean
	@true
clean-blocks: blocks/v1.clean
	@true


# ======================================================================
# Build time dependencies

deps.go: /usr/local/bin/protoc
	cd /tmp && GO111MODULE=on go get github.com/protocolbuffers/protobuf@latest
	cd /tmp && GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go@latest
	cd /tmp && GO111MODULE=on go get github.com/grpc-ecosystem/grpc-gateway@latest
	cd /tmp && GO111MODULE=on go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: deps deps.go

# ======================================================================
# Generic and pattern targets
# You shouldn't change anything below this line

protos.go = $(foreach l,$(protos),$l.go)
protos-lib = $(foreach l,$(protos),lib-$l)
protos-clean = $(foreach l,$(protos),clean-$l)

list:
	@echo $(protos) | tr ' ' '\n' | sort

list-go:
	@echo $(protos.go)

all.go: $(protos.go)
all: $(protos.go)

# pattern rules evidently require a non-empty recipe to match
lib-%: clean-% %.go
	@echo "done"

clean-%:
	rm -rf $*/dist
	rm -f $*/protos/*.protoset

# required to match tunnels with its version prefix
%.clean:
	rm -rf $*/dist
	rm -f $*/protos/*.protoset

clean: $(protos-clean)

.PHONEY: help list all-go all
.PHONEY: $(protos.go) $(protos-lib) $(protos-clean)

go.go:
	@true

%.go: protoset=$(basename $*)
%.go:
	@echo "making $*.go"
	mkdir -p $*/dist/golang
	protoc $(proto_std) \
	       $(proto_i) \
		--proto_path=$*/protos \
		--go_out=$*/dist/golang \
		--go_opt=paths=source_relative \
		--go-grpc_out=$*/dist/golang \
		--go-grpc_opt=paths=source_relative \
		--descriptor_set_out=$*/protos/$(protoset).protoset \
		--include_source_info \
		--include_imports \
		$*/protos/*.proto

%.docs: protoset=$(basename $*)
%.docs:
	@echo "making $*.docs"
	mkdir -p $*/dist/docs
	protoc $(proto_std) \
	       $(proto_i) \
		   -I./common/protos  \
		--proto_path=$*/protos \
		--doc_out=$*/dist/docs --doc_opt=json,docs.json \
		--descriptor_set_out=$*/protos/$(protoset).protoset \
		--include_source_info \
		--include_imports \
		$*/protos/*.proto

