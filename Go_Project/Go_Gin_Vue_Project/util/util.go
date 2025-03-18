package util

import (
	"math/rand"
	"time"
)

// RandmString 生成指定长度的随机字符串
func RandmString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM123456789") // 可用字符集
	result := make([]byte, n)                                                             // 初始化结果切片

	rand.Seed(time.Now().Unix()) // 设置随机种子
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))] // 随机选择一个字符
	}
	return string(result) // 返回生成的字符串
}
