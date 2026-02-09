package securitydef

import "encoding/xml"

type SecurityEngAddrInfo struct {
	XMLName xml.Name `xml:"security_eng_addr"`
	Text    string   `xml:",chardata"`
	IpAddr  string   `xml:"ipaddr"`
	Port    int      `xml:"port"`
	Timeout int      `xml:"timeout"`
}
