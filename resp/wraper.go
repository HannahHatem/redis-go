package resp

import (
	"strconv"
	"strings"
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

func WrapArrayRESP(arr []string) string {
	var builder strings.Builder
	builder.WriteString("*" + strconv.Itoa(len(arr)) + "\r\n")
	for _, s := range arr {
		builder.WriteString(WrapBulkStringRESP(s))
	}
	return builder.String()
}

func GetNullBulkStringRESP() string {
	return "$-1\r\n"
}
