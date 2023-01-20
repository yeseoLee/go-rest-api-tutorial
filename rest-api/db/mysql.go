package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	dbpool                *sql.DB
	current_connectstring string
}

func MakeMySQLConnectionString(ip string, port int, id string, pw string, dbName string) string {
	//id + ":" + pw + "@tcp(" + ip + ":" + strconv.Itoa(port) + ")/" + dbName
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", id, pw, ip, port, dbName)
}

func (m *MySQL) Connect(connectionstring string) bool {
	dbpool, err := sql.Open("mysql", connectionstring)
	if err != nil {
		fmt.Println("mysql connect error", err)
		return false
	}
	m.dbpool = dbpool
	m.current_connectstring = connectionstring
	//m.dbpool.SetConnMaxLifetime(0)
	//m.dbpool.SetMaxIdleConns(5)
	//m.dbpool.SetMaxOpenConns(5)

	return true
}

func (m *MySQL) Reconnect() bool {
	dbpool, err := sql.Open("mysql", m.current_connectstring)
	if err != nil {
		fmt.Println("mysql reconnect error", err)
		return false
	}
	m.dbpool = dbpool
	//m.dbpool.SetConnMaxLifetime(0)
	//m.dbpool.SetMaxIdleConns(5)
	//m.dbpool.SetMaxOpenConns(5)

	return true
}

func (m *MySQL) PingDB() error {
	err := m.dbpool.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQL) DBPool() *sql.DB {
	return m.dbpool
}
