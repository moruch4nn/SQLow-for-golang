package sqlow

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

var database *DB

// DB is data struct
type DB struct {
	Database *sql.DB
	DBName   string
}

// Error is Original Error
type Error struct {
	Msg string
}

func contains(list []string, target string) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}

func (err *Error) Error() string {
	return fmt.Sprintf("Sqlow Error: %v", err.Msg)
}

// New returns sqlow.Data if sql.DB already exists.
func New(db *sql.DB, dbname string) *DB {
	database = &DB{
		DBName:   dbname,
		Database: db,
	}
	return database
}

// Ping is Ping to SQL
func (data DB) Ping() error {
	return data.Database.Ping()
}

// Close is close the connection to SQL.
func (data DB) Close() error {
	return data.Database.Close()
}

func toSQLDate(val time.Time) string {
	return fmt.Sprintf("%v-%v-%v", val.Year(), val.Month(), val.Day())
}

func toSQLDateTime(val time.Time) string {
	return fmt.Sprintf("%v-%v-%v %v:%v:%v", val.Year(), val.Month(), val.Day(), val.Hour(), val.Minute(), val.Second())
}

func toSQLTime(val time.Time) string {
	return fmt.Sprintf("%v:%v:%v", val.Hour(), val.Minute(), val.Second())
}

func toSQLList(val []string) string {
	tmp := []string{}
	for _, value := range val {
		tmp = append(tmp, fmt.Sprintf("`%v`", string(value)))
	}
	return strings.Join(tmp, ",")
}
