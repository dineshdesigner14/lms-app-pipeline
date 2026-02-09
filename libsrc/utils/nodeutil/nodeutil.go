package nodeutil

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/nodeconfig"
	"os"
)

var nodeConfigLoadedFlag = false
var nodeListInfo nodeconfig.NodeListInfo

func GetNodeInfo(NodeName string, nodeInfo *nodeconfig.NodeInfo) int {
	var rval = -1
	if !nodeConfigLoadedFlag {
		if loadNodeConfig() < 0 {
			return rval
		}
	}
	for i := 0; i < len(nodeListInfo.Node); i++ {
		if nodeListInfo.Node[i].Name == NodeName {
			*nodeInfo = nodeListInfo.Node[i]
			rval = 1
			break
		}
	}
	if rval < 0 {
		return rval
	}
	return rval
}

func loadNodeConfig() int {
	if nodeConfigLoadedFlag {
		return 1
	}
	configFile := fmt.Sprintf("%s/config/rt/node.xml", globaldef.GetAppBaseDir())
	xmlFile, err := os.Open(configFile)
	if err != nil {
		//trace.Lg("os.Open() failed for configFile(%s)", configFile)
		return -1
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &nodeListInfo)
	if err != nil {
		//trace.Lg("xml.Unmarshal() failed for configFile(%s)", configFile)
		return -1
	}
	nodeConfigLoadedFlag = true
	return 1
}

func GetNodeListInfo() (int, nodeconfig.NodeListInfo) {
	if !nodeConfigLoadedFlag {
		if loadNodeConfig() < 0 {
			return -1, nodeListInfo
		}
	}
	return 1, nodeListInfo
}
