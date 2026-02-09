package nodeconfig

import "encoding/xml"

type NodeInfo struct {
	Text                         string `xml:",chardata"`
	Name                         string `xml:"name,attr"`
	Ipaddr                       string `xml:"ipaddr"`
	ServiceEngPort               int    `xml:"service_eng_port"`
	ServiceEngTimeout            int    `xml:"service_eng_timeout"`
	ServiceEngAliveProbeInterval int    `xml:"service_eng_alive_probe_interval"`
	DiagEngIpaddr                string `xml:"diag_eng_ipaddr"`
	DiagEngPort                  int    `xml:"diag_eng_port"`
	DiagEngTimeout               int    `xml:"diag_eng_timeout"`
	DBLoaderIpaddr               string `xml:"db_loader_ipaddr"`
	DBLoaderPort                 int    `xml:"db_loader_port"`
	DBLoaderTimeout              int    `xml:"db_loader_timeout"`
}

type NodeListInfo struct {
	XMLName xml.Name   `xml:"node_list"`
	Text    string     `xml:",chardata"`
	Node    []NodeInfo `xml:"node"`
}
