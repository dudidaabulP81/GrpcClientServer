package internal

import (
	"GrpcClientServer/grpcInterfaces"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	SingleDay = "day"
	Week      = "week"
)

type SensorsSrv struct {
	cmds   chan command
	rdb    *redis.Client
	rwLock sync.RWMutex
}

func (s *SensorsSrv) GetMeasurementsData(ctx context.Context, req *grpcInterfaces.GetRequest) (*grpcInterfaces.MeasurementsDataResult, error) {
	duration := req.Duration
	validDuration := duration == SingleDay || duration == Week
	if !validDuration {
		return &grpcInterfaces.MeasurementsDataResult{Value: fmt.Sprintf("Invalid perion of time: %s", duration)}, nil
	}
	replyChan := make(chan string)
	s.cmds <- getCommand{duration: duration, replyChan: replyChan}
	reply := <-replyChan
	return &grpcInterfaces.MeasurementsDataResult{Value: reply}, nil
}

func (s *SensorsSrv) SetMeasurementData(ctx context.Context, req *grpcInterfaces.SetRequest) (*grpcInterfaces.Empty, error) {
	sensorId := req.SensorId
	currentMeasurement := req.Temperature
	s.cmds <- setCommand{lastMeasurement: currentMeasurement, sensorId: sensorId}
	return &grpcInterfaces.Empty{}, nil
}

func NewSensorsServer() *SensorsSrv {
	return &SensorsSrv{make(chan command), nil, sync.RWMutex{}}
}

func (s *SensorsSrv) Run() {
	s.rdb = redis.NewClient(&redis.Options{
		Addr:     "172.20.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := s.rdb.Ping().Result()
	if err != nil {
		fmt.Println("Failed to connect to database")
		os.Exit(1)
	}
	fmt.Println("Connected to database")
	coresNum := runtime.NumCPU()
	for i := 0; i < coresNum; i++ {
		go s.worker()
	}
}

func (s *SensorsSrv) worker() {
	for cmd := range s.cmds {
		cmd.execute(s)
	}
}

func (s *SensorsSrv) addMeasurement(sensorId int32, lastMeasurement int32) {
	today := DateToString(time.Now())
	key := today + ":" + fmt.Sprint(sensorId)

	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	haveSetForToday := s.rdb.SIsMember(today, "dummyInitiator").Val()
	if !haveSetForToday {
		s.rdb.SAdd(today, "dummyInitiator")
	}

	arbitraryFieldToCheckKeyExistence := "NumOfMeasurements"
	_, err := s.rdb.HGet(key, arbitraryFieldToCheckKeyExistence).Result()
	foundKeyInDb := err == nil
	if !foundKeyInDb {
		s.rdb.SAdd(today, fmt.Sprintf("%v", sensorId))
		data := make(map[string]interface{})
		data["Minimum"] = lastMeasurement
		data["Maximum"] = lastMeasurement
		data["NumOfMeasurements"] = 1
		data["Average"] = float32(lastMeasurement)
		s.rdb.HMSet(key, data)
	} else {
		data := s.rdb.HGetAll(key).Val()
		minimum, _ := strconv.ParseInt(data["Minimum"], 10, 32)
		maximum, _ := strconv.ParseInt(data["Maximum"], 10, 32)
		numOfMeasurements, _ := strconv.ParseInt(data["NumOfMeasurements"], 10, 32)
		average, _ := strconv.ParseFloat(data["Average"], 32)

		average = (average*float64(numOfMeasurements) + float64(lastMeasurement)) / (float64(numOfMeasurements + 1))
		if minimum > int64(lastMeasurement) {
			minimum = int64(lastMeasurement)
		}
		if maximum < int64(lastMeasurement) {
			maximum = int64(lastMeasurement)
		}
		numOfMeasurements++
		s.rdb.HSet(key, "Minimum", fmt.Sprintf("%v", minimum))
		s.rdb.HSet(key, "Maximum", fmt.Sprintf("%v", maximum))
		s.rdb.HSet(key, "NumOfMeasurements", fmt.Sprintf("%v", numOfMeasurements))
		s.rdb.HSet(key, "Average", fmt.Sprintf("%v", average))
	}
}

func (s *SensorsSrv) keysRelevantToDate(date string) *[]string {
	keys := s.rdb.SMembers(date).Val()
	return &keys
}

func (s *SensorsSrv) getMeasurementsOfDay(date string, sb *strings.Builder) {
	relevantKeys := s.keysRelevantToDate(date)
	sb.WriteString(date + ":\n")
	for _, key := range *relevantKeys {
		if key != "dummyInitiator" {
			data := s.rdb.HGetAll(date + ":" + key).Val()
			sb.WriteString(fmt.Sprintf("\tSensor id: %s\t\tMin: %s\t\tMax: %s\t\tAverage: %s\n", key, data["Minimum"], data["Maximum"], data["Average"]))
		}
	}
}

func (s *SensorsSrv) getMeasurementsResults(duration string) string {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	sb := strings.Builder{}
	if duration == SingleDay {
		today := DateToString(time.Now())
		s.getMeasurementsOfDay(today, &sb)
		result := sb.String()
		sb.Reset()
		return result
	}
	if duration == Week {
		const daysInWeek int = 7
		for i := 0; i < daysInWeek; i++ {
			date := DateToString(time.Now().AddDate(0, 0, -i))
			s.getMeasurementsOfDay(date, &sb)
		}
		result := sb.String()
		sb.Reset()
		return result
	}
	return fmt.Sprintf("The requested duration, %s, is not valid", duration)
}
