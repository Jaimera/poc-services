package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jaimera/poc-services/domain/contract"
)

// Conn is the MySQL connection manager
type Conn struct {
	db *sql.DB
}

// Connect returns a new connection, representing a DataManager.
func Connect(
	user string,
	pass string,
	name string,
	host string,
	port int,
) (contract.DataManager, error) {

	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%d", host, port)
	config.User = user
	config.DBName = name
	config.Passwd = pass

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	conn := new(Conn)
	conn.db = db

	return conn, nil
}

// Begin starts a transaction
func (c *Conn) Begin() (*sql.Tx, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// Close closes the db connection
func (c *Conn) Close() (err error) {
	return c.db.Close()
}

func (c *Conn) Port() contract.PortRepository {
	return newPortRepository(c.db)
}
