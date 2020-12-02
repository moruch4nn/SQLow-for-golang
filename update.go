package sqlow

import (
	"fmt"
	"strings"
)

// Update is change the INSERT statement to UPDATE syntax.
func (ins InsertData) Update(key string, value interface{}) string {
	result := fmt.Sprintf("UPDATE %s SET", ins.TableName)
	for i, x := range ins.Values {
		result += fmt.Sprintf(", %s = %s", i, toSQLTypeS(x))
	}
	result = strings.Replace(result, ",", "", 1)
	result += fmt.Sprintf(" WHERE %s = %s", key, toSQLTypeS(value))
	return result
}
