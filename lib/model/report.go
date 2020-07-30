package model

type Report struct {
	QueryAggs []*QueryAgg
}

type QueryAgg struct {
	Query string

	Time         *IntAgg
	PlanningTime *IntAgg
	CPUTime      *IntAgg
	WaitingTime  *IntAgg
	PageHit      *IntAgg
	PageFault    *IntAgg

	Records []*Record
}

type IntAgg struct {
	Max   int
	Mean  int
	Total int
}
