package args

import (
	"os"
	"strconv"
)

func Exists(key string) bool {
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == key {
			return true
		}
	}
	return false
}

func GetValue(key string) string {
	for i := 1; i < len(os.Args)-1; i++ {
		if os.Args[i] == key {
			return os.Args[i+1]
		}
	}
	return ""
}

func GetValueList(key string) []string {
	valueList := make([]string, 0)
	for i := 1; i < len(os.Args)-1; i++ {
		if os.Args[i] == key {
			valueList = append(valueList, os.Args[i+1])
		}
	}
	return valueList
}

func GetInt(key string) int {
	str := GetValue(key)
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return val
}
