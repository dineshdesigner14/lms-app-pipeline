package reqbrokerutil

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/headerdef"
	"lmsapieng/include/common/nodeconfig"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/include/common/servicedef"
	"lmsapieng/libsrc/utils/nodeutil"
	"os"
	"strconv"
	"strings"
)

var brokerSelectorInfo reqbrokerdef.ReqBrokerSelectorStruct
var loadBrokerSelectorInfoFlag = false

func GetBrokerSelectionArgs(brokerMsg []byte, brokerSelectionArgs *string, rejectDesc *string) int {
	brokerReqMap := make(map[string]interface{})
	err := json.Unmarshal(brokerMsg, &brokerReqMap)
	if err != nil {
		*rejectDesc = fmt.Sprintf("json unmarshal failed for %s", "brokerReqMap")
		return -1
	}

	_, ok := brokerReqMap[reqbrokerdef.ReqBrokerReqHeaderJSONObj]
	if !ok {
		*rejectDesc = fmt.Sprintf("no %s json in request", reqbrokerdef.ReqBrokerReqHeaderJSONObj)
		return -1
	}
	header_data_map := make(map[string]interface{})
	header_data_map = brokerReqMap[reqbrokerdef.ReqBrokerReqHeaderJSONObj].(map[string]interface{})
	lbrokerSelectionArgs := fmt.Sprintf("%s=%s,", reqbrokerdef.ReqBrokerSelectorModuleType, header_data_map[headerdef.App_Header_Broker_MsgSrcType])
	lbrokerSelectionArgs += fmt.Sprintf("%s=%s,", reqbrokerdef.ReqBrokerSelectorModuleName, header_data_map[headerdef.App_Header_Broker_MsgSrc])
	lbrokerSelectionArgs += fmt.Sprintf("%s=%s", reqbrokerdef.ReqBrokerSelectorChannel, header_data_map[headerdef.App_Header_Broker_Channel])
	*brokerSelectionArgs = lbrokerSelectionArgs
	return 1
}

func GetBrokerInstance(brokerSelectionArgs string, brokerInstance *[]reqbrokerdef.ReqBrokerInstanceStruct) int {
	var nodeInfo nodeconfig.NodeInfo
	if !loadBrokerSelectorInfoFlag {
		if loadBrokerSelectorInfo() < 0 {
			return -1
		}
	}
	var selectorCriteria reqbrokerdef.ReqBrokerSelectorCriteriaStruct
	var lbrokerInstance reqbrokerdef.ReqBrokerInstanceStruct
	selectorCriteria.ModuleType = "*"
	selectorCriteria.ModuleName = "*"
	selectorCriteria.Channel = "*"
	selectorCriteria.ReqType = "*"
	selectorCriteria.EntityID = "*"
	selectorCriteria.SubentityID = "*"

	brokerSelectionArgsSlice := strings.Split(brokerSelectionArgs, ",")
	for i := 0; i < len(brokerSelectionArgsSlice); i++ {
		selectorArg := strings.TrimSpace(brokerSelectionArgsSlice[i])
		selectionParamSlice := strings.Split(selectorArg, "=")
		if len(selectionParamSlice) >= 2 {
			if selectionParamSlice[0] == reqbrokerdef.ReqBrokerSelectorModuleType {
				selectorCriteria.ModuleType = selectionParamSlice[1]
			} else if selectionParamSlice[0] == reqbrokerdef.ReqBrokerSelectorModuleName {
				selectorCriteria.ModuleName = selectionParamSlice[1]
			} else if selectionParamSlice[0] == reqbrokerdef.ReqBrokerSelectorChannel {
				selectorCriteria.Channel = selectionParamSlice[1]
			} else if selectionParamSlice[0] == reqbrokerdef.ReqBrokerSelectorReqType {
				selectorCriteria.ReqType = selectionParamSlice[1]
			} else if selectionParamSlice[0] == reqbrokerdef.ReqBrokerSelectorEntityID {
				selectorCriteria.EntityID = selectionParamSlice[1]
			} else if selectionParamSlice[0] == reqbrokerdef.ReqBrokerSelectorSubEntityID {
				selectorCriteria.SubentityID = selectionParamSlice[1]
			}
		}
	}
	found := false
	offset := 0
	if selectorCriteria.ModuleType != "*" {
		for i := 0; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType {
				found = true
				offset = i
				break
			}
		}
		if !found {
			selectorCriteria.ModuleType = "*"
		}
	}
	if !found {
		for i := 0; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType {
				found = true
				offset = i
				break
			}
		}
		if !found {
			//trace.Lg("GetBrokerInstance() failed for ModuleType:%s", selectorCriteria.ModuleType)
			return -1
		}
	}
	found = false
	if selectorCriteria.ModuleName != "*" {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName {
				found = true
				offset = i
				break
			}
		}
		if !found {
			selectorCriteria.ModuleName = "*"
		}
	}
	if !found {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName {
				found = true
				offset = i
				break
			}
		}
		if !found {
			//trace.Lg("GetBrokerInstance() failed for ModuleType:%s ModuleName:%s", selectorCriteria.ModuleType, selectorCriteria.ModuleName)
			return -1
		}
	}
	found = false
	if selectorCriteria.Channel != "*" {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.Channel == selectorCriteria.Channel {
				found = true
				offset = i
				break
			}
		}
		if !found {
			selectorCriteria.Channel = "*"
		}
	}
	if !found {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.Channel == selectorCriteria.Channel {
				found = true
				offset = i
				break
			}
		}
		if !found {
			//trace.Lg("GetBrokerInstance() failed for ModuleType:%s ModuleName:%s Channel:%s", selectorCriteria.ModuleType, selectorCriteria.ModuleName, selectorCriteria.Channel)
			return -1
		}
	}
	found = false
	if selectorCriteria.ReqType != "*" {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.Channel == selectorCriteria.Channel &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ReqType == selectorCriteria.ReqType {
				found = true
				offset = i
				break
			}
		}
		if !found {
			selectorCriteria.ReqType = "*"
		}
	}
	if !found {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.Channel == selectorCriteria.Channel &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ReqType == selectorCriteria.ReqType {
				found = true
				offset = i
				break
			}
		}
		if !found {
			//trace.Lg("GetBrokerInstance() failed for ModuleType:%s ModuleName:%s Channel:%s ReqType:%s", selectorCriteria.ModuleType, selectorCriteria.ModuleName, selectorCriteria.Channel, selectorCriteria.ReqType)
			return -1
		}
	}
	found = false
	if selectorCriteria.EntityID != "*" {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.Channel == selectorCriteria.Channel &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ReqType == selectorCriteria.ReqType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.EntityID == selectorCriteria.EntityID {
				found = true
				offset = i
				break
			}
		}
		if !found {
			selectorCriteria.EntityID = "*"
		}
	}
	if !found {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.Channel == selectorCriteria.Channel &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ReqType == selectorCriteria.ReqType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.EntityID == selectorCriteria.EntityID {
				found = true
				offset = i
				break
			}
		}
		if !found {
			//trace.Lg("GetBrokerInstance() failed for ModuleType:%s ModuleName:%s Channel:%s ReqType:%s EntityID:%s", selectorCriteria.ModuleType, selectorCriteria.ModuleName, selectorCriteria.Channel, selectorCriteria.ReqType, selectorCriteria.EntityID)
			return -1
		}
	}
	found = false
	if selectorCriteria.SubentityID != "*" {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.Channel == selectorCriteria.Channel &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ReqType == selectorCriteria.ReqType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.EntityID == selectorCriteria.EntityID &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.SubentityID == selectorCriteria.SubentityID {
				if nodeutil.GetNodeInfo(brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerNode, &nodeInfo) < 0 {
					//trace.Lg("GetNodeInfo() failed for Node:%s", brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerNode)
					continue
				}
				brokerPortSlice := strings.Split(brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerInstance, servicedef.ServicePortSeparator)
				if len(brokerPortSlice) != 2 {
					//trace.Lg("Cannot get broker port from instance:%s", brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerInstance)
					continue
				}
				brokerPort, err := strconv.Atoi(brokerPortSlice[1])
				if err != nil {
					//trace.Lg("Cannot get broker port from instance:%s", brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerInstance)
					continue
				}
				found = true
				offset = i
				lbrokerInstance.BrokerModule = brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerInstance
				lbrokerInstance.IpAddr = nodeInfo.Ipaddr
				lbrokerInstance.PortNum = brokerPort
				lbrokerInstance.TimeOut = brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.Timeout
				*brokerInstance = append(*brokerInstance, lbrokerInstance)
			}
		}
		if !found {
			selectorCriteria.SubentityID = "*"
		}
	}
	if !found {
		for i := offset; i < len(brokerSelectorInfo.ReqBrokerInstance); i++ {
			if brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleType == selectorCriteria.ModuleType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ModuleName == selectorCriteria.ModuleName &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.Channel == selectorCriteria.Channel &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.ReqType == selectorCriteria.ReqType &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.EntityID == selectorCriteria.EntityID &&
				brokerSelectorInfo.ReqBrokerInstance[i].SelectorCriteria.SubentityID == selectorCriteria.SubentityID {
				if nodeutil.GetNodeInfo(brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerNode, &nodeInfo) < 0 {
					//trace.Lg("GetNodeInfo() failed for Node:%s", brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerNode)
					continue
				}
				brokerPortSlice := strings.Split(brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerInstance, servicedef.ServicePortSeparator)
				if len(brokerPortSlice) != 2 {
					//trace.Lg("Cannot get broker port from instance:%s", brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerInstance)
					continue
				}
				brokerPort, err := strconv.Atoi(brokerPortSlice[1])
				if err != nil {
					//trace.Lg("Cannot get broker port from instance:%s", brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerInstance)
					continue
				}
				found = true
				offset = i
				lbrokerInstance.BrokerModule = brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.BrokerInstance
				lbrokerInstance.IpAddr = nodeInfo.Ipaddr
				lbrokerInstance.PortNum = brokerPort
				lbrokerInstance.TimeOut = brokerSelectorInfo.ReqBrokerInstance[i].AddrInfo.Timeout
				*brokerInstance = append(*brokerInstance, lbrokerInstance)
				break
			}
		}
		if !found {
			//trace.Lg("GetBrokerInstance() failed for ModuleType:%s ModuleName:%s Channel:%s ReqType:%s EntityID:%s SubentityID:%s", selectorCriteria.ModuleType, selectorCriteria.ModuleName, selectorCriteria.Channel, selectorCriteria.ReqType, selectorCriteria.EntityID, selectorCriteria.SubentityID)
			return -1
		}
	}
	return 1
}

func loadBrokerSelectorInfo() int {
	selectorFile := fmt.Sprintf("%s/config/reqbroker/reqbrokerinstance.xml", globaldef.GetAppBaseDir())
	xmlFile, err := os.Open(selectorFile)
	if err != nil {
		//trace.Lg("os.Open() failed for configFile(%s)", selectorFile)
		return -1
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &brokerSelectorInfo)
	if err != nil {
		//trace.Lg("xml.Unmarshal() failed for configFile(%s)", selectorFile)
		return -1
	}
	loadBrokerSelectorInfoFlag = true
	return 1
}
