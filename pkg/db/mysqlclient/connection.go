package mysqlclient

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlClient(
	host,
	port,
	user,
	pass,
	dbName string,
	maxConn,
	maxIdleConn int,
) (*gorm.DB, *sql.DB, error) {
	connection := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Etc/UTC")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil && viper.GetBool("DEBUG") {
		return nil, nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetMaxIdleConns(maxIdleConn)

	return db, sqlDB, nil
}
