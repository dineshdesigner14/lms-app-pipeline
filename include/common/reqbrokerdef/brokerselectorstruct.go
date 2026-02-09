package reqbrokerdef

import "encoding/xml"

type ReqBrokerSelectorCriteriaStruct struct {
	Text        string `xml:",chardata"`
	ModuleType  string `xml:"module_type"`
	ModuleName  string `xml:"module_name"`
	Channel     string `xml:"channel"`
	ReqType     string `xml:"req_type"`
	EntityID    string `xml:"entity_id"`
	SubentityID string `xml:"subentity_id"`
}

type ReqBrokerAddrStruct struct {
	Text           string `xml:",chardata"`
	BrokerInstance string `xml:"broker_instance"`
	BrokerNode     string `xml:"broker_node"`
	Timeout        int    `xml:"timeout"`
}

type ReqBrokerSelectorStruct struct {
	XMLName           xml.Name `xml:"req_broker_selector"`
	Text              string   `xml:",chardata"`
	ReqBrokerInstance []struct {
		SelectorCriteria ReqBrokerSelectorCriteriaStruct `xml:"req_broker_selection_criteria"`
		AddrInfo         ReqBrokerAddrStruct             `xml:"req_broker_addr"`
	} `xml:"req_broker_instance"`
}
