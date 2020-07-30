package service

import (
	"github.com/okeyaki/neo4j-query-log-analyzer/lib/model"
)

type ReportBuilder struct {
}

func NewReportBuilder() *ReportBuilder {
	return &ReportBuilder{}
}

func (b *ReportBuilder) Run(recs []*model.Record) (*model.Report, error) {
	aggs := []*model.QueryAgg{}
	for qr, qrRecs := range b.groupRecordsByQuery(recs) {
		aggs = append(aggs, &model.QueryAgg{
			Query:        qr,
			Records:      qrRecs,
			Time:         b.aggregateInt(qrRecs, func(rec *model.Record) int { return rec.Time }),
			PlanningTime: b.aggregateInt(qrRecs, func(rec *model.Record) int { return rec.PlanningTime }),
			CPUTime:      b.aggregateInt(qrRecs, func(rec *model.Record) int { return rec.CPUTime }),
			WaitingTime:  b.aggregateInt(qrRecs, func(rec *model.Record) int { return rec.WaitingTime }),
			PageHit:      b.aggregateInt(qrRecs, func(rec *model.Record) int { return rec.PageHits }),
			PageFault:    b.aggregateInt(qrRecs, func(rec *model.Record) int { return rec.PageFaults }),
		})
	}

	rep := &model.Report{
		QueryAggs: aggs,
	}
	return rep, nil
}

func (b *ReportBuilder) groupRecordsByQuery(recs []*model.Record) map[string][]*model.Record {
	qrs := map[string][]*model.Record{}
	for _, rec := range recs {
		qrs[rec.Query] = append(qrs[rec.Query], rec)
	}

	return qrs
}

func (b *ReportBuilder) aggregateInt(
	recs []*model.Record,
	get func(rec *model.Record) int,
) *model.IntAgg {
	total := 0
	max := 0
	for _, rec := range recs {
		tm := get(rec)

		total += tm

		if tm > max {
			max = tm
		}
	}

	return &model.IntAgg{
		Total: total,
		Max:   max,
		Mean:  total / len(recs),
	}
}
