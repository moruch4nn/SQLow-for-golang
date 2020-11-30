package sqlow

import "fmt"

//TableData is Stores SQL Table data.
type TableData struct {
	Name    string
	Columns []*ColumnData
	DBName  string
}

// Table is creates a Table and returns TableData. The arguments are Name string, columns []ColumnData.
func Table(name string, columns []*ColumnData) *TableData {
	table := TableData{Name: name, Columns: columns}
	return &table
}

// Build is Converts table data to SQL syntax.
func (tab TableData) Build() (string, error) {
	if database == nil {
		return "", &Error{Msg: "Currently the Database has not been created yet. Create it with New (*sql.DB, Name)."}
	}
	autoIncrement := false
	columns := tab.Columns
	name := tab.Name
	primarys := []interface{}{}
	uniqueIndex := []string{}
	result := fmt.Sprintf("CREATE TABLE `%v`.`%v` (", database.DBName, name)
	columnCounts := len(columns) - 1
	for n, i := range columns {
		if i.AutoIncremental {
			if autoIncrement {
				return "", &Error{Msg: "Up to one AutoIncrement can be set for each table."}
			}
			autoIncrement = true
		}
		if i.PrimaryKey {
			primarys = append(primarys, i.Name)
		}
		if i.UniqueIndex {
			uniqueIndex = append(uniqueIndex, i.Name)
		}
		r, err := i.Build()
		if err != nil {
			return "", err
		}
		result += r
		if n != columnCounts {
			result += ","
		}
	}
	if len(primarys) >= 1 {
		result += ","
		result += fmt.Sprintf("PRIMARY KEY (%v)", toSQLList(primarys))
	}
	if len(uniqueIndex) >= 1 {
		for _, i := range uniqueIndex {
			result += ","
			result += fmt.Sprintf("UNIQUE INDEX `%v_UNIQUE` (`%v` ASC) VISIBLE", i, i)
		}
	}
	result += ");"
	return result, nil
}

// Send is if a table with a matching name exists, it will pass and if it does not exist, a table will be added.
func (tab *TableData) Send() (Status, error) {
	if database == nil {
		return "", &Error{Msg: "Currently the Database has not been created yet. Create it with New (*sql.DB, Name)."}
	}
	tableName := tab.Name
	result, err := database.Database.Query(fmt.Sprintf("SHOW TABLES LIKE \"%v\";", tableName))
	if err != nil {
		return ERROR, err
	}
	defer result.Close()
	if result.Next() {
		return PASS, nil
	}
	res, err := tab.Build()
	if err != nil {
		return ERROR, err
	}
	_, err = database.Database.Exec(res)
	if err != nil {
		return ERROR, err
	}
	return ADD, nil
}
