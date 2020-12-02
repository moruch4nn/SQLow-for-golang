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
	result += toSQLListS(values) + ")"
	return result, nil
}

// Send synchronizes InsertData to the database.
func (ins InsertData) Send() (Status, error) {
	res, err := ins.Build()
	if err != nil {
		return ERROR, err
	}
	_, err = database.Database.Exec(res)
	if err != nil {
		return ERROR, err
	}
	return ADD, nil
}

// SendPK will pass if the specified key exists and INSERT if it does not exist.
func (ins InsertData) SendPK(key string) (Status, error) {
	if value := ins.Values[key]; value != nil {
		if res, err := database.Database.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s = %s LIMIT 1;", ins.TableName, key, toSQLTypeS(value))); err == nil {
			if res.Next() {
				return PASS, nil
			}
		} else {
			return ERROR, err
		}
	} else {
		return "", &Error{Msg: "Key is nil"}
	}
	res, err := ins.Build()
	if err != nil {
		return ERROR, err
	}
	_, err = database.Database.Exec(res)
	if err != nil {
		return ERROR, err
	}
	return ADD, nil
}

// ForceSend will UPDATE if the specified key exists and INSERT if it does not.
func (ins InsertData) ForceSend(key string) (Status, error) {
	if value := ins.Values[key]; value != nil {
		if res, err := database.Database.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s = %s LIMIT 1;", ins.TableName, key, toSQLTypeS(value))); err == nil {
			if res.Next() {
				_, err := database.Database.Exec(ins.Update(key, ins.Values[key]))
				if err != nil {
					return ERROR, err
				}
				return UPDATE, err
			}
		} else {
			return ERROR, err
		}
	} else {
		return "", &Error{Msg: "PK is nil"}
	}
	res, err := ins.Build()
	if err != nil {
		return ERROR, err
	}
	_, err = database.Database.Exec(res)
	if err != nil {
		return ERROR, err
	}
	return ADD, nil
}
