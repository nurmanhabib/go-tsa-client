package handler_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/nurmanhabib/go-tsa-client/interface/handler"
	generate2 "github.com/nurmanhabib/go-tsa-client/pkg/generate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDigicertTSAHandler(t *testing.T) {
	input := path.Join("../..", "tests/output", "tsa-handler-file.txt")
	output := path.Join("../..", "tests/output", "tsa-handler-file.tsq")

	data := []byte("Hello world")
	err := os.WriteFile(input, data, 0644)
	require.NoError(t, err)

	errTSR := generate2.TimestampQuery(input, output)
	require.NoError(t, errTSR)

	f, errRead := os.ReadFile(output)
	require.NoError(t, errRead)

	b := bytes.NewReader(f)
	req := httptest.NewRequest(http.MethodPost, "/digicert-tsa", b)
	w := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/timestamp-query")
	req.Header.Add("Accept", "application/timestamp-reply, application/timestamp-response")
	req.Header.Add("Pragma", "no-cache")

	// Handler execution
	handler.DigicertTSAHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, errResponse := ioutil.ReadAll(res.Body)
	require.NoError(t, errResponse)

	// Prepare to Verify TSR
	tsrPath, errTSR := ioutil.TempFile("", "test-*.tsr")
	require.NoError(t, errTSR)

	defer os.Remove(tsrPath.Name())

	_, errWrite := tsrPath.Write(data)
	require.NoError(t, errWrite)

	verify, err := generate2.TSRVerify(tsrPath.Name())

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
