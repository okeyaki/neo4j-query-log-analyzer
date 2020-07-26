package service

import (
	"strconv"
	"strings"

	"github.com/okeyaki/neo4j-query-log-analyzer/lib/model"
	"github.com/pkg/errors"
)

type RecordParser struct {
}

func NewRecordParser() *RecordParser {
	return &RecordParser{}
}

func (p *RecordParser) Run(raw string) (*model.Record, error) {
	switch {
	case strings.Contains(raw, "bolt-session"):
		return p.parseBolt(raw)
	case strings.Contains(raw, "embedded-session"):
		return p.parseEmbedded(raw)
	case strings.Contains(raw, "server-session"):
		return p.parseServer(raw)
	default:
		return nil, errors.Errorf("failed to parse record - %s", raw)
	}
}

func (p *RecordParser) parseEmbedded(raw string) (*model.Record, error) {
	flds := strings.Split(raw, "\t")

	firstFld, err := p.parseFirstField(flds[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rec := &model.Record{
		Query:          flds[len(flds)-1],
		Time:           firstFld.Time,
		PlanningTime:   firstFld.PlanningTime,
		CPUTime:        firstFld.CPUTime,
		WaitingTime:    firstFld.WaitingTime,
		AllocatedBytes: firstFld.AllocatedBytes,
		PageHits:       firstFld.PageHits,
		PageFaults:     firstFld.PageFaults,
		SessionType:    firstFld.SessionType,
	}
	return rec, nil
}

func (p *RecordParser) parseBolt(raw string) (*model.Record, error) {
	flds := strings.Split(raw, "\t")

	firstFld, err := p.parseFirstField(flds[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rec := &model.Record{
		Query:          flds[len(flds)-1],
		Time:           firstFld.Time,
		PlanningTime:   firstFld.PlanningTime,
		CPUTime:        firstFld.CPUTime,
		WaitingTime:    firstFld.WaitingTime,
		AllocatedBytes: firstFld.AllocatedBytes,
		PageHits:       firstFld.PageHits,
		PageFaults:     firstFld.PageFaults,
		SessionType:    firstFld.SessionType,
	}
	return rec, nil
}

func (p *RecordParser) parseServer(raw string) (*model.Record, error) {
	flds := strings.Split(raw, "\t")

	firstFld, err := p.parseFirstField(flds[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rec := &model.Record{
		Query:          flds[len(flds)-1],
		Time:           firstFld.Time,
		PlanningTime:   firstFld.PlanningTime,
		CPUTime:        firstFld.CPUTime,
		WaitingTime:    firstFld.WaitingTime,
		AllocatedBytes: firstFld.AllocatedBytes,
		PageHits:       firstFld.PageHits,
		PageFaults:     firstFld.PageFaults,
		SessionType:    firstFld.SessionType,
	}
	return rec, nil
}

func (p *RecordParser) parseFirstField(field string) (*firstField, error) {
	subFlds := strings.Split(field, " ")

	fld := &firstField{}

	tm, err := strconv.Atoi(subFlds[4])
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fld.Time = tm

	planningTm, err := strconv.Atoi(strings.TrimSuffix(subFlds[7], ","))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fld.PlanningTime = planningTm

	cpuTm, err := strconv.Atoi(strings.TrimSuffix(subFlds[9], ","))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fld.CPUTime = cpuTm

	waitingTm, err := strconv.Atoi(strings.TrimSuffix(subFlds[11], ")"))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fld.WaitingTime = waitingTm

	allocatedBytes, err := strconv.Atoi(subFlds[13])
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fld.AllocatedBytes = allocatedBytes

	hits, err := strconv.Atoi(subFlds[16])
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fld.PageHits = hits

	faults, err := strconv.Atoi(subFlds[19])
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fld.PageFaults = faults

	fld.SessionType = subFlds[21]

	return fld, nil
}

type firstField struct {
	Time           int
	PlanningTime   int
	CPUTime        int
	WaitingTime    int
	AllocatedBytes int
	PageHits       int
	PageFaults     int
	SessionType    string
}
