package db

import "database/sql"

type DBManager interface {
	Connect(string) bool
	Reconnect() bool
	PingDB() error
	DBPool() *sql.DB
}

var instance DBManager

func GetInstance() DBManager {
	if instance == nil {
		instance = *new(DBManager)
	}
	return instance
}

func SetDatabase(ds DBManager) {
	instance = ds
}
