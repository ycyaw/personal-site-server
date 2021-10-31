package utils

import (
	"encoding/json"
	"os"
	"personal-site/log"
)

type MysqlConf struct{
	Addr string `json:"addr"`
	Port string `json:"port"`
	Database string `json:"database"`
	User string `json:"user"`
	Password string `json:"password"`
}

// 可配置项结构体
type Configuration struct {
	Addr string	`json:"addr"`
	Port string `json:"port"`
	MysqlConf MysqlConf `json:"mysqlConf"`
}

// 配置数据
var Config Configuration

var path string="config.json"

func init() {
	LoadConfig(path)
}
// 加载json文件配置
func LoadConfig(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		log.Error("conf file open error :",err.Error())
	}

	// 创建json解析器
	decoder := json.NewDecoder(file)

	// json解析并将数据填入Configuration中
	err = decoder.Decode(&Config)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("配置参数: ",Config)
}
