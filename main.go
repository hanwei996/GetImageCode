package main

import (
	"GetImageCode/handler"
	pb "GetImageCode/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("getimagecode"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterGetImageCodeHandler(srv.Server(), new(handler.GetImageCode))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
