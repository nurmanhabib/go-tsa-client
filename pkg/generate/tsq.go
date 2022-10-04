package generate

import (
	"os"
	"os/exec"
)

func TimestampQuery(input, output string) error {
	cmd := exec.Command("openssl", "ts", "-query", "-data", input, "-cert", "-sha256", "-no_nonce", "-out", output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
