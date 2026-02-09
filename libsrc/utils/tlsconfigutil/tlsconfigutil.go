package tlsconfigutil

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/libsrc/utils/fileutil"
	"os"
)

type TLSConfigStruct struct {
	XMLName      xml.Name `xml:"tlsconfig"`
	Text         string   `xml:",chardata"`
	Serverconfig []struct {
		Text       string `xml:",chardata"`
		Ipaddr     string `xml:"ipaddr,attr"`
		Portnum    string `xml:"portnum,attr"`
		ServerName string `xml:"server_name,attr"`
		ServerCrt  string `xml:"server_crt"`
	} `xml:"serverconfig"`
}

var tlsConfigInfo TLSConfigStruct
var tlsConfigLoadFlag = false

func loadTLSConfig() int {
	configFile := fmt.Sprintf("%s/config/rt/tlsconfig.xml", globaldef.GetAppBaseDir())
	xmlFile, err := os.Open(configFile)
	if err != nil {
		//trace.Lg("os.Open() failed for configFile(%s)", configFile)
		return -1
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &tlsConfigInfo)
	if err != nil {
		//trace.Lg("xml.Unmarshal() failed for configFile(%s)", configFile)
		return -1
	}
	tlsConfigLoadFlag = true
	return 1
}

func IsTLSReq(serverIpAddr string, serverPort string, serverCrt *string, serverName *string) bool {
	// trace.Lg("IsTLSReq() called for serverIpAddr:%s serverPort:%s", serverIpAddr, serverPort)
	if !tlsConfigLoadFlag {
		if loadTLSConfig() < 0 {
			//trace.Lg("return false----1 for serverIpAddr:%s serverPort:%s", serverIpAddr, serverPort)
			return false
		}
	}
	srvrIp := ""
	if serverIpAddr == "localhost" {
		srvrIp = "127.0.0.1"
	} else {
		srvrIp = serverIpAddr
	}
	tlsConfigFound := false
	for i := 0; i < len(tlsConfigInfo.Serverconfig); i++ {
		if tlsConfigInfo.Serverconfig[i].Ipaddr == srvrIp && tlsConfigInfo.Serverconfig[i].Portnum == serverPort {
			if len(tlsConfigInfo.Serverconfig[i].ServerCrt) != 0 {
				serverCrtFile := fmt.Sprintf("%s/%s/%s", globaldef.GetAppBaseDir(), "config/tls/certificates/client", tlsConfigInfo.Serverconfig[i].ServerCrt)
				if fileutil.IsFileExists(serverCrtFile) {
					tlsConfigFound = true
					*serverCrt = tlsConfigInfo.Serverconfig[i].ServerCrt
					*serverName = tlsConfigInfo.Serverconfig[i].ServerName
				}
			}
			break
		}
	}
	//trace.Lg("return ftlsConfigFound:%t for serverIpAddr:%s serverPort:%s", tlsConfigFound, serverIpAddr, serverPort)
	return tlsConfigFound
}
