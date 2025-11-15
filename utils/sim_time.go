package utils

import (
	"sync"
	"time"

	"github.com/fajryhamzah/go-loan-sim/constants"
)

type SimulationTime struct {
	currentDate time.Time
}

var (
	once    sync.Once
	simTime *SimulationTime
)

func (s *SimulationTime) Now() time.Time {
	return s.currentDate
}

func (s *SimulationTime) AddWeek() time.Time {
	return s.AddDay(7)
}

func (s *SimulationTime) AddDay(day int) time.Time {
	s.currentDate = s.currentDate.AddDate(0, 0, day)
	return s.Now()
}

func GetSimulationTime() *SimulationTime {
	once.Do(func() {
		simTime = &SimulationTime{
			currentDate: time.Now(),
		}
	})

	return simTime
}

func Format(time time.Time) string {
	return time.Format(constants.TIME_FORMAT)
}

func Now() time.Time { // for simulation sake, we could change it to regular time.Time based on env
	return GetSimulationTime().Now()
}

func NowFormatted() string {
	return Format(GetSimulationTime().Now())
}
