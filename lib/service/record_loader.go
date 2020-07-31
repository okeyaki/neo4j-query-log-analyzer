package service

import (
	"bufio"
	"io"
	"os"

	"github.com/okeyaki/neo4j-query-log-analyzer/lib/model"
	"github.com/pkg/errors"
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

	reader := bufio.NewReader(l.normalizer.Run(os.Stdin))
	for {
		raw, err := l.read(reader)
		if err != nil {
			if err == io.EOF {
				break
			}

			return recs, err
		}

		rec, err := l.parser.Run(raw)
		if err != nil {
			return recs, err
		}
		if rec != nil {
			recs = append(recs, rec)
		}
	}

	return recs, nil
}

func (l *RecordLoader) read(reader *bufio.Reader) (string, error) {
	var bRaw []byte
	for {
		buf, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return "", err
			}
			return "", errors.WithStack(err)
		}

		bRaw = append(bRaw, buf[:]...)

		if !isPrefix {
			break
		}
	}

	raw := string(bRaw)

	return raw, nil
}
