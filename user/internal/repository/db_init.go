package repository

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	// 部署到docker 时区错误 删除loc
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username, password, host, port, database, charset, url.QueryEscape(loc))
	fmt.Println(dsn)
	err := Database(dsn)
	if err != nil {
		panic(err)
	}

}

func Database(dsn string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,  // 禁用datetime的精度，mysql5.6之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引的时候采用删除并新建的方式，mysql5.7之前的数据库不支持
		DontSupportRenameColumn:   true,  // 用change重命名列，mysql8之前的数据库不支持
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 设置最大空闲连接池
	sqlDB.SetMaxOpenConns(100)                 // 最大打开数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 最大连接时间
	DB = db
	migration() // 迁移
	return nil
}
