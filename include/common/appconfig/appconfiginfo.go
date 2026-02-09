package appconfig

import "encoding/xml"

type AppGlobalConfig struct {
	Text          string `xml:",chardata"`
	NodeName      string `xml:"node_name"`
	DBEncryptFlag string `xml:"db_encrypt_flag"`
}

type AppServiceConfig struct {
	Text                string `xml:",chardata"`
	Node                string `xml:"node"`
	Name                string `xml:"name"`
	ServiceTimeOut      int    `xml:"service_timeout"`
	ServiceAliveTimeOut int    `xml:"service_alive_timeout"`
	ServiceDebugLevel   string `xml:"service_debug_level"`
	TLSFlag             string `xml:"tls_flag"`
}

type AppServiceGroupConfig struct {
	Text                       string             `xml:",chardata"`
	ServiceDefaultTimeOut      int                `xml:"service_default_timeout"`
	ServiceDefaultAliveTimeOut int                `xml:"service_default_alive_timeout"`
	Service                    []AppServiceConfig `xml:"service"`
}

type AppConfigInfoStruct struct {
	XMLName      xml.Name              `xml:"App"`
	Text         string                `xml:",chardata"`
	GlobalConfig AppGlobalConfig       `xml:"global_config"`
	ServiceGroup AppServiceGroupConfig `xml:"service_group"`
}
