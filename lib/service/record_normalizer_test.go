package service_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/okeyaki/neo4j-query-log-analyzer/lib/service"
	"github.com/stretchr/testify/assert"
)

func Test_RecordNormalizer_CanNormalizeRecord(t *testing.T) {
	recNormalizer := service.NewRecordNormalizer()

	dataset := []struct {
		src      io.Reader
		expected string
	}{
		{
			src:      bytes.NewBufferString("foo \n bar \n baz {}"),
			expected: "foo bar baz {}\n",
		},
	}
	for i, data := range dataset {
		dest := recNormalizer.Run(data.src)

		actual, err := ioutil.ReadAll(dest)

		assert.NoErrorf(t, err, "%d", i)
		assert.Equalf(t, data.expected, string(actual), "%d", i)
	}
}
