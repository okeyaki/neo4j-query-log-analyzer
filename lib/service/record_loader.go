package service

import (
	"bufio"
	"io"
	"os"

	"github.com/okeyaki/neo4j-query-log-analyzer/lib/model"
)

type RecordLoader struct {
	normalizer *RecordNormalizer
	parser     *RecordParser
}

func NewRecordLoader() *RecordLoader {
	return &RecordLoader{
		normalizer: NewRecordNormalizer(),
		parser:     NewRecordParser(),
	}
}

func (l *RecordLoader) Run() ([]*model.Record, error) {
	recs := []*model.Record{}

	scanner := bufio.NewScanner(l.normalizer.Run(os.Stdin))
	for scanner.Scan() {
		rec, err := l.parser.Run(scanner.Text())
		if err != nil {
			continue
		}
		if rec != nil {
			recs = append(recs, rec)
		}
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			return recs, err
		}
	}

	return recs, nil
}
