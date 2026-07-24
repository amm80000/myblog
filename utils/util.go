package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {
	// ✅ 使用 DefaultString 避免 error 处理，并提供默认值（可选）
	driverName := beego.AppConfig.DefaultString("driverName", "mysql")
	user := beego.AppConfig.DefaultString("mysqluser", "root")
	pwd := beego.AppConfig.DefaultString("mysqlpwd", "")
	host := beego.AppConfig.DefaultString("host", "127.0.0.1")
	port := beego.AppConfig.DefaultString("port", "3306")
	dbname := beego.AppConfig.DefaultString("dbname", "test")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
	}
}

// 创建用户类
func CreateTableWithUser() {

	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`

	ModifyDB(sql)
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {

	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	count, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

// 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
