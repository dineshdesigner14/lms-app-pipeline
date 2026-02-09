package microsv_urmg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/httpdef"
	"lmsapieng/include/common/urmgdef"
	"os"
)

var locateUrmgApiEngFlag = false
var addrInfo urmgdef.UrmgApiEngAddrInfo

func locateUrmgApiEngAddr(UrmgApiEngAddr *httpdef.HttpServerAddr) int {
	if !locateUrmgApiEngFlag {
		configFile := fmt.Sprintf("%s/config/urmgapieng/urmgapieng_addr.xml", globaldef.GetAppBaseDir())
		xmlFile, err := os.Open(configFile)
		if err != nil {
			//trace.Lg("os.Open() failed for configFile(%s)...locateUrmgApiEngAddr() Failed", configFile)
			return -1
		}
		defer xmlFile.Close()
		byteValue, _ := ioutil.ReadAll(xmlFile)
		err = xml.Unmarshal(byteValue, &addrInfo)
		if err != nil {
			//trace.Lg("xml.Unmarshal() failed for configFile(%s)...loadGlobalDBConfig() Failed", configFile)
			return -1
		}
		locateUrmgApiEngFlag = true
	}
	UrmgApiEngAddr.ServerIpAddr = addrInfo.IpAddr
	UrmgApiEngAddr.ServerPort = addrInfo.Port
	UrmgApiEngAddr.ServerTimeout = addrInfo.Timeout
	return 1
}
