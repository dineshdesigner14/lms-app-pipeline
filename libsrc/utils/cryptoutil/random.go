package cryptoutil

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenRandomNumber(Len int) (int, string) {
	randLenStr := "1"
	for i := 0; i < Len; i++ {
		randLenStr += "0"
	}
	randomRange, _ := strconv.Atoi(randLenStr)
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(randomRange)
	return 1, fmt.Sprintf("%0*d", Len, randNumber)
}

func GenRandomHex(Len int) (int, string) {
	bytes := make([]byte, Len/2)
	if _, err := rand.Read(bytes); err != nil {
		return -1, ""
	}
	return 1, strings.ToUpper(hex.EncodeToString(bytes))
}
