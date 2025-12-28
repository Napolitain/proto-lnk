// Package proto contains the generated protobuf and gRPC code.
// Run `go generate ./...` from the repository root to regenerate.
package proto

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative config.proto
