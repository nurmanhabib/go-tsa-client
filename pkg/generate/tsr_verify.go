package generate

import (
	"bytes"
	"os"
	"os/exec"
)

func TSRVerify(input string) ([]byte, error) {
	b := new(bytes.Buffer)
	cmd := exec.Command("openssl", "ts", "-reply", "-in", input, "-text")
	cmd.Stdout = b
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
