package mysql

import (
	"fmt"
	"webWorks02/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		settings.Conf.MySqlConfig.Username,
		settings.Conf.MySqlConfig.Password,
		settings.Conf.MySqlConfig.Host,
		settings.Conf.MySqlConfig.Port,
		settings.Conf.MySqlConfig.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(settings.Conf.MySqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(settings.Conf.MySqlConfig.MaxIdleConns)
	return
}

func Close() {
	db.Close()
}
