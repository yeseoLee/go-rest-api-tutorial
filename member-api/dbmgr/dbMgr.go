package dbmgr

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type dbMgr struct {
	db *sql.DB
}

var instance = dbMgr{}

func GetDBMgr() *dbMgr {
	return &instance
}

func (m *dbMgr) GetDB() (db *sql.DB, err error) {
	return m.db, nil
}

func (m *dbMgr) Connect(driverName string, dataSourceName string) error {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	m.db = db
	return nil
}

func MySQLConnectionName(ip string, port int, id string, pw string, dbName string) (driverName string, dataSourceName string) {
	driverName = "mysql"
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", id, pw, ip, port, dbName)
	return
}
