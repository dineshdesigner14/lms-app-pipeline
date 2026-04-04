package osutil

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func IsLinux() bool {
	os := runtime.GOOS
	if os == "linux" {
		return true
	}
	return false
}

func InitOS() int {
	return 1
}

func GetServiceName() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return getWindowsServiceName()
	case "darwin":
		return getMACServiceName()
	case "linux":
		return getLinuxServiceName()
	default:
		return getLinuxServiceName()
	}
}

func getWindowsServiceName() string {
	str := strings.Split(os.Args[0], "/")
	return str[len(str)-1]
}

func getMACServiceName() string {
	str := strings.Split(os.Args[0], "/")
	return str[len(str)-1]
}

func getLinuxServiceName() string {
	str := strings.Split(os.Args[0], "/")
	return str[len(str)-1]
}

func IsProcessWithNameRunning(processName string, parentPid string) bool {
	cmd := fmt.Sprintf("ps -ef | grep %s | grep -v grep | grep -v vi", processName)
	out, e := exec.Command("bash", "-c", cmd).Output()
	if e != nil {
		return false
	}
	found := false
	output := string(out[:])
	str := strings.Split(output, "\n")
	for i := 0; i < len(str); i++ {
		strArray := strings.Fields(str[i])
		if len(strArray) >= 8 {
			if strArray[2] == parentPid {
				found = true
				break
			}
		}
	}
	return found
}
