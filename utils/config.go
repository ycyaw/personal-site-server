package utils

import (
	"encoding/json"
	"os"
	"personal-site/log"
)


type Configuration struct {
	Addr string
	Port string
}

var Config Configuration

func LoadConfig(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		log.Error(err.Error())
	}

	// 创建json解析器
	decoder := json.NewDecoder(file)

	// json解析并将数据填入Configuration中
	err = decoder.Decode(&Config)
	if err != nil {
		log.Error(err.Error())
	}
}