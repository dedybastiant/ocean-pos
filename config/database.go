package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func NewDB(config *viper.Viper) *sql.DB {
	username := config.GetString("DATABASE_USERNAME")
	password := config.GetString("DATABASE_PASSWORD")
	db_url := config.GetString("DATABASE_URL")
	db_port := config.GetString("DATABASE_PORT")
	db_name := config.GetString("DATABASE_NAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, db_url, db_port, db_name)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
