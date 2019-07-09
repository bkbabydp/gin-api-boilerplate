package util

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"os"
	"reflect"
	"runtime"
)

// Assert 断言 err != nil
func Assert(err error, msg interface{}) {
	if err != nil {
		panic(msg)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$@")

// RandomStr 生成随机字符串
func RandomStr(length int) string {
	var lenthLetter = len(letterRunes)

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(lenthLetter)]
	}
	return string(b)
}

// Base64Encoding base64 编码
func Base64Encoding(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

// Base64Decoding base64 解码
func Base64Decoding(enc string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// FuncName 获取函数的名字
func FuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// AssignMap 合并多个map
func AssignMap(maps ...map[string]interface{}) map[string]interface{} {
	m := map[string]interface{}{}

	for _, mp := range maps {
		for key, val := range mp {
			m[key] = val
		}
	}

	return m
}

// 把能转成字符串的都转成JSON字符串
func ToString(v interface{}) string {
	bt, _ := json.Marshal(v)
	return string(bt)
}

// HasFile 是否存在该文件
func HasFile(f string) bool {
	if _, err := os.Stat(f); !os.IsNotExist(err) {
		return true
	}
	return false
}
