package resp

import (
	"strconv"
)

func WrapSimpleStringRESP(simpleString string) string {
	if len(simpleString) == 0 {
		return ""
	}
	return "+" + simpleString + "\r\n"
}

func WrapBulkStringRESP(s string) string {
	return "$" + strconv.Itoa(len(s)) + "\r\n" + s + "\r\n"
}

func GetNullBulkStringRESP() string {
	return "$-1\r\n"
}
