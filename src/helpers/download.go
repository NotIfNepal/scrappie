package helpers

import (
	"fmt"
	"os/exec"
)

func WgetFile(url string, fileName string) ([]byte, error) {
	cmd := exec.Command("wget", fmt.Sprintf("--output-document=%s", fileName), url)

	return cmd.Output()
}
