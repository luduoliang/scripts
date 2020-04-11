package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"scripts/config"
)

var DB *sqlx.DB

func init(){
	connStr := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("MYSQL_USERNAME"),
		config.Get("MYSQL_PASSWORD"),
		"tcp("+config.Get("MYSQL_HOST")+":"+config.Get("MYSQL_PORT")+")",
		config.Get("MYSQL_DATABASE"),
	)
	DB = sqlx.MustConnect("mysql", connStr)
}