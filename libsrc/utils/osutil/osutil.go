package osutil

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func IsLinux() bool {
	os := runtime.GOOS
	if os == "linux" {
		return true
	}
	return false
}

func InitOS() int {
	var rval int
	os := runtime.GOOS
	switch os {
	case "windows":
		rval = performWindowsInit()
		break
	case "darwin":
		rval = performMACInit()
		break
	case "linux":
		rval = performLinuxInit()
		break
	default:
		rval = performLinuxInit()
	}
	return rval
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

func performWindowsInit() int {
	return 1
}

func performMACInit() int {
	return 1
}

func performLinuxInit() int {
	os.Stdin.Close()
	os.Stdout.Close()
	os.Stderr.Close()
	ret, _, errNo := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if errNo != 0 {
		fmt.Printf("\n syscall.RawSyscall failed for syscall.SYS_FORK with errNo(%d)\n", errNo)
		return -1
	}
	if ret != 0 {
		os.Exit(0)
	}
	_, err := syscall.Setsid()
	if err != nil {
		fmt.Printf("\n syscall.Setsid() failed with err(%s)\n", err)
		return -1
	}
	return 1
}
