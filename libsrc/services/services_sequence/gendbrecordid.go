package services_sequence

import (
	"fmt"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/rtutil"
	"os"
	"sync"
)

var dbrecSeqNumMutex = &sync.Mutex{}
var dbrecSeqNum = 0

func getdbrecSeqNum() string {
	dbrecSeqNumMutex.Lock()
	defer dbrecSeqNumMutex.Unlock()
	if dbrecSeqNum >= 999999 {
		dbrecSeqNum = 1
	} else {
		dbrecSeqNum += 1
	}
	return fmt.Sprintf("%06d", dbrecSeqNum)
}

func GetDBRecordID() string {
	nodeName := rtutil.GetCurrentNodeName()
	if len(nodeName) == 0 {
		return nodeName
	}
	dbRecordNum := fmt.Sprintf("%s_%d_%s_%s", nodeName, os.Getpid(), dtutil.GetDateTimeVal(), getdbrecSeqNum())
	return dbRecordNum
}
