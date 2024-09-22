// utils/db_util.go
package utils

import (
	"github.com/jackc/pgx"
)

// DBClient defines the interface for database operations
type DBClient interface {
	Exec(query string, args ...interface{}) error
	Query(query string, args ...interface{}) *pgx.Row
	// Add other DB methods as needed
}

type PgxClient struct {
	Conn *pgx.Conn
}

func (c *PgxClient) Exec(query string, args ...interface{}) error {
	_, err := c.Conn.Exec(query, args...)
	return err
}

func (c *PgxClient) Query(query string, args ...interface{}) *pgx.Row {
	row := c.Conn.QueryRow(query, args...)
	return row
}
