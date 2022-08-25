package bootstrap

import (
	"log"
	"os"
	"time"

	"giligili/app/model/user"
	"giligili/app/model/video"
	"giligili/pkg/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetUpDB() {
	
	var dbConfig gorm.Dialector	
	dsn := os.Getenv("MYSQL_DSN")
	
	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})


	// 初始化 GORM 日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)


	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, newLogger)


	
	// 设置连接池（最大连接数、最大空闲连接数、每个连接的过期时间）
	database.SQLDB.SetMaxOpenConns(20)  // 打开
	database.SQLDB.SetMaxIdleConns(10)  // 空闲
	database.SQLDB.SetConnMaxLifetime(time.Duration(300) * time.Second)

	// 自动迁移
	database.DB.AutoMigrate((&user.User{}))
	database.DB.AutoMigrate((&video.Video{}))

}


