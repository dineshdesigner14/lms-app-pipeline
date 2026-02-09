package urmgdef

import "encoding/xml"

type UrmgApiEngAddrInfo struct {
	XMLName xml.Name `xml:"urmg_api_eng_addr"`
	Text    string   `xml:",chardata"`
	IpAddr  string   `xml:"ipaddr"`
	Port    int      `xml:"port"`
	Timeout int      `xml:"timeout"`
}

const (
	UserDataTableKey        = "user_data_table"
	VerifyTokenMicroService = "VerifyToken"
)
