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

const numberOfServers = 5
const serverPort = 9092

func main() {
	fmt.Println("Running client side")
	clientConns := make([]*grpc.ClientConn, numberOfServers)
	for i := 1; i <= numberOfServers; i++ {
		ip := "172.20.0." + fmt.Sprintf("%d", i+1)
		serverAddress := ip + ":" + fmt.Sprintf("%d", serverPort)
		fmt.Println("Connecting server at ip: ", serverAddress)
		conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal("Could not connect the servers: %w", err)
		}
		fmt.Println("connected to server %w", serverAddress)
		clientConns[i-1] = conn
	}
	defer closeConnections(clientConns)

	connectionIndex := 1
	for {
		c := temperatureMeasurements.NewTemperatureMeasurementsServiceClient(clientConns[connectionIndex-1])
		if connectionIndex == 4 {
			fmt.Println("Sending get request")
			getRequest := temperatureMeasurements.GetRequest{Duration: "week"}
			response, err := c.GetMeasurementsData(context.Background(), &getRequest)
			if err != nil {
				go log.Fatal("Error when calling GetMeasurementsData: %w", err)
			}
			log.Printf("Response from server is:\n %s", response.Value)
		} else {
			sensorId, currentMeasurement := twoRandomValues()
			fmt.Println("sending set request, sensorId = ", sensorId, ", temperature = ", currentMeasurement)
			setRequest := temperatureMeasurements.SetRequest{SensorId: int32(sensorId), Temperature: int32(currentMeasurement)}
			c.SetMeasurementData(context.Background(), &setRequest)
		}

		connectionIndex++
		if connectionIndex > numberOfServers {
			connectionIndex = 1
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
	max := 5
	seconds := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(seconds) * time.Second)
}

func twoRandomValues() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5), rand.Intn(30)
}
