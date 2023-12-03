package helpers

import (
	"strconv"
)

func ContainsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ContainsByte(slice []byte, item byte) bool {
	for _, b := range slice {
		if b == item {
			return true
		}
	}
	return false
}

func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func ByteIsInt(b byte) bool {
	_, err := strconv.Atoi(string(b))
	return err == nil
}