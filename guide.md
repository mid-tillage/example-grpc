go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/protobuf

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative catalog/catalog.proto