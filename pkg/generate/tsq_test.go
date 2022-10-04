package generate_test

import (
	"os"
	"path"
	"testing"

	"github.com/nurmanhabib/go-tsa-client/pkg/generate"

	"github.com/stretchr/testify/require"
)

func TestTimestampRequest(t *testing.T) {
	input := path.Join("..", "tests/output", "file.txt")
	output := path.Join("..", "tests/output", "file.tsq")

	data := []byte("Hello world")
	err := os.WriteFile(input, data, 0644)

	require.NoError(t, err)

	errTSR := generate.TimestampQuery(input, output)

	require.NoError(t, errTSR)
}
