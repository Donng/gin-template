package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
)

type AppConf struct {
	RunMode string
}

type ServerConf struct {
	HttpPort int
}

type DatabaseConf struct {
	Type        string
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	TablePrefix string
}

var App = &AppConf{}
var Server = &ServerConf{}
var Database = &DatabaseConf{}

// 全局的配置对象
var cfg *ini.File

func init() {
	var e error
	cfg, e = ini.Load("conf/app.ini")
	if e != nil {
		log.Printf("Failed to load file app.ini: %s", e)
	}

	mapTo("", App)
	mapTo("server", Server)
	mapTo("database", Database)
	fmt.Println(Server, Database)
}

// 将配置信息映射到对应的变量
func mapTo(section string, v interface{}) {
	e := cfg.Section(section).MapTo(v)
	if e != nil {
		log.Printf("Failed to parse conf %s: %s", e, section)
	}
}
