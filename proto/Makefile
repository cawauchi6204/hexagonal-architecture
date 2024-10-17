.PHONY: gen-protoc
gen-protoc:
	docker compose run --rm --build protoc

go-proto:
	if [ ! -d "./go/protos" ]; then mkdir -p ./go/protos; fi
	protoc -I. \
	--proto_path=${HOME}/ghq/github.com/googleapis/googleapis \
	--experimental_allow_proto3_optional \
	--go_out=./go/protos --go_opt paths=source_relative \
	-I. \
	*.proto

clean:
	rm -fr ./go/protos

all:
	make clean
	make go-proto
