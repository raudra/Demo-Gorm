package db

import(
	"demo_gorm/configs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var Conn *gorm.DB

var err error

type dbConf struct{
	database string
	host string
	username string
	password string
	port int
}

var DbConf dbConf

func init(){
	initDbConf()
	fmt.Println("Initialization the database connection")
	Conn, err = gorm.Open("mysql", buildconnectionString())	
	Conn.LogMode(true)
	config.RaiseException(err)
}

func initDbConf(){
	var c config.Conf
	data := c.FetchConfig("/Users/unbxd/go-workspace/src/demo_gorm/configs/db/config.yml")
	DbConf = dbConf{
			data["database"].(string),
			data["host"].(string),
			data["username"].(string),
			data["password"].(string),
			data["port"].(int),
		}
}

func buildconnectionString() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", DbConf.username, DbConf.password, DbConf.host, DbConf.port, DbConf.database)
}

