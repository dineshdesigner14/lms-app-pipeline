package rtutil

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/appconfig"
	"lmsapieng/include/common/debugdef"
	"lmsapieng/include/common/globaldef"
	"os"
	"strings"
)

var appConfig appconfig.AppConfigInfoStruct
var loadAppConfigFlag = false

func LoadAppConfig() int {
	configFile := fmt.Sprintf("%s/config/rt/app.xml", globaldef.GetAppBaseDir())
	xmlFile, err := os.Open(configFile)
	if err != nil {
		//trace.Lg("os.Open() failed for configFile(%s)", configFile)
		return -1
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &appConfig)
	if err != nil {
		//trace.Lg("xml.Unmarshal() failed for configFile(%s)", configFile)
		return -1
	}
	loadAppConfigFlag = true
	return 1
}

func GetCurrentNodeName() string {
	if !loadAppConfigFlag {
		if LoadAppConfig() < 0 {
			return ""
		}
	}
	return appConfig.GlobalConfig.NodeName
}

func IsDBEncryptReq() bool {
	if !loadAppConfigFlag {
		if LoadAppConfig() < 0 {
			return false
		}
	}
	if strings.EqualFold(appConfig.GlobalConfig.DBEncryptFlag, "Yes") {
		return true
	}
	return false
}

func GetServiceListInfo() []appconfig.AppServiceConfig {
	if !loadAppConfigFlag {
		if LoadAppConfig() < 0 {
			return nil
		}
	}
	return appConfig.ServiceGroup.Service
}

func GetServiceTimeOut(serviceName string) int {
	if !loadAppConfigFlag {
		if LoadAppConfig() < 0 {
			return -1
		}
	}
	serviceTimeOut := 0
	for i := 0; i < len(appConfig.ServiceGroup.Service); i++ {
		if appConfig.ServiceGroup.Service[i].Name == serviceName {
			serviceTimeOut = appConfig.ServiceGroup.Service[i].ServiceTimeOut
			break
		}
	}
	if serviceTimeOut != 0 {
		return serviceTimeOut
	}
	if appConfig.ServiceGroup.ServiceDefaultTimeOut != 0 {
		return appConfig.ServiceGroup.ServiceDefaultTimeOut
	}
	return 30
}

func GetServiceAliveTimeOut(serviceName string) int {
	if !loadAppConfigFlag {
		if LoadAppConfig() < 0 {
			return -1
		}
	}
	serviceAliveTimeOut := 0
	for i := 0; i < len(appConfig.ServiceGroup.Service); i++ {
		if appConfig.ServiceGroup.Service[i].Name == serviceName {
			serviceAliveTimeOut = appConfig.ServiceGroup.Service[i].ServiceAliveTimeOut
			break
		}
	}
	if serviceAliveTimeOut != 0 {
		return serviceAliveTimeOut
	}
	if appConfig.ServiceGroup.ServiceDefaultAliveTimeOut != 0 {
		return appConfig.ServiceGroup.ServiceDefaultAliveTimeOut
	}
	return 30
}

func GetServiceDebugLevel(serviceName string) int {
	if !loadAppConfigFlag {
		if LoadAppConfig() < 0 {
			return -1
		}
	}
	var debugLevel string
	var traceLevel int
	var rval = -1
	for i := 0; i < len(appConfig.ServiceGroup.Service); i++ {
		if appConfig.ServiceGroup.Service[i].Name == serviceName {
			debugLevel = appConfig.ServiceGroup.Service[i].ServiceDebugLevel
			rval = 1
			break
		}
	}
	if rval < 0 {
		traceLevel = debugdef.DEBUG_LEVEL_SECURED
	} else {
		if strings.EqualFold(debugLevel, debugdef.DEBUG_LEVEL_NORMAL_STR) {
			traceLevel = debugdef.DEBUG_LEVEL_NORMAL
		} else if strings.EqualFold(debugLevel, debugdef.DEBUG_LEVEL_SECURED_STR) {
			traceLevel = debugdef.DEBUG_LEVEL_SECURED
		} else if strings.EqualFold(debugLevel, debugdef.DEBUG_LEVEL_TEST_STR) {
			traceLevel = debugdef.DEBUG_LEVEL_TEST
		} else if strings.EqualFold(debugLevel, debugdef.DEBUG_LEVEL_ERROR_STR) {
			traceLevel = debugdef.DEBUG_LEVEL_ERROR
		} else {
			traceLevel = debugdef.DEBUG_LEVEL_SECURED
		}
	}
	return traceLevel
}

func IsServiceTLSFlagSet(serviceName string) bool {
	if !loadAppConfigFlag {
		if LoadAppConfig() < 0 {
			return false
		}
	}
	tlsFlag := false
	for i := 0; i < len(appConfig.ServiceGroup.Service); i++ {
		if appConfig.ServiceGroup.Service[i].Name == serviceName {
			if strings.EqualFold(appConfig.ServiceGroup.Service[i].TLSFlag, "Yes") {
				tlsFlag = true
			}
			break
		}
	}
	return tlsFlag
}
