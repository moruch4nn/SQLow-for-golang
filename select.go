package sqlow

import (
	"fmt"
	"strings"
	"time"
)

// SelectData is data that generates a SELECT statement.
type SelectData struct {
	TableName string
	Values    []string
	Cond      map[string]interface{}
	Greater   map[string]interface{}
	Less      map[string]interface{}
	Like      map[string]string
	Join      map[string]map[string]string
	Limit     int
	Distinct  bool
}

// Select returns SelectData when passed a TableName.
func Select(tableName string, values []string) *SelectData {
	data := SelectData{
		TableName: tableName,
		Values:    values,
		Cond:      map[string]interface{}{},
		Greater:   map[string]interface{}{},
		Less:      map[string]interface{}{},
		Like:      map[string]string{},
	}
	return &data
}

// Build converts SelectData to SELECT syntax.
func (sel *SelectData) Build() (string, error) {
	result := "SELECT "
	for _, i := range sel.Values {
		result += fmt.Sprintf(", %s", i)
	}
	result = strings.Replace(result, ", ", "", 1)
	result += fmt.Sprintf(" FROM %s", sel.TableName)
	if (len(sel.Less) + len(sel.Greater) + len(sel.Cond) + len(sel.Like)) > 0 {
		result += " WHERE"
		for i, v := range sel.Cond {
			result += fmt.Sprintf(" AND %s = %s", i, toSQLTypeS(v))
		}
		result = strings.Replace(result, " AND", "", 1)
		for i, v := range sel.Greater {
			result += fmt.Sprintf(" AND %s > %s", i, toSQLTypeS(v))
		}
		result = strings.Replace(result, " AND", "", 1)
		for i, v := range sel.Less {
			result += fmt.Sprintf(" AND %s < %s", i, toSQLTypeS(v))
		}
		result = strings.Replace(result, " AND", "", 1)
		for i, v := range sel.Like {
			result += fmt.Sprintf(" AND %s LIKE %s", i, v)
		}
		result = strings.Replace(result, " AND", "", 1)
	}
	return result, nil
}

// SetLike sets the LIKE condition to SelectData.
func (sel *SelectData) SetLike(key string, value string) *SelectData {
	sel.Like[key] = value
	return sel
}

// SetCond sets the condition to SelectData.
func (sel *SelectData) SetCond(key string, value interface{}) *SelectData {
	sel.Cond[key] = toSQLTypeS(value)
	return sel
}

// SetGreater sets a greater condition on SelectData.
func (sel *SelectData) SetGreater(key string, value interface{}) *SelectData {
	switch value.(type) {
	case int, int16, int32, int64, int8, uint, uint8, uint16, uint32, uint64, float32, float64:
		sel.Greater[key] = fmt.Sprintf("%s", value)
	case time.Time:
		if val, ok := value.(time.Time); ok {
			sel.Greater[key] = fmt.Sprintf("'%s'", val.Format("2006/1/2 15:04:05"))
		}
	}
	return sel
}

// SetLesser sets a lesser condition to SelectData.
func (sel *SelectData) SetLesser(key string, value interface{}) *SelectData {
	switch value.(type) {
	case int, int16, int32, int64, int8, uint, uint8, uint16, uint32, uint64, float32, float64:
		sel.Less[key] = fmt.Sprintf("%s", value)
	case time.Time:
		if val, ok := value.(time.Time); ok {
			sel.Less[key] = fmt.Sprintf("'%s'", val.Format("2006/1/2 15:04:05"))
		}
	}
	return sel
}
