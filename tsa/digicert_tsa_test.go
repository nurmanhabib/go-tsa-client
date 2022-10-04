package tsa_test

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"testing"

	"github.com/nurmanhabib/go-tsa-client/generate"
	"github.com/nurmanhabib/go-tsa-client/tsa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDigiCertClient_TSARequest(t *testing.T) {
	input := path.Join("..", "tests/output", "file.txt")
	outputTSQ := path.Join("..", "tests/output", "file.tsq")
	outputTSR := path.Join("..", "tests/output", "file.tsr")

	data := []byte("Hello world")
	err := os.WriteFile(input, data, 0644)
	require.NoError(t, err)

	errTSR := generate.TimestampQuery(input, outputTSQ)
	require.NoError(t, errTSR)

	tsq, errTSQ := os.ReadFile(outputTSQ)
	require.NoError(t, errTSQ)

	client := tsa.NewDigiCertClient()
	response, errReq := client.TSARequest(tsq)
	require.NoError(t, errReq)

	f, errF := os.Create(outputTSR)
	require.NoError(t, errF)

	defer f.Close()

	_, errW := f.Write(response)
	require.NoError(t, errW)

	verify, err := generate.TSRVerify(outputTSR)

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
}
