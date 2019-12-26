package database

import(
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(

	//Dialect
	Dialect = "mysql"

	//DBUser ユーザー名
	DBUser = os.Getenv("DBuser")

	//DBPass パスワード
	DBPass = os.Getenv("DBpass")

	//DBProtocol プロトコル
	DBProtocol = "tcp(127.0.0.1)"

	//DBName DB名
	DBName = "slope"

	//DBchar 文字コード
	DBchar = "charset=utf8mb4"
)

//ConnectDB DBにアクセス
func ConnectDB() (*gorm.DB, error) {
	connectTemplate := "%s:%s@%s/%s?%s&parseTime=true"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName, DBchar,)

	db, err := gorm.Open(Dialect, connect)
	return db, err
}