package services_security

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/httpdef"
	"lmsapieng/include/common/securitydef"
	"os"
)

var locateSecurityEngFlag = false
var addrInfo securitydef.SecurityEngAddrInfo

func locateSecurityEngAddr(SecurityEngAddr *httpdef.HttpServerAddr) int {
	if !locateSecurityEngFlag {
		configFile := fmt.Sprintf("%s/config/security/securityeng_addr.xml", globaldef.GetAppBaseDir())
		xmlFile, err := os.Open(configFile)
		if err != nil {
			//trace.Lg("os.Open() failed for configFile(%s)...loadGlobalDBConfig() Failed", configFile)
			return -1
		}
		defer xmlFile.Close()
		byteValue, _ := ioutil.ReadAll(xmlFile)
		err = xml.Unmarshal(byteValue, &addrInfo)
		if err != nil {
			//trace.Lg("xml.Unmarshal() failed for configFile(%s)...loadGlobalDBConfig() Failed", configFile)
			return -1
		}
		locateSecurityEngFlag = true
	}
	SecurityEngAddr.ServerIpAddr = addrInfo.IpAddr
	SecurityEngAddr.ServerPort = addrInfo.Port
	SecurityEngAddr.ServerTimeout = addrInfo.Timeout
	return 1
}
