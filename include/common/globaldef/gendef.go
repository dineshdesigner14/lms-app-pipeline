package globaldef

import (
	"os"
	"strings"
)

const (
	NOT_INITIALIZED string = "NA"
)

const (
	OS_TYPE_LINUX   string = "LINUX"
	OS_TYPE_WINDOWS string = "WINDOWS"
)

const (
	STATUS_NOT_OK     string = "0"
	STATUS_OK         string = "1"
	STATUS_NOT_OK_STR string = "IN-ACTIVE"
	STATUS_OK_STR     string = "ACTIVE"
)

const (
	APPBASEDIR string = "SEMBASE"
)

const (
	EXIT_INIT_LOG_FAILED int = 2
	EXIT_INIT_FAILED     int = 3
	EXIT_NORMAL          int = 1
	EXIT_LOGINIT_FAILED  int = 4
)

const (
	ReqCtxObj = "ReqCtxObj"
)

func IsAppBaseDirExists() bool {
	_, ok := os.LookupEnv(APPBASEDIR)
	return ok
}

func GetAppBaseDir() string {
	BasePath, _ := os.LookupEnv(APPBASEDIR)
	return BasePath
}

func GetLinuxProcessName() string {
	str := strings.Split(os.Args[0], "/")
	return str[len(str)-1]
}
