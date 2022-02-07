package internal

type command interface {
	execute(server *Server)
}

type getCommand struct {
	periodOfTime string
	replyChan    chan PeriodicSensorsData
}

func (cmd getCommand) execute(server *Server) {
	cmd.replyChan <- server.getMeasurementsResults(cmd.periodOfTime)
}

type setCommand struct {
	lastMeasurement int
	sensorId        int
}

func (cmd setCommand) execute(server *Server) {
	server.addMeasurement(cmd.sensorId, cmd.lastMeasurement)
}
