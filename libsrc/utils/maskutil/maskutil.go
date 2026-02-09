package maskutil

func GetMaskStr(MaskStr string, displayNum1 int, displayNum2 int) string {
	var MaskMaskStr string
	var maskdigits int
	if len(MaskStr) >= displayNum1 {
		MaskMaskStr = MaskStr[:displayNum1]
		if len(MaskStr) > 10 {
			maskdigits = len(MaskStr) - displayNum1 - displayNum2
			for i := 0; i < maskdigits; i++ {
				MaskMaskStr += "*"
			}
			MaskMaskStr += MaskStr[len(MaskStr)-displayNum2:]
		} else {
			MaskMaskStr = MaskStr
		}
	} else {
		MaskMaskStr = MaskStr
	}
	return MaskMaskStr
}
