package sqlow

import "fmt"

func Delete(tableName string, key string, value interface{}) (Status, error) {
	_, err := database.Database.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = %s", tableName, key, toSQLTypeS(value)))
	if err != nil {
		return ERROR, err
	}
	return DELETE, nil
}
