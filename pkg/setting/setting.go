package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type CommonConf struct {
	RunMode string
}

type AppConf struct {
	PageSize  int
	JwtSecret int
}

type ServerConf struct {
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
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

var Common = &CommonConf{}

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

	mapTo("", Common)
	mapTo("app", App)
	mapTo("server", Server)
	mapTo("database", Database)
}

// 将配置信息映射到对应的变量
func mapTo(section string, v interface{}) {
	e := cfg.Section(section).MapTo(v)
	if e != nil {
		log.Printf("Failed to parse conf %s: %s", e, section)
	}
}
