package resp

import (
	"fmt"
	// "reflect"
	// "sort"
	"log"
	"strconv"
	"strings"
)

const (
	Error        = '-'
	Integer      = ':'
	SimpleString = '+'
	BulkString   = '$'
	Array        = '*' // Array is a special case, it's not a type of RESP, but it's used to represent an array of RESP values.
)

func StartDeserializeParser(byteArray []byte) []string {
	// fmt.Println("Start deserialize parser")

	var ans []string

	if len(byteArray) == 0 {
		return []string{}
	}
	if byteArray[0] == Array {
		ans = DeserializeArray(byteArray)
	} else if byteArray[0] == SimpleString {
		ans = DeserializeSimpleString(byteArray)
	}
	return ans
}

func DeserializeArray(byteArray []byte) []string {

	var ans []string

	if len(byteArray) == 0 {
		return []string{}
	}

	strByteArray := string(byteArray)
	for i := 4; i < len(byteArray); i++ {
		if byteArray[i] == BulkString {
			// log.Println("Desrializng : ", string(byteArray[i:]))
			index := strings.Index(strByteArray[i:], "\r\n")
			length := string(byteArray[i+1 : i+index])
			intValue, _ := strconv.Atoi(length)
			start := i + 3 + len(length)
			end := start + intValue
			ans = append(ans, string(byteArray[start:end]))
			i = end
		} else if byteArray[i] == Integer {
			index, err := DeserializeInteger(byteArray, i, &ans)
			if err != nil {
				log.Println("Error deserializing integer: ", err)
				return []string{}
			}
			i = index
		}
	}
	return ans
}

func DeserializeInteger(byteArray []byte, index int, ans *[]string) (int, error) {
	if len(byteArray) == 0 || byteArray[index] != ':' {
		return 0, fmt.Errorf("invalid RESP integer format")
	}

	var temp []byte

	for i := (index + 1); i < len(byteArray); i++ {
		if byteArray[i] != '\r' {
			temp = append(temp, byteArray[i])
		} else if byteArray[i] == '\r' {
			valueStr := string(temp)
			fmt.Println("temp: ", valueStr)
			*ans = append(*ans, valueStr)
			return i, nil
		}
	}

	return index, nil
}

func DeserializeSimpleString(byteArray []byte) []string {
	var ans []string
	var temp []byte
	for i := 1; i < len(byteArray); i++ {
		if byteArray[i] != '\r' {
			temp = append(temp, byteArray[i])
		} else {
			break
		}
	}
	ans = append(ans, string(temp))

	return ans
}

// Wrapers
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
