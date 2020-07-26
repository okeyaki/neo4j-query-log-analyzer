package model

type Report struct {
	QueryAggs []*QueryAgg
}

type QueryAgg struct {
	Query string

	Time         *QueryTimeAgg
	PlanningTime *QueryTimeAgg
	CPUTime      *QueryTimeAgg
	WaitingTime  *QueryTimeAgg

	Records []*Record
}

type QueryTimeAgg struct {
	Max   int
	Mean  int
	Total int
}
