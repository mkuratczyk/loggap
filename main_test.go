package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	in := strings.NewReader(`
2023-12-01 17:14:16.017759+00:00 [info] line 1
2023-12-01 17:14:16.117759+00:00 [info] line 2
2023-12-01 17:14:16.017759+00:00 [info] line 3
2023-12-01 17:14:17.007759+00:00 [info] line 4
2023-12-01 17:14:18.017759+00:00 [info] line 5
2023-12-01 17:14:18.017759+00:00 [info] line 6
`)

	expected_out := `
2023-12-01 17:14:16.017759+00:00 [info] line 1
2023-12-01 17:14:16.117759+00:00 [info] line 2
2023-12-01 17:14:16.017759+00:00 [info] line 3
.......... 990ms later
2023-12-01 17:14:17.007759+00:00 [info] line 4
.......... 1.01s later
2023-12-01 17:14:18.017759+00:00 [info] line 5
2023-12-01 17:14:18.017759+00:00 [info] line 6
`
	var buf bytes.Buffer
	out := io.Writer(&buf)

	processLogs(in, out, 500*time.Millisecond)

	assert.Equal(t, expected_out, buf.String())

}
