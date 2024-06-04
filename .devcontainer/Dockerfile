FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

USER vscode

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4.0