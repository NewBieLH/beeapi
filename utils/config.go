package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var (
	G_mysql_addr   string //mysql ip 地址
	G_mysql_port   string //mysql 端口
	G_mysql_dbname string //mysql db name
	G_mysql_uname  string //mysql username
	G_mysql_passwd string //mysql password
	G_mysql_maxidle int   //mysql maxidle
	G_mysql_maxconn int   // mysql maxconn
	G_mysql_prefix string
)

func InitConfig() {
	//从配置文件读取配置信息
	appconf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		beego.Debug(err)
		return
	}
	G_mysql_addr = appconf.String("mysqladdr")
	G_mysql_port = appconf.String("mysqlport")
	G_mysql_dbname = appconf.String("mysqldbname")
	G_mysql_uname = appconf.String("mysqlusername")
	G_mysql_passwd = appconf.String("mysqlpassword")
	G_mysql_maxidle,_ = appconf.Int("mysqlmaxIdle")
	G_mysql_maxconn,_ = appconf.Int("mysqlmaxConn")
	G_mysql_prefix = appconf.String("mysqlprefix")

	return
}

func init() {
	InitConfig()
}