package seqdef

import "encoding/xml"

type SeqEngAddrInfo struct {
	XMLName xml.Name `xml:"seq_eng_addr"`
	Text    string   `xml:",chardata"`
	IpAddr  string   `xml:"ipaddr"`
	Port    int      `xml:"port"`
	Timeout int      `xml:"timeout"`
}
