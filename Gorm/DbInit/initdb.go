package DbInit

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var MyDb *gorm.DB
var sqlDb *sql.DB

func InitMysqlDB() {
	DSN := "root:123456@tcp(127.0.0.1:3306)/ForGorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MyDb, err = gorm.Open(mysql.New(mysql.Config{
		//DriverName: "my_mysql_driver",	//自定义驱动
		DSN: DSN, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		CreateBatchSize: 100,	//批量创建时，一批最多的数量
	})

	if err != nil {
		fmt.Println("Mysql Connect Error: ", err)
	}else {
		fmt.Println("Mysql Connect Success!!!")
	}

	//管理连接池
	sqlDb, err = MyDb.DB()
	if err != nil {
		fmt.Println("Mysql pool init Error: ", err)
	}else {
		fmt.Println("Mysql pool init Success!!!")
	}

	sqlDb.SetMaxIdleConns(10)	//设置空闲连接池中连接的最大数量
	sqlDb.SetMaxOpenConns(100)	//设置打开数据库连接的最大数量
	sqlDb.SetConnMaxLifetime(time.Hour)	//设置了连接可复用的最大时间
}
