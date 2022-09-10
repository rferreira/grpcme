all: build-server build-client

proto-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

proto-gen:
	protoc --version
	protoc-gen-go --version
	protoc --go_out=. \
		--go-grpc_out=. \
		--proto_path=. \
		grpcme.proto

clean:
	rm -rf dist

build-server:
	go build -o dist//grpcmed cmd/daemon/grpcmed.go

build-client:
	go build -o dist/grpcme cmd/client/grpcme.go