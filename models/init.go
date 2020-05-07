package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/beeapi/utils"
)

func init()  {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",utils.G_mysql_uname+":"+utils.G_mysql_passwd+"@tcp("+utils.G_mysql_addr+":"+utils.G_mysql_port+")/"+utils.G_mysql_dbname+"?charset=utf8mb4",utils.G_mysql_maxidle,utils.G_mysql_maxconn)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default",false,true)
}

func GetTableName(dbname string) string {
	prefix := utils.G_mysql_prefix
	return prefix+dbname
}
//表名
func GetUserTableName() string  {
	return GetTableName("backend_user")
}
