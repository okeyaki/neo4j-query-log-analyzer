package model

type Record struct {
	Timestamp      string
	Level          string
	Time           int
	PlanningTime   int
	CPUTime        int
	WaitingTime    int
	AllocatedBytes int
	PageHits       int
	PageFaults     int
	SessionType    string
	Query          string
}
