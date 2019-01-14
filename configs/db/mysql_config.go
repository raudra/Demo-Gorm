package db

import(
	"demo_gorm/configs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type MysqlConf struct{
	Conn *gorm.DB
	DbConf
}

func(self *MysqlConf)initDbConf(){
	fmt.Println("Initialization the database connection for mysql")
	var c config.Conf
	data := c.FetchConfig("/Users/unbxd/go-workspace/src/demo_gorm/configs/db/mysql.yml")
	self.setConfDetail(data)
	self.buildConn()	
}

func(self *MysqlConf)setConfDetail(data map[string]interface{}) {
	conf := DbConf{
				data["database"].(string),
				data["host"].(string),
				data["username"].(string),
				data["password"].(string),
				data["port"].(int),
			}

	self.DbConf=conf
}

func(self *MysqlConf)buildConn() {
	self.Conn, err = gorm.Open("mysql", self.connString())	
	self.Conn.LogMode(true)
	config.RaiseException(err)
}

func(self *MysqlConf)connString() string{
	str:= fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", self.DbConf.Username, self.DbConf.Password, self.DbConf.Host, self.DbConf.Port, self.DbConf.Database)
	return str
}

