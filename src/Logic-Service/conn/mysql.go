package conn

import (
	"fmt"
	"github.com/gkzy/gow/lib/config"
	"github.com/gkzy/gow/lib/logy"
	"github.com/gkzy/gow/lib/mysql"
)

var (
	debug bool
	name  = "kimi"
)

// InitMysql
func InitMysql() {
	if config.DefaultString("run_mode", "dev") == "dev" {
		debug = true
	}
	conf := &mysql.DBConfig{
		Name:     name,
		User:     config.GetString(name + "::user"),
		Password: config.GetString(name + "::password"),
		Host:     config.GetString(name + "::host"),
		Port:     config.DefaultInt(name+"::port", 3306),
		Debug:    debug,
	}
	mysql.InitDefaultDB(conf)
	logy.Info(fmt.Sprintf("[DB]-[%v] initialized successfully => %v", conf.Name, conf))
}
