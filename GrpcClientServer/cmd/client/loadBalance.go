package main

import (
	temperatureMeasurements "GrpcClientServer/grpcInterfaces"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"time"
)

const numberOsServers = 1

func main() {
	fmt.Println("Running client side")
	clientConns := make([]*grpc.ClientConn, numberOsServers)
	for i := 1; i <= numberOsServers; i++ {
		ip := "172.20.0." + fmt.Sprintf("%d", i)
		port := 9092 + i
		serverAddress := ip + ":" + fmt.Sprintf("%d", port)
		conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal("Could not connect the servers: %w", err)
		}
		clientConns[i-1] = conn
	}
	defer closeConnections(clientConns)

	counter := 0
	connectionIndex := 0
	for {
		c := temperatureMeasurements.NewTemperatureMeasurementsServiceClient(clientConns[connectionIndex])
		if connectionIndex > 0 && connectionIndex%numberOsServers == 0 {
			getRequest := temperatureMeasurements.GetRequest{Duration: "week"}
			response, err := c.GetMeasurementsData(context.Background(), &getRequest)
			if err != nil {
				go log.Fatal("Error when calling GetMeasurementsData: %w", err)
			}
			log.Printf("Response from server is:\n %s", response.Value)
		} else {
			sensorId, currentMeasurement := twoRandomValues()
			setRequest := temperatureMeasurements.SetRequest{SensorId: int32(sensorId), Temperature: int32(currentMeasurement)}
			c.SetMeasurementData(context.Background(), &setRequest)
		}

		counter++
		if counter > numberOsServers {
			counter = 0
		}
		connectionIndex++
		if connectionIndex%numberOsServers == 0 {
			connectionIndex = 0
		}

		waitForRandomTime()
	}
}

func closeConnections(conns []*grpc.ClientConn) {
	for _, conn := range conns {
		err := conn.Close()
		if err != nil {
			log.Fatal("Failed to close connection")
		}
	}
}

func waitForRandomTime() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10
	seconds := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(seconds) * time.Second)
}

func twoRandomValues() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5), rand.Intn(30)
}
