package main

import (
	"os"

	"github.com/derarken/vlr-api/src/service"
	"google.golang.org/grpc/grpclog"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	grpclog.SetLoggerV2(log)

	service.Start()
}
