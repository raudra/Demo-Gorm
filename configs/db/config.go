package db

import(
	"log"
	"github.com/jinzhu/gorm"
	// "reflect"
)

type Dbinterface interface{
	initDbConf()
}

var err error

type DbConf struct{
	Database string
	Host string
	Username string
	Password string
	Port int
}

var MysqlConn *gorm.DB

func init(){
	log.Println("Initialization the database connection")
	intiMySql()
}

func intiMySql(){
	var dbInfc Dbinterface
	conf := MysqlConf{}
	dbInfc =&conf
	dbInfc.initDbConf()	
	MysqlConn =conf.Conn
}