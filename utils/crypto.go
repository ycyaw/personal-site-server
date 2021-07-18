package utils

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"personal-site/log"
)

// 将字符串使用md5加密
func EncodeMd5(password string) string {
	// 加密器
	hash := md5.New()

	// MD5加密
	hashByte := hash.Sum([]byte(password))

	// 将数据转换为16进制
	return fmt.Sprintf("%x", hashByte)
}

// 生成token
func EncodeToken() string {
	// token
	token := make([]byte, 32)

	// 生成token
	_, err := rand.Read(token)
	if err != nil {
		log.Warning(err.Error())
	}

	return fmt.Sprintf("%x", token)
}