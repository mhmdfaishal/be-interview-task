package boot

import (
	_mysql "be-interview-task/pkg/db/mysqlclient"
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	MainDBConn *gorm.DB
	MainSqlDB  *sql.DB
)

func init() {

	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("DEBUG") {
		fmt.Println("Service RUN on DEBUG mode")
	}

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPass := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbMaxConn := viper.GetInt("DB_MAX_CONN")
	dbMaxIdleConn := viper.GetInt("DB_MAX_IDLE_CONN")
	MainDBConn, MainSqlDB, err = _mysql.NewMysqlClient(dbHost, dbPort, dbUser, dbPass, dbName, dbMaxConn, dbMaxIdleConn)
	if err != nil {
		log.Fatalf("MainDBConn.Init: %s", err)
	}
}

func FlushResources() {
	fmt.Println("stopping main db connection")
	err := MainSqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main db connection stopped")
}
