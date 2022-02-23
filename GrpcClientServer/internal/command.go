package internal

type command interface {
	execute(server *SensorsSrv)
}

type getCommand struct {
	duration  string
	replyChan chan string
}

func (cmd getCommand) execute(server *SensorsSrv) {
	cmd.replyChan <- server.getMeasurementsResults(cmd.duration)
}

type setCommand struct {
	lastMeasurement int32
	sensorId        int32
}

func (cmd setCommand) execute(server *SensorsSrv) {
	server.addMeasurement(cmd.sensorId, cmd.lastMeasurement)
}
