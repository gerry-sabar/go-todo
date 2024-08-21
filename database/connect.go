package database

import (
	"database/sql"
	"fmt"
	"go-todo/config"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

type MySQLConfig struct {
	Hostname           string `json:"hostname" mapstructure:"HOSTNAME"`
	Username           string `json:"username" mapstructure:"USERNAME"`
	Password           string `json:"password" mapstructure:"PASSWORD"`
	Protocol           string `json:"protocol" mapstructure:"PROTOCOL"`
	Name               string `json:"name" mapstructure:"NAME"`
	MaxOpenConnections int    `json:"max_open_connections" mapstructure:"DB_MAX_OPEN_CONNECTION"`
	MaxIdleConnections int    `json:"max_idle_connections" mapstructure:"DB_MAX_IDLE_SIZE"`
	TimeoutInSeconds   int    `json:"timeout_in_seconds" mapstructure:"DB_TIMEOUT_IN_SECOND"`
}

func ConnectDB() (*sql.DB, error) {
	var dsn string
	if config.Cfg.APPEnv == "sandbox" {
		dsn = config.Cfg.DBUser + ":" + config.Cfg.DBPassword + "@unix(/cloudsql/" + config.Cfg.DBHost + ")/" + config.Cfg.DBName + "?parseTime=true"
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Cfg.DBUser, config.Cfg.DBPassword, config.Cfg.DBHost, config.Cfg.DBPort, config.Cfg.DBName)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db, nil
}
