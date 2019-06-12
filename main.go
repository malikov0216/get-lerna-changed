package execCmdOutput

import (
	"os/exec"
	"strings"
)
func ExecuteLerna () (string, error) {
	out, err := exec.Command("lerna", "changed").Output()
	if err != nil {
		return "[Error in executing]", err
	}
	a := string(out)
	b := strings.Split(a, "\n")
	return b[0], nil
}
