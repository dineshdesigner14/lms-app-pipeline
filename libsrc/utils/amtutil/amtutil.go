package amtutil

import (
	"fmt"
	"lmsapieng/libsrc/utils/mathutil"
	"math/big"
	"strings"
)

func ConvertFloatToString(FloatVal float64, DecPos int, StringVal *string) int {
	ba := big.NewFloat(FloatVal)
	*StringVal = ba.Text('f', DecPos)
	return 1
}

func IsValidIsoAmt(IsoAmt string) bool {
	if !mathutil.IsValidNumericStr(IsoAmt, 12, 12) {
		return false
	}
	return true
}

func CompareAmt(val1 string, val2 string) int {
	ba, bb := big.NewFloat(0), big.NewFloat(0)
	if _, ok := ba.SetString(val1); !ok {
		ba.SetString("0")
	}
	if _, ok := bb.SetString(val2); !ok {
		bb.SetString("0")
	}
	return ba.Cmp(bb)
}

func AddAmt(val1 string, val2 string, DecPos int) string {
	ba, bb := big.NewFloat(0), big.NewFloat(0)
	if _, ok := ba.SetString(val1); !ok {
		ba.SetString("0")
	}
	if _, ok := bb.SetString(val2); !ok {
		bb.SetString("0")
	}
	sum := big.NewFloat(0).Add(ba, bb)
	return sum.Text('f', DecPos)
}

func SubAmt(val1 string, val2 string, DecPos int) string {
	ba, bb := big.NewFloat(0), big.NewFloat(0)
	if _, ok := ba.SetString(val1); !ok {
		ba.SetString("0")
	}
	if _, ok := bb.SetString(val2); !ok {
		bb.SetString("0")
	}
	sum := big.NewFloat(0).Sub(ba, bb)
	return sum.Text('f', DecPos)
}

func MulAmt(val1 string, val2 string, DecPos int) string {
	ba, bb := big.NewFloat(0), big.NewFloat(0)
	if _, ok := ba.SetString(val1); !ok {
		ba.SetString("0")
	}
	if _, ok := bb.SetString(val2); !ok {
		bb.SetString("0")
	}
	sum := big.NewFloat(0).Mul(ba, bb)
	return sum.Text('f', DecPos)
}

func DivAmt(val1 string, val2 string, DecPos int) string {
	ba, bb := big.NewFloat(0), big.NewFloat(0)
	if _, ok := ba.SetString(val1); !ok {
		ba.SetString("0")
	}
	if _, ok := bb.SetString(val2); !ok {
		bb.SetString("0")
	}
	sum := big.NewFloat(0).Quo(ba, bb)
	return sum.Text('f', DecPos)
}

func TransformIsoToInternalAmt(IsoAmt string, DecPos int) (int, string) {
	if !IsValidIsoAmt(IsoAmt) {
		return -1, IsoAmt
	}
	IsoAmtSlice := []rune(IsoAmt)
	IsoAmtWhole := string(IsoAmtSlice[:len(IsoAmtSlice)-DecPos])
	IsoAmtFraction := string(IsoAmtSlice[len(IsoAmtSlice)-DecPos:])
	IsoAmtString := IsoAmtWhole + "." + IsoAmtFraction
	IsoAmtFloat := big.NewFloat(0)
	IsoAmtFloat.SetString(IsoAmtString)
	return 1, IsoAmtFloat.Text('f', DecPos)
}

func TransformInternalToIsoAmt(AmountStr string, DecPos int) (int, string) {
	amountArr := strings.Split(AmountStr, ".")
	return 1, fmt.Sprintf("%010s%s", amountArr[0], amountArr[1])
}

func IsValidAmount(value string, MinWholeNumLen int, MaxWholeNumLen int, MinDecLen int, MaxDecLen int) bool {
	amountArr := strings.Split(value, ".")
	if len(amountArr) == 1 {
		if !mathutil.IsValidNumericString(amountArr[0]) {
			return false
		}
		return true
	}
	if len(amountArr) != 2 {
		return false
	}
	if len(amountArr[0]) == 0 || len(amountArr[1]) == 0 {
		return false
	}
	if len(amountArr[0]) < MinWholeNumLen || len(amountArr[0]) > MaxWholeNumLen {
		return false
	}
	if len(amountArr[1]) < MinDecLen || len(amountArr[1]) > MaxDecLen {
		return false
	}
	if !mathutil.IsValidNumericString(amountArr[0]) {
		return false
	}
	if !mathutil.IsValidNumericString(amountArr[1]) {
		return false
	}
	if amountArr[0][0] == '0' {
		if len(amountArr[0]) != 1 {
			return false
		}
	}
	return true
}
