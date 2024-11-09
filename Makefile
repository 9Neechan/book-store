include .env
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-book-store-api

generate-book-store-api:
	mkdir -p pkg/book_store/v1/book 
	protoc --proto_path api/book_store/v1 \
	--go_out=pkg/book_store/v1/book --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/book_store/v1/book --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/book_store/v1/book.proto 

	mkdir -p pkg/book_store/v1/author
	protoc --proto_path api/book_store/v1 \
	--go_out=pkg/book_store/v1/author --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/book_store/v1/author --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/book_store/v1/author.proto 

	mkdir -p pkg/book_store/v1/buy
	protoc --proto_path api/book_store/v1 \
	--go_out=pkg/book_store/v1/buy --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/book_store/v1/buy --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/book_store/v1/buy.proto 

	mkdir -p pkg/book_store/v1/genre
	protoc --proto_path api/book_store/v1 \
	--go_out=pkg/book_store/v1/genre --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/book_store/v1/genre --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/book_store/v1/genre.proto 