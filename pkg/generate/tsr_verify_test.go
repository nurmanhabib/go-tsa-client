package generate_test

import (
	"bufio"
	"bytes"
	"fmt"
	"path"
	"testing"

	"github.com/nurmanhabib/go-tsa-client/pkg/generate"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTSRVerify(t *testing.T) {
	input := path.Join("..", "tests/output", "example.tsr")

	verify, err := generate.TSRVerify(input)

	require.NoError(t, err)
	require.NotEmpty(t, verify)

	verifyReader := bytes.NewReader(verify)
	sc := bufio.NewScanner(verifyReader)

	var line int

	for sc.Scan() {
		line++

		if line == 2 {
			assert.Equal(t, "Status: Granted.", sc.Text())
			break
		}
	}

	fmt.Printf("%s", verify)
}
