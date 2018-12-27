package lib

import (
	"crypto/sha256"
	"strconv"
	"strings"
)

func EncryptSha256(input string) string {
	h := sha256.Sum256([]byte(input))
	var hex strings.Builder
	for _, hc := range h {
		hs := strconv.FormatInt(int64(0xff&hc), 16)
		if len(hs) == 1 {
			hex.WriteString("0")
		}
		hex.WriteString(hs)
	}
	return hex.String()
}
