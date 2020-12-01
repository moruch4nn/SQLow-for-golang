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

func toSQLList(val []interface{}) string {
	tmp := []string{}
	for _, value := range val {
		switch value.(type) {
		case int, int16, int32, int64, int8, uint, uint8, uint16, uint32, uint64, float32, float64:
			tmp = append(tmp, fmt.Sprintf("%v", value))
		default:
			tmp = append(tmp, fmt.Sprintf("`%v`", value))
		}
	}
	return strings.Join(tmp, ",")
}

func toSQLListS(val []interface{}) string {
	tmp := []string{}
	for _, value := range val {
		switch value.(type) {
		case int, int16, int32, int64, int8, uint, uint8, uint16, uint32, uint64, float32, float64:
			tmp = append(tmp, fmt.Sprintf("%d", value))
		default:
			tmp = append(tmp, fmt.Sprintf("'%s'", value))
		}
	}
	return strings.Join(tmp, ",")
}

//ToSQLType is converts the type you entered to SQL.
func ToSQLType(val interface{}) string {
	switch val.(type) {
	case int, int16, int32, int64, int8, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprintf("%d", val)
	case time.Time:
		if val, ok := val.(time.Time); ok {
			return fmt.Sprintf("`%s`", val.Format("2006/1/2 15:04:05"))
		}
	default:
		return fmt.Sprintf("`%d`", val)
	}
	return ""
}

//ToSQLTypeS is converts the type you entered to SQL.
func ToSQLTypeS(val interface{}) string {
	switch val.(type) {
	case int, int16, int32, int64, int8, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprintf("%d", val)
	case time.Time:
		if val, ok := val.(time.Time); ok {
			return fmt.Sprintf("'%s'", val.Format("2006/1/2 15:04:05"))
		}
	default:
		return fmt.Sprintf("'%s'", val)
	}
	return ""
}
