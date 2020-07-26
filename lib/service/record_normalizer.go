package service

import (
	"bufio"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/tidwall/transform"
)

type RecordNormalizer struct {
}

func NewRecordNormalizer() *RecordNormalizer {
	return &RecordNormalizer{}
}

func (n *RecordNormalizer) Run(src io.Reader) io.Reader {
	src = n.Trim(src)
	src = n.Merge(src)

	return src
}

func (n *RecordNormalizer) Trim(src io.Reader) io.Reader {
	reader := bufio.NewReader(src)

	return transform.NewTransformer(func() ([]byte, error) {
		raw, err := n.read(reader)
		if err != nil {
			if err == io.EOF {
				return nil, err
			}
			return nil, errors.WithStack(err)
		}

		raw = strings.TrimSpace(raw)

		return []byte(raw + "\n"), nil
	})
}

func (n *RecordNormalizer) Merge(src io.Reader) io.Reader {
	reader := bufio.NewReader(src)

	return transform.NewTransformer(func() ([]byte, error) {
		base, err := n.read(reader)
		if err != nil {
			if err == io.EOF {
				return nil, err
			}
			return nil, errors.WithStack(err)
		}

		for !strings.HasSuffix(base, "{}") {
			next, err := n.read(reader)
			if err != nil {
				if err == io.EOF {
					return nil, err
				}
				return nil, errors.WithStack(err)
			}

			base += " " + next
		}

		return []byte(base + "\n"), nil
	})
}

func (n *RecordNormalizer) read(reader *bufio.Reader) (string, error) {
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
