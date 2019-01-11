package db

import(
	"log"
	"github.com/jinzhu/gorm"
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
	msqlConf := MysqlConf{}
	dbInfc = &msqlConf
	dbInfc.initDbConf()	
	MysqlConn = msqlConf.Conn
}