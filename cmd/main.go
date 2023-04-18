package main

import (
	"fmt"
	"log"
	"net"

	"github.com/s-shaida/xm-go-company-svc/pkg/config"
	"github.com/s-shaida/xm-go-company-svc/pkg/db"
	pb "github.com/s-shaida/xm-go-company-svc/pkg/pb"
	services "github.com/s-shaida/xm-go-company-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Company Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCompanyServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
