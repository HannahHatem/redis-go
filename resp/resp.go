package resp

import (
	"fmt"
	// "reflect"
	// "sort"
	"strconv"
	// "strings"
)

const (
	Error        = '-'
	Integer      = ':'
	SimpleString = '+'
	BulkString   = '$'
	Array        = '*' // Array is a special case, it's not a type of RESP, but it's used to represent an array of RESP values.
)

func StartDeserializeParser(byteArray []byte) []string {
	fmt.Println("Start deserialize parser")

	var ans []string

	if len(byteArray) == 0 {
		return []string{}
	}
	if byteArray[0] == Array {
		ans = DeserializeArray(byteArray)
		fmt.Println(ans)
	}
	if byteArray[0] == SimpleString {
		var temp []byte
		for i := 1; i < len(byteArray); i++ {
			if byteArray[i] != '\r' {
				temp = append(temp, byteArray[i])
			} else {
				break
			}
		}
		ans = append(ans, string(temp))
	}
	return ans
}

func DeserializeArray(byteArray []byte) []string {

	var ans []string

	fmt.Println("Deserialize array")
	if len(byteArray) == 0 {
		return []string{}
	}
	//size := byteArray[1]
	fmt.Println("hello 1 ")
	for i := 4; i < len(byteArray); i++ {
		if byteArray[i] == BulkString {
			fmt.Println("BulkString")
			length := string(byteArray[i+1])
			intValue, _ := strconv.Atoi(length)
			start := i + 4
			end := start + intValue
			ans = append(ans, string(byteArray[start:end]))
			i = end
		} else if byteArray[i] == SimpleString {
			i++
			var temp []byte
			for byteArray[i] != '\r' {
				temp = append(temp, byteArray[i])
				i++
			}
			ans = append(ans, string(temp))
		}
	}
	return ans
}
