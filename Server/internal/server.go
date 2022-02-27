package internal

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	SingleDay = "day"
	Week      = "week"
)

const maxLoggedDays int = 7

//const bufferSize int = 512

type DailySensorsData map[int]*sensorData
type PeriodicSensorsData map[string]*DailySensorsData

type sensorData struct {
	minimum           int
	maximum           int
	numOfMeasurements int
	average           float32
}

type accumulatedMeasurements struct {
	minimum                int
	maximum                int
	totalOfAllMeasurements int
	numOfMeasurements      int
}

type Server struct {
	cmds                chan command
	sensorsMeasurements PeriodicSensorsData
	//sb                  strings.Builder
}

func (s *Server) Run() {
	//s.sb.Grow(bufferSize)
	go func() {
		for cmd := range s.cmds {
			cmd.execute(s)
		}
	}()
}

func NewServer() *Server {
	sensorsMeasurements := make(PeriodicSensorsData)
	return &Server{make(chan command), sensorsMeasurements}
	//return &Server{make(chan command), sensorsMeasurements, strings.Builder{}}
}

func (s *Server) Get(w http.ResponseWriter, req *http.Request) {
	//for test only, remove later
	//time.Sleep(3 * time.Second)

	period := req.URL.Query().Get("period")
	var validPeriod bool = (period == SingleDay || period == Week)
	if !validPeriod {
		fmt.Fprintf(w, "period of %s is invalid", period)
		return
	}
	replyChan := make(chan PeriodicSensorsData)
	s.cmds <- getCommand{periodOfTime: period, replyChan: replyChan}
	reply := <-replyChan
	if len(reply) > 0 {
		fmt.Fprintf(w, s.serializeResponse(reply, period))
	} else {
		fmt.Fprintf(w, "No data received")
	}
}

func (s *Server) Set(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	id := req.Form.Get("id")
	clientId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}
	currentMeasurement := req.Form.Get("temperature")
	clientMeasurement, err := strconv.Atoi(currentMeasurement)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}
	fmt.Println("Received data: clientId = ", clientId, "clientMeasurement = ", clientMeasurement)
	s.cmds <- setCommand{lastMeasurement: clientMeasurement, sensorId: clientId}
}

func (s *Server) getMeasurementsResults(period string) PeriodicSensorsData {
	if period == Week {
		return s.sensorsMeasurements
	}
	if period == SingleDay {
		if len(s.sensorsMeasurements) == 0 {
			return s.sensorsMeasurements
		}
		today := DateToString(time.Now())
		measurementsOfToday := s.getMeasurementsOfDay(today)
		result := make(PeriodicSensorsData)
		result[today] = measurementsOfToday
		return result
	}
	fmt.Printf("Invalid requested period parameter: %s", period)
	return make(PeriodicSensorsData)
}

func (s *Server) getMeasurementsOfDay(day string) *DailySensorsData {
	return s.sensorsMeasurements[day]
}

func (s *Server) addMeasurement(sensorId, measurement int) {
	today := DateToString(time.Now())
	if measurementsOfToday, ok := s.sensorsMeasurements[today]; ok {
		if measurementsOfRelevantSensor, ok := (*measurementsOfToday)[sensorId]; ok {
			currentAverage := measurementsOfRelevantSensor.average
			currentNumOfMeasurements := measurementsOfRelevantSensor.numOfMeasurements
			(*s.sensorsMeasurements[today])[sensorId].average = (currentAverage*float32(currentNumOfMeasurements) + float32(measurement)) / (float32(currentNumOfMeasurements + 1))
			((*s.sensorsMeasurements[today])[sensorId].numOfMeasurements)++
			if (*s.sensorsMeasurements[today])[sensorId].minimum > measurement {
				(*s.sensorsMeasurements[today])[sensorId].minimum = measurement
			} else if (*s.sensorsMeasurements[today])[sensorId].maximum < measurement {
				(*s.sensorsMeasurements[today])[sensorId].maximum = measurement
			}
		} else {
			(*s.sensorsMeasurements[today])[sensorId] = &(sensorData{measurement, measurement, 1, float32(measurement)})
		}
	} else {
		if len(s.sensorsMeasurements) >= maxLoggedDays {
			s.removeOldestDailyMeasurements()
		}
		currentDayMeasurements := make(DailySensorsData)
		currentDayMeasurements[sensorId] = &(sensorData{measurement, measurement, 1, float32(measurement)})
		s.sensorsMeasurements[today] = &currentDayMeasurements
	}
}

func (s *Server) removeOldestDailyMeasurements() {
	dateOfDataToRemove := time.Now().AddDate(0, 0, -maxLoggedDays)
	keyToRemove := DateToString(dateOfDataToRemove)
	_, ok := s.sensorsMeasurements[keyToRemove]
	if ok {
		delete(s.sensorsMeasurements, keyToRemove)
	}
}

func (s *Server) calculateAccumulatedMeasurements() map[int]*accumulatedMeasurements {
	sensorIdToAccumulatedMeasurements := make(map[int]*accumulatedMeasurements)
	numOfDocumentedDays := len(s.sensorsMeasurements)
	for i := 0; i < numOfDocumentedDays; i++ {
		dateToScan := time.Now().AddDate(0, 0, -i)
		currentDate := DateToString(dateToScan)
		if _, ok := s.sensorsMeasurements[currentDate]; ok {
			for sensorId, sensorData := range *(s.sensorsMeasurements[currentDate]) {
				sumOfCurrentDayMeasurement := int(float32(sensorData.numOfMeasurements) * sensorData.average)
				_, foundDataForSensorId := sensorIdToAccumulatedMeasurements[sensorId]
				if !foundDataForSensorId {
					sensorIdToAccumulatedMeasurements[sensorId] = &accumulatedMeasurements{sensorData.minimum, sensorData.maximum,
						sumOfCurrentDayMeasurement, sensorData.numOfMeasurements}
				} else {
					if sensorIdToAccumulatedMeasurements[sensorId].minimum > sensorData.minimum {
						sensorIdToAccumulatedMeasurements[sensorId].minimum = sensorData.minimum
					}
					if sensorIdToAccumulatedMeasurements[sensorId].maximum < sensorData.maximum {
						sensorIdToAccumulatedMeasurements[sensorId].maximum = sensorData.maximum
					}
					sensorIdToAccumulatedMeasurements[sensorId].totalOfAllMeasurements += sumOfCurrentDayMeasurement
					sensorIdToAccumulatedMeasurements[sensorId].numOfMeasurements += sensorData.numOfMeasurements
				}
			}
		}
	}
	return sensorIdToAccumulatedMeasurements
}

func (s *Server) serializeAccumulatedMeasurements() string {
	//func (s *Server) serializeAccumulatedMeasurements() {
	var currentSensorData string
	var numberOfAverages float32 = 0.0
	var sumOfAllAverages float32 = 0.0
	//var buffer bytes.Buffer
	//var buffer strings.Builder
	//buffer.Grow(512)
	//buffer.WriteString("\nAccumulated data:\n")
	//buffer.WriteString("\tid\t\tmin\t\tmax\t\taverage\n")
	//s.sb.WriteString("\nAccumulated data:\n")
	//s.sb.WriteString("\tid\t\tmin\t\tmax\t\taverage\n")
	result := "\nAccumulated data:\n"
	result += "\tid\t\tmin\t\tmax\t\taverage\n"
	accumulatedData := s.calculateAccumulatedMeasurements()
	for sensorId, totalMeasurementData := range accumulatedData {
		averageOfCurrentSensor := float32(totalMeasurementData.totalOfAllMeasurements) / float32(totalMeasurementData.numOfMeasurements)
		numberOfAverages++
		sumOfAllAverages += averageOfCurrentSensor
		//fmt.Fprintf(&buffer, "\t%d\t\t%d\t\t%d\t\t%f\n", sensorId, totalMeasurementData.minimum, totalMeasurementData.maximum, averageOfCurrentSensor)
		//fmt.Fprintf(&(s.sb), "\t%d\t\t%d\t\t%d\t\t%f\n", sensorId, totalMeasurementData.minimum, totalMeasurementData.maximum, averageOfCurrentSensor)
		currentSensorData = fmt.Sprintf("\t%d\t\t%d\t\t%d\t\t%f\n", sensorId, totalMeasurementData.minimum, totalMeasurementData.maximum, averageOfCurrentSensor)
		result += currentSensorData
	}
	averageMeasurement := sumOfAllAverages / numberOfAverages
	//fmt.Fprintf(&buffer, "average measurement = %f", averageMeasurement)
	//fmt.Fprintf(&(s.sb), "average measurement = %f", averageMeasurement)
	averageOfAverages := fmt.Sprintf("average measurement = %f", averageMeasurement)
	result += averageOfAverages

	//return buffer.String()
	// nothing to return - everything is in s.sb !!
	return result
}

func (s *Server) serializeResponse(data PeriodicSensorsData, duration string) string {
	//var buffer bytes.Buffer
	//var buffer strings.Builder
	var result string = ""
	var currentSensorData string
	//s.sb.Reset()
	numOfDocumentedDays := len(data)
	for i := 0; i < numOfDocumentedDays; i++ {
		dateToScan := time.Now().AddDate(0, 0, -i)
		currentDate := DateToString(dateToScan)
		if _, ok := data[currentDate]; ok {
			//buffer.WriteString(currentDate)
			//buffer.WriteString(":\n")
			//s.sb.WriteString(currentDate)
			//s.sb.WriteString(":\n")
			result = currentDate + ":\n"
			printedTitles := false
			for sensorId, sensorData := range *(data)[currentDate] {
				if !printedTitles {
					//buffer.WriteString("\tid\t\tmin\t\tmax\t\taverage\n")
					//s.sb.WriteString("\tid\t\tmin\t\tmax\t\taverage\n")
					result += "\tid\t\tmin\t\tmax\t\taverage\n"
					printedTitles = true
				}
				//fmt.Fprintf(&buffer, "\t%d\t\t%d\t\t%d\t\t%f\n", sensorId, sensorData.minimum, sensorData.maximum, sensorData.average)
				//fmt.Fprintf(&(s.sb), "\t%d\t\t%d\t\t%d\t\t%f\n", sensorId, sensorData.minimum, sensorData.maximum, sensorData.average)
				currentSensorData = fmt.Sprintf("\t%d\t\t%d\t\t%d\t\t%f\n", sensorId, sensorData.minimum, sensorData.maximum, sensorData.average)
				result += currentSensorData
			}
		}
	}

	if duration == Week {
		//buffer.WriteString(s.serializeAccumulatedMeasurements())
		//s.serializeAccumulatedMeasurements()
		result += s.serializeAccumulatedMeasurements()
	}

	//return buffer.String()
	//return s.sb.String()
	return result
}
