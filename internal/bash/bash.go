package bash

import (
	"fmt"
	"os/exec"
	"strings"
)

var UploadFormat = "clockify-cli manual --interactive=0 --project='%s' --task='%s' --when='%s' --when-to-close='%s' --description='%s'\n"

func BashExec(format string, a ...any) (string, error) {
	cmd := fmt.Sprintf(format, a...)

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", fmt.Errorf("could not bash exec: %s", err)
	}

	return strings.TrimRight(string(out), "\n"), nil
}
