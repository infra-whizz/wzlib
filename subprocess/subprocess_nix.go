package wzlib_subprocess

import (
	"os"
	"os/exec"
	"strings"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/sirupsen/logrus"
)

// ExecCommand is a small platform specific wrapper around os/exec.Command
func ExecCommand(name string, arg ...string) *Cmd {
	wzlib_logger.GetTextLogger(logrus.DebugLevel, os.Stderr).Debugf("Call: %s %s", name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	return newCmd(cmd)
}
