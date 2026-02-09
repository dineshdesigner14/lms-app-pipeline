package sequtil

import (
	"fmt"
	"math/big"
	"strings"
)

func isSequenceZero(CurrentSequence string) bool {
	for i := 0; i < len(CurrentSequence); i++ {
		if CurrentSequence[i] != '0' {
			return false
		}
	}
	return true
}

func GetNextSequence(CurrentSequence string, SequenceLen int) (int, string) {
	var SeqStr, lStr string
	for i := 0; i < SequenceLen; i++ {
		lStr += "9"
	}
	if CurrentSequence == lStr {
		return -1, ""
	}
	if isSequenceZero(CurrentSequence) {
		SeqStr = "0"
	} else {
		SeqStr = strings.TrimLeft(CurrentSequence, "0")
	}
	ba, bb := big.NewInt(0), big.NewInt(0)
	if _, ok := ba.SetString(SeqStr, 0); !ok {
		return -1, ""
	}
	if _, ok := bb.SetString("1", 0); !ok {
		return -1, ""
	}
	sum := big.NewInt(0).Add(ba, bb)
	return 1, fmt.Sprintf("%0*s", SequenceLen, sum.String())
}

func GetNextSequenceWithOverflow(CurrentSequence string, SequenceLen int) (int, string) {
	var SeqStr, lStr string
	for i := 0; i < SequenceLen; i++ {
		lStr += "9"
	}
	if CurrentSequence == lStr {
		return 1, fmt.Sprintf("%0*s", SequenceLen, "1")
	}
	if isSequenceZero(CurrentSequence) {
		SeqStr = "0"
	} else {
		SeqStr = strings.TrimLeft(CurrentSequence, "0")
	}
	ba, bb := big.NewInt(0), big.NewInt(0)
	if _, ok := ba.SetString(SeqStr, 0); !ok {
		return -1, ""
	}
	if _, ok := bb.SetString("1", 0); !ok {
		return -1, ""
	}
	sum := big.NewInt(0).Add(ba, bb)
	return 1, fmt.Sprintf("%0*s", SequenceLen, sum.String())
}
