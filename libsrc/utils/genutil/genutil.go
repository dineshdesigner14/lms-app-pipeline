package genutil

import (
	"lmsapieng/include/common/servicedef"
	"runtime"
	"strings"
)

var moduleName, moduleVersion, serviceName string

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func SetModule(ServiceName string, ModuleName string, ModuleVersion string) {
	serviceName = ServiceName
	moduleName = ModuleName
	moduleVersion = ModuleVersion
}

func GetModuleName() string {
	return moduleName
}

func GetModuleVersion() string {
	return moduleVersion
}

func GetListeningPort() (int, string) {
	ProcessNameSlice := strings.Split(serviceName, servicedef.ServicePortSeparator)
	if len(ProcessNameSlice) != 2 {
		return -1, ""
	}
	return 1, ProcessNameSlice[1]
}

func GetListeningPortFromServiceName(ServiceName string) string {
	ProcessNameSlice := strings.Split(ServiceName, servicedef.ServicePortSeparator)
	return ProcessNameSlice[1]
}
