package service_test

import (
	"strings"
	"testing"

	"github.com/okeyaki/neo4j-query-log-analyzer/lib/model"
	"github.com/okeyaki/neo4j-query-log-analyzer/lib/service"
	"github.com/stretchr/testify/assert"
)

func Test_RecordParser_CanParseRecord(t *testing.T) {
	recParser := service.NewRecordParser()

	dataset := []struct {
		raw      string
		expected *model.Record
	}{
		{
			raw: strings.Join([]string{
				"1000-01-01 00:00:00.000+0000 INFO  6 ms: (planning: 3, cpu: 2, waiting: 1) - 0 B - 0 page hists, 0 page faults - embedded-session",
				"",
				"",
				"",
				"- MATCH (a:` Arbitrary label name that really doesn't matter `) RETURN a LIMIT 0 - {} - {}",
			}, "\t"),
			expected: &model.Record{
				Query:        "- MATCH (a:` Arbitrary label name that really doesn't matter `) RETURN a LIMIT 0 - {} - {}",
				Time:         6,
				PlanningTime: 3,
				CPUTime:      2,
				WaitingTime:  1,
			},
		},
		{
			raw: strings.Join([]string{
				"1000-01-01 00:00:00.000+0000 INFO  6 ms: (planning: 3, cpu: 2, waiting: 1) - 0 B - 0 page hists, 0 page faults - bolt-session	bolt",
				"neo4j",
				"neo4j-javascript/1.0.0",
				"",
				"client/127.0.0.1:42",
				"server/127.0.0.1:42>",
				"neo4j - MATCH (f:Foo) RETURN f; - {} - {}",
			}, "\t"),
			expected: &model.Record{
				Query:        "neo4j - MATCH (f:Foo) RETURN f; - {} - {}",
				Time:         6,
				PlanningTime: 3,
				CPUTime:      2,
				WaitingTime:  1,
			},
		},
		{
			raw: strings.Join([]string{
				"1000-01-01 00:00:00.000+0000 INFO  6 ms: (planning: 3, cpu: 2, waiting: 1) - 0 B - 0 page hists, 0 page faults - server-session",
				"http",
				"127.0.0.1",
				"/db/data/transaction/commit",
				"neo4j - MATCH (f:Foo) RETURN f; - {} - {}",
			}, "\t"),
			expected: &model.Record{
				Query:        "neo4j - MATCH (f:Foo) RETURN f; - {} - {}",
				Time:         6,
				PlanningTime: 3,
				CPUTime:      2,
				WaitingTime:  1,
			},
		},
	}
	for i, data := range dataset {
		rec, err := recParser.Run(data.raw)

		assert.NoErrorf(t, err, "%d", i)
		assert.Equal(t, data.expected.Query, rec.Query, "%d", i)
		assert.Equal(t, data.expected.Time, rec.Time, "%d", i)
	}
}
