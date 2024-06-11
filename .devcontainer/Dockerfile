FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

USER vscode

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2