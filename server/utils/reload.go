package utils

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func Reload() error {
	if runtime.GOOS == "windows" {
		return errors.New("The system does not support")
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	return cmd.Run()
}
