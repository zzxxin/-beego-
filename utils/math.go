package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Sub returns the difference of two integers
func Sub(a, b int) int {
	return a - b
}

// InArray 检查给定值是否存在于数组中
func InArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// Str 将任何类型转换为字符串
func Str(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func JoinUintSlice(slice []uint) string {
	str := make([]string, len(slice))
	for i, v := range slice {
		str[i] = strconv.Itoa(int(v))
	}
	return strings.Join(str, ",")
}
