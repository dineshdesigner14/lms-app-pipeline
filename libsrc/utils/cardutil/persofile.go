package cardutil

import (
	"fmt"
	"lmsapieng/include/common/carddef"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/mathutil"
	"strings"
)

func leftPadWithZeros(strVal string, strSize int) string {
	if len(strVal) >= strSize {
		return strVal[:strSize]
	}
	return fmt.Sprintf("%0*s", strSize, strVal)
}

func rightPadWithSpaces(strVal string, strSize int) string {
	if len(strVal) >= strSize {
		return strVal[:strSize]
	}
	return fmt.Sprintf("%-*s", strSize, strVal)
}

func replaceSpaceAndTab(s string) string {
	var builder strings.Builder

	for _, char := range s {
		if char == ' ' || char == '\t' {
			builder.WriteRune('/')
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func ComposeTrack1Data(CardNum string, CardName string, ExpiryDate string, ServiceCode string, CVV1 string, Track1DataSize int, track1Data *string) int {
	trackCardNameSize := 26
	track1Str := ""
	track1Str += "B"
	track1Str += CardNum
	track1Str += "^"
	track1Str += strings.ToUpper(rightPadWithSpaces(replaceSpaceAndTab(CardName), trackCardNameSize))
	track1Str += "^"
	track1Str += ExpiryDate
	track1Str += ServiceCode
	track1Str += "000000000000000"
	track1Str += CVV1
	track1Str += "000000"
	offsetLen := Track1DataSize - len(track1Str)
	track1Str += strings.Repeat(" ", offsetLen)
	*track1Data = track1Str
	return 1
}

func ComposeTrack2Data(CardNum string, ExpiryDate string, ServiceCode string, CVV1 string, serverFlag string, Track2DataSize int, track2Data *string) int {
	track2Str := ""
	track2Str += CardNum
	track2Str += "="
	track2Str += ExpiryDate
	track2Str += ServiceCode
	track2Str += "00000"
	track2Str += CVV1
	track2Str += "00"
	track2Str += serverFlag
	track2Str += "00"
	offsetLen := Track2DataSize - len(track2Str)
	track2Str += strings.Repeat(" ", offsetLen)
	*track2Data = track2Str
	return 1
}

func ComposeTrack3Data(Track3DataSize int, track3Data *string) int {
	track3Str := strings.Repeat(" ", 107)
	if len(track3Str) != Track3DataSize {
		//trace.Lg("track3StrLen[%d] != Track3DataSize[%d]", len(track3Str), Track3DataSize)
		return -1
	}
	*track3Data = track3Str
	return 1
}

func GetCardEmbRecord(cardEmbFileInfo carddef.CardEmbossingFileStruct, cardEmbRecordStr *string, rejectReason *string) int {
	cardTypeIDSize := 10
	embCardNumSize := 20
	embExpiryDateSize := 20
	cardNameSize := 26
	coBrandedNameSize := 26
	cvv2Size := 10
	track1Size := 79
	track2Size := 40
	track3Size := 107
	customerNameSize := 30
	addrLine1Size := 30
	addrLine2Size := 30
	addrLine3Size := 30
	addrLine4Size := 30
	deliveryMtdSize := 10
	trailingSpacesSize := 200

	if len(cardEmbFileInfo.CardTypeID) == 0 {
		*rejectReason = "CardTypeID is NULL"
		return -1
	}
	if len(cardEmbFileInfo.CardNum) == 0 {
		*rejectReason = "CardNum is NULL"
		return -1
	}
	if len(cardEmbFileInfo.CardNum) != 16 && len(cardEmbFileInfo.CardNum) != 19 {
		*rejectReason = "CardNum Length should 16 or 19"
		return -1
	}
	if !mathutil.IsValidNumericString(cardEmbFileInfo.CardNum) {
		*rejectReason = fmt.Sprintf("CardNum[%s] Should be Numeric", cardEmbFileInfo.CardNum)
		return -1
	}
	if len(cardEmbFileInfo.CardSeqNum) == 0 {
		*rejectReason = "CardSeqNum is NULL"
		return -1
	}
	if len(cardEmbFileInfo.CardSeqNum) != 2 {
		*rejectReason = "CardSeqNum Length should 2"
		return -1
	}
	if !mathutil.IsValidNumericString(cardEmbFileInfo.CardSeqNum) {
		*rejectReason = fmt.Sprintf("CardSeqNum[%s] Should be Numeric", cardEmbFileInfo.CardSeqNum)
		return -1
	}
	if len(cardEmbFileInfo.ExpiryDate) == 0 {
		*rejectReason = "ExpiryDate is NULL"
		return -1
	}
	if len(cardEmbFileInfo.ExpiryDate) != 4 {
		*rejectReason = "ExpiryDate Length should 4"
		return -1
	}
	if !dtutil.IsValidDateFormat(cardEmbFileInfo.ExpiryDate, "YYMM") {
		*rejectReason = fmt.Sprintf("ExpiryDate[%s] Should be YYMM", cardEmbFileInfo.ExpiryDate)
		return -1
	}
	if len(cardEmbFileInfo.CardName) == 0 {
		*rejectReason = "CardName is NULL"
		return -1
	}
	if len(cardEmbFileInfo.CoBrandedName) == 0 {
		*rejectReason = "CoBrandedName is NULL"
		return -1
	}
	if len(cardEmbFileInfo.CVV1) == 0 {
		*rejectReason = "CVV1 is NULL"
		return -1
	}
	if len(cardEmbFileInfo.CVV1) != 3 {
		*rejectReason = "CVV1 Length should 3"
		return -1
	}
	if !mathutil.IsValidNumericString(cardEmbFileInfo.CVV1) {
		*rejectReason = fmt.Sprintf("CVV1[%s] Should be Numeric", cardEmbFileInfo.CVV1)
		return -1
	}
	if len(cardEmbFileInfo.CVV2) == 0 {
		*rejectReason = "CVV2 is NULL"
		return -1
	}
	if len(cardEmbFileInfo.CVV2) != 3 {
		*rejectReason = "CVV2 Length should 3"
		return -1
	}
	if !mathutil.IsValidNumericString(cardEmbFileInfo.CVV2) {
		*rejectReason = fmt.Sprintf("CVV2[%s] Should be Numeric", cardEmbFileInfo.CVV2)
		return -1
	}
	if len(cardEmbFileInfo.ICVV) == 0 {
		*rejectReason = "ICVV is NULL"
		return -1
	}
	if len(cardEmbFileInfo.ICVV) != 3 {
		*rejectReason = "ICVV Length should 3"
		return -1
	}
	if !mathutil.IsValidNumericString(cardEmbFileInfo.ICVV) {
		*rejectReason = fmt.Sprintf("ICVV[%s] Should be Numeric", cardEmbFileInfo.ICVV)
		return -1
	}
	if len(cardEmbFileInfo.ServiceCode) == 0 {
		*rejectReason = "ServiceCode is NULL"
		return -1
	}
	if len(cardEmbFileInfo.ServiceCode) != 3 {
		*rejectReason = "ServiceCode Length should 3"
		return -1
	}
	if !mathutil.IsValidNumericString(cardEmbFileInfo.ServiceCode) {
		*rejectReason = fmt.Sprintf("ServiceCode[%s] Should be Numeric", cardEmbFileInfo.ServiceCode)
		return -1
	}
	if len(cardEmbFileInfo.CustomerName) == 0 {
		*rejectReason = "CustomerName is NULL"
		return -1
	}
	if len(cardEmbFileInfo.CardIssDate) == 0 {
		*rejectReason = "CardIssDate is NULL"
		return -1
	}
	if !dtutil.IsValidDateFormat(cardEmbFileInfo.CardIssDate, "DDMMYYYY") {
		*rejectReason = fmt.Sprintf("CardIssDate[%s] Should be DDMMYYYY", cardEmbFileInfo.CardIssDate)
		return -1
	}
	recordStr := ""
	recordStr += strings.ToUpper(leftPadWithZeros(cardEmbFileInfo.CardTypeID, cardTypeIDSize))
	recordStr += rightPadWithSpaces(cardEmbFileInfo.CardNum, embCardNumSize)
	recordStr += rightPadWithSpaces(cardEmbFileInfo.ExpiryDate, embExpiryDateSize)
	recordStr += strings.ToUpper(rightPadWithSpaces(cardEmbFileInfo.CardName, cardNameSize))
	recordStr += strings.ToUpper(rightPadWithSpaces(cardEmbFileInfo.CoBrandedName, coBrandedNameSize))
	recordStr += rightPadWithSpaces(cardEmbFileInfo.CVV2, cvv2Size)
	if len(cardEmbFileInfo.Track1) != track1Size {
		*rejectReason = fmt.Sprintf("Track1Size[%d] != RequiredSize[%d]", len(cardEmbFileInfo.Track1), track1Size)
		return -1
	}
	recordStr += cardEmbFileInfo.Track1
	if len(cardEmbFileInfo.Track2) != track2Size {
		*rejectReason = fmt.Sprintf("Track1Size[%d] != RequiredSize[%d]", len(cardEmbFileInfo.Track2), track2Size)
		return -1
	}
	recordStr += cardEmbFileInfo.Track2
	tempStr := ""
	if ComposeTrack3Data(track3Size, &tempStr) < 0 {
		*rejectReason = "composeTrack3Data Failed"
		return -1
	}
	recordStr += tempStr
	recordStr += strings.ToUpper(rightPadWithSpaces(cardEmbFileInfo.CustomerName, customerNameSize))
	recordStr += strings.Repeat(" ", addrLine1Size)
	recordStr += strings.Repeat(" ", addrLine2Size)
	recordStr += strings.Repeat(" ", addrLine3Size)
	recordStr += strings.Repeat(" ", addrLine4Size)
	recordStr += strings.Repeat(" ", deliveryMtdSize)
	recordStr += cardEmbFileInfo.ICVV
	recordStr += cardEmbFileInfo.CardIssDate
	recordStr += cardEmbFileInfo.CardSeqNum
	recordStr += strings.Repeat(" ", trailingSpacesSize)
	*cardEmbRecordStr = recordStr
	return 1
}
