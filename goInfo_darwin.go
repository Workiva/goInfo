package goInfo

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func GetInfo() (*GoInfoObject, error) {
	out, err := _getInfo()
	if err != nil {
		return nil, err
	}

	for strings.Index(out, "broken pipe") != -1 {
		out, err = _getInfo()
		if err != nil {
			return nil, err
		}
		time.Sleep(500 * time.Millisecond)
	}
	osStr := strings.Replace(out, "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osInfo := strings.Split(osStr, " ")
	gio := &GoInfoObject{Kernel: osInfo[0], Core: osInfo[1], Platform: osInfo[2], OS: osInfo[0], GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
	gio.Hostname, _ = os.Hostname()

	return gio, nil
}

func _getInfo() (string, error) {
	cmd := exec.Command("uname", "-srm")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
