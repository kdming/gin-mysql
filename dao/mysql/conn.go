package mysql

import (
	"app/common/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	var err error
	dsn := config.GetConfig().MySqlUrl
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql连接成功，准备更新表格......")
	migrate()
}
