package main

import (
	temperatureMeasurements "GrpcBased/grpcInterfaces"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	fmt.Println("Running client side")
	conn, err := grpc.Dial("127.0.0.1:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect the server: %w", err)
	}
	defer closeConnection(conn)

	c := temperatureMeasurements.NewTemperatureMeasurementsServiceClient(conn)

	//setRequest1 := temperatureMeasurements.SetRequest{SensorId: 1, Temperature: 16}
	//c.SetMeasurementData(context.Background(), &setRequest1)
	//
	//setRequest2 := temperatureMeasurements.SetRequest{SensorId: 3, Temperature: 7}
	//c.SetMeasurementData(context.Background(), &setRequest2)
	//
	//setRequest3 := temperatureMeasurements.SetRequest{SensorId: 12, Temperature: 18}
	//c.SetMeasurementData(context.Background(), &setRequest3)

	getRequest := temperatureMeasurements.GetRequest{Duration: "day"}
	response, err := c.GetMeasurementsData(context.Background(), &getRequest)
	if err != nil {
		go log.Fatal("Error when calling GetMeasurementsData: %w", err)
	}
	log.Printf("Response from server is:\n %s", response.Value)
}

func closeConnection(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatal("Failed to close connection")
	}
}
