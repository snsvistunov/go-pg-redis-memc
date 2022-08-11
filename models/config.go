package models

import (
	"fmt"

	"github.com/go-redis/redis"
)

type dbConfig struct {
	dbType  string
	host    string
	port    uint
	user    string
	dbName  string
	psw     string
	sslMode string
}

var dbConf = dbConfig{
	dbType:  "postgres",
	host:    "kandula.db.elephantsql.com",
	port:    5432,
	user:    "vwxkizsp",
	dbName:  "vwxkizsp ",
	psw:     "ke5GRJ4Z5jfGbuPrd35SRFuANp5QwCdd",
	sslMode: "disable",
}

func loadDbConfig(conf dbConfig) (string, string) {
	confString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", conf.host, conf.port, conf.user, conf.dbName, conf.psw, conf.sslMode)
	return conf.dbType, confString
}

var rConf = &redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
}
