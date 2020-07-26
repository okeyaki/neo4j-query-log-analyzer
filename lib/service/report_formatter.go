package service

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/okeyaki/neo4j-query-log-analyzer/lib/model"
	"github.com/pkg/errors"
)

type ReportFormatter struct {
}

func NewReportFormatter() *ReportFormatter {
	return &ReportFormatter{}
}

func (f *ReportFormatter) Run(rep *model.Report) (string, error) {
	aggs := rep.QueryAggs
	sort.Slice(aggs, func(i int, j int) bool {
		return aggs[i].Time.Total > aggs[j].Time.Total
	})

	qrs := []string{}
	for _, agg := range aggs {
		bEnc, err := json.Marshal(struct {
			Query             string `json:"query"`
			Count             int    `json:"count"`
			TimeTotal         int    `json:"timeTotal"`
			TimeMax           int    `json:"timeMax"`
			TimeMean          int    `json:"timeMean"`
			PlanningTimeTotal int    `json:"planningTimeTotal"`
			PlanningTimeMax   int    `json:"planningTimeMax"`
			PlanningTimeMean  int    `json:"planningTimeMean"`
			CPUTimeTotal      int    `json:"cpuTimeTotal"`
			CPUTimeMax        int    `json:"cpuTimeMax"`
			CPUTimeMean       int    `json:"cpuTimeMean"`
			WaitingTimeTotal  int    `json:"waitingTimeTotal"`
			WaitingTimeMax    int    `json:"waitingTimeMax"`
			WaitingTimeMean   int    `json:"waitingTimeMean"`
		}{
			Query:             agg.Query,
			Count:             len(agg.Records),
			TimeTotal:         agg.Time.Total,
			TimeMax:           agg.Time.Max,
			TimeMean:          agg.Time.Mean,
			CPUTimeTotal:      agg.CPUTime.Total,
			CPUTimeMax:        agg.CPUTime.Max,
			CPUTimeMean:       agg.CPUTime.Mean,
			PlanningTimeTotal: agg.PlanningTime.Total,
			PlanningTimeMax:   agg.PlanningTime.Max,
			PlanningTimeMean:  agg.PlanningTime.Mean,
			WaitingTimeTotal:  agg.WaitingTime.Total,
			WaitingTimeMax:    agg.WaitingTime.Max,
			WaitingTimeMean:   agg.WaitingTime.Mean,
		})
		if err != nil {
			return "", errors.WithStack(err)
		}

		qrs = append(qrs, string(bEnc))
	}

	out := strings.Join(qrs, "\n")

	return out, nil
}
