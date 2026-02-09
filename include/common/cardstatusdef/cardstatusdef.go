package cardstatusdef

import "lmsapieng/include/common/globaldef"

type cardStatusInfo struct {
	CardStatus     string
	CardStatusDesc string
}

const (
	SEM_CardPersoDownloadStatus_Downloaded    = "Downloaded"
	SEM_CardPersoDownloadStatus_NotDownloaded = "NotDownloaded"
)

const (
	SEM_CardStatus_WaitingForAuth     = "001"
	SEM_CardStatus_WaitingForPerso    = "002"
	SEM_CardStatus_WaitingForRcvCard  = "003"
	SEM_CardStatus_WaitingForIssuance = "004"
	SEM_CardStatus_InActive           = "005"
	SEM_CardStatus_FirstUseCard       = "006"
	SEM_CardStatus_Active             = "007"
	SEM_CardStatus_TempBlock          = "008"
	SEM_CardStatus_Block              = "009"
	SEM_CardStatus_Closed             = "010"
	SEM_CardStatus_Renewal            = "011"
	SEM_CardStatus_ReIssue            = "012"
	SEM_CardStatus_RePIN              = "013"
	SEM_CardStatus_WaitingForLink     = "014"
)

var cardStatusTable = []cardStatusInfo{
	{SEM_CardStatus_WaitingForAuth, "WaitingForAuth"},
	{SEM_CardStatus_WaitingForPerso, "WaitingForPerso"},
	{SEM_CardStatus_WaitingForRcvCard, "WaitingForRcvCard"},
	{SEM_CardStatus_WaitingForIssuance, "WaitingForIssuance"},
	{SEM_CardStatus_InActive, "InActive"},
	{SEM_CardStatus_FirstUseCard, "FirstUseCard"},
	{SEM_CardStatus_Active, "Active"},
	{SEM_CardStatus_TempBlock, "TempBlock"},
	{SEM_CardStatus_Block, "Blocked"},
	{SEM_CardStatus_Closed, "Closed"},
	{SEM_CardStatus_Renewal, "Renewal"},
	{SEM_CardStatus_ReIssue, "ReIssue"},
	{SEM_CardStatus_RePIN, "RePIN"},
	{SEM_CardStatus_WaitingForLink, "WaitingForLink"},
}

func GetCardStatusDesc(cardStatus string) string {
	rejectOffset := -1
	for i := 0; i < len(cardStatusTable); i++ {
		if cardStatus == cardStatusTable[i].CardStatus {
			rejectOffset = i
			break
		}
	}
	if rejectOffset < 0 {
		return globaldef.NOT_INITIALIZED
	}
	return cardStatusTable[rejectOffset].CardStatusDesc
}
