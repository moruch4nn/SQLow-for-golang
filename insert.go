package sqlow

import "fmt"

//InsertData is Stores the data used to generate the INSERT syntax.
type InsertData struct {
	TableName string
	Values    map[string]interface{}
}

// Insert is If you pass the required data, it will return InsertData.
func Insert(tableName string, values map[string]interface{}) *InsertData {
	data := InsertData{TableName: tableName, Values: values}
	return &data
}

// Build is Converts InsertData to INSERT syntax.
func (ins InsertData) Build() (string, error) {
	keys := []interface{}{}
	values := []interface{}{}
	result := fmt.Sprintf("INSERT INTO %v (", ins.TableName)
	for k := range ins.Values {
		keys = append(keys, k)
		values = append(values, ins.Values[k])
	}
	result += toSQLList(keys) + ") VALUES ("
	result += toSQLList(values) + ")"
	return result, nil
}
