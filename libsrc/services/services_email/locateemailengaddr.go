package services_email

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/emaildef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/httpdef"
	"os"
)

var locateEmailEngFlag = false
var addrInfo emaildef.EmailEngAddrInfo

func locateEmailEngAddr(EmailEngAddr *httpdef.HttpServerAddr) int {
	if !locateEmailEngFlag {
		configFile := fmt.Sprintf("%s/config/email/emaileng_addr.xml", globaldef.GetAppBaseDir())
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
		locateEmailEngFlag = true
	}
	EmailEngAddr.ServerIpAddr = addrInfo.IpAddr
	EmailEngAddr.ServerPort = addrInfo.Port
	EmailEngAddr.ServerTimeout = addrInfo.Timeout
	return 1
}
