package main

import (
	temperatureMeasurements "GrpcBased/grpcInterfaces"
	"GrpcBased/internal"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Running server side")
	server := internal.NewSensorsServer()
	server.Run()
	grpcServer := grpc.NewServer()

	temperatureMeasurements.RegisterTemperatureMeasurementsServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	grpcServer.Serve(lis)

}
