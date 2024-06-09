package main

import (
	"os"

	"github.com/derarken/vlr-api/src/api"
	"google.golang.org/grpc/grpclog"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	grpclog.SetLoggerV2(log)

	// service.Start()

	match, err := api.GetMatch("348476/gen-g-vs-team-heretics-champions-tour-2024-masters-shanghai-gf")
	if err != nil {
		panic(err)
	}
	grpclog.Infof("Match: %v", match)
}
