package utils

import (
	"strconv"

	"github.com/larspensjo/config"
)

var MONGODB = "mongodb://root:root@192.168.2.173:27017/"
var PORT = 7999
var Corpid = ""
var Agentid = ""
var Secret = ""

var CurrentMode = "DEV"

func RunMode(CurrentMode string) {
	conf, _ := config.ReadDefault("conf/app.conf")
	MONGODB, _ = conf.String(CurrentMode, "mongodb")
	_port, _ := conf.String(CurrentMode, "port")
	PORT, _ = strconv.Atoi(_port)

	Corpid, _ = conf.String(CurrentMode, "corpid")
	Agentid, _ = conf.String(CurrentMode, "agentid")
	Secret, _ = conf.String(CurrentMode, "secret")

}
