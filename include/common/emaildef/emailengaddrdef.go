package emaildef

import "encoding/xml"

type EmailEngAddrInfo struct {
	XMLName xml.Name `xml:"email_eng_addr"`
	Text    string   `xml:",chardata"`
	IpAddr  string   `xml:"ipaddr"`
	Port    int      `xml:"port"`
	Timeout int      `xml:"timeout"`
}
