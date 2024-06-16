package main

import (
	"os"

	_ "embed"

	"github.com/derarken/vlr-api/src/service"
	"google.golang.org/grpc/grpclog"
)

//go:embed gen/openapiv2/vlr/api/api_service.swagger.json
var swaggerJSON []byte

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	grpclog.SetLoggerV2(log)

	service.Start(swaggerJSON)
}
