package sqlow

import (
	"fmt"
	"time"
)

//ColumnData is Stores SQL Column data.
type ColumnData struct {
	Name            string
	DataType        DataType
	PrimaryKey      bool
	NotNull         bool
	UniqueIndex     bool
	Unsigned        bool
	ZeroFill        bool
	AutoIncremental bool
	Default         interface{}
	Property        interface{}
}

// SetDefault can be set by passing Boolean whether to set Column to SetDefault
func (col *ColumnData) SetDefault(value interface{}) *ColumnData {
	col.Default = value
	return col
}

// Build is Converts column data to SQL syntax.
func (col ColumnData) Build() (string, error) {
	if database == nil {
		return "", &Error{Msg: "Currently the Database has not been created yet. Create it with New (*sql.DB, Name)."}
	}
	result := fmt.Sprintf("`%s` ", col.Name)
	switch col.DataType.TypeName {
	case "ENUM":
		if defa := col.Property; defa != nil {
			if res, ok := defa.([]string); ok {
				result += fmt.Sprintf(" %v", fmt.Sprintf("ENUM(%v)", toSQLList(res)))
			}
		} else {
			return "", &Error{Msg: "The SQL syntax could not be created successfully because the property is not set in the ENUM of the Column."}
		}
	case "SET":
		if prop := col.Property; prop != nil {
			if res, ok := prop.([]string); ok {
				result += fmt.Sprintf(" %v", fmt.Sprintf("SET(%v)", toSQLList(res)))
			}
		} else {
			return "", &Error{Msg: "The SQL syntax could not be created successfully because the property is not set in the SET of the Column."}
		}
	default:
		if prop := col.Property; prop != nil {
			if res, ok := prop.(int); ok {
				result += fmt.Sprintf(" %v(%v)", col.DataType.TypeName, res)
			} else if res, ok := prop.(string); ok {
				result += fmt.Sprintf(" %v(`%v`)", col.DataType.TypeName, res)
			}
		} else {
			result += col.DataType.TypeName
		}
	}

	dataType := col.DataType.TypeName
	if col.ZeroFill {
		if col.DataType.ZEROFILL {
			result += " ZEROFILL"
		}
	}
	if col.Unsigned {
		if col.DataType.UNSIGNED {
			result += " UNSIGNED"
		}
	}
	if col.NotNull {
		result += " NOT"
	}
	result += " NULL"
	if col.AutoIncremental {
		if col.DataType.AutoIncrement {
			result += " AUTO_INCREMENT"
		}
	}
	if !col.DataType.AutoIncrement {
		if defa := col.Default; defa != nil {
			switch dataType {
			case "TINYINT", "SMALLINT", "MEDIUMINT", "INT", "INTEGER", "BIGINT", "FLOAT", "DOUBLE":
				if col.Unsigned {
					if res, ok := defa.(uint64); ok {
						result += fmt.Sprintf(" %v", res)
					} else {
						return "", &Error{Msg: "The default value could not be successfully converted to a number."}
					}
				} else {
					if res, ok := defa.(int64); ok {
						result += fmt.Sprintf(" %v", res)
					} else {
						return "", &Error{Msg: "The default value could not be successfully converted to a number."}
					}
				}
			case "BOOL", "BOOLEAN":
				if res, ok := defa.(bool); ok {
					result += fmt.Sprintf(" %v", res)
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to Boolean."}
				}
			case "DATE":
				if res, ok := defa.(time.Time); ok {
					result += fmt.Sprintf(" `%v`", toSQLDate(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to time.Time."}
				}
			case "DATETIME":
				if res, ok := defa.(time.Time); ok {
					result += fmt.Sprintf(" `%v`", toSQLDateTime(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to time.Time."}
				}
			case "TIMESTAMP":
				if res, ok := defa.(time.Time); ok {
					result += fmt.Sprintf(" `%v`", toSQLDateTime(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to time.Time."}
				}
			case "TIME":
				if res, ok := defa.(time.Time); ok {
					result += fmt.Sprintf(" `%v`", toSQLTime(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to time.Time."}
				}
			case "TEXT", "VARCHAR", "MIDIUMTEXT", "LONGTEXT", "ENUM":
				if res, ok := defa.(string); ok {
					result += fmt.Sprintf(" `%v`", res)
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to String."}
				}
			case "SET":
				if res, ok := defa.([]string); ok {
					result += fmt.Sprintf(" %v", toSQLList(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to List."}
				}
			}
		}
	}
	return result, nil
}

// SetPrimaryKey can be set by passing Boolean whether to set Column to PrimaryKey
func (col *ColumnData) SetPrimaryKey(check bool) *ColumnData {
	col.PrimaryKey = check
	return col
}

// SetNotNull can be set by passing Boolean whether to set Column to SetNotNull
func (col *ColumnData) SetNotNull(check bool) *ColumnData {
	col.NotNull = check
	return col
}

// SetUniqueIndex can be set by passing Boolean whether to set Column to SetUniqueIndex
func (col *ColumnData) SetUniqueIndex(check bool) *ColumnData {
	col.UniqueIndex = check
	return col
}

// SetUnsigned can be set by passing Boolean whether to set Column to SetUnsigned
func (col *ColumnData) SetUnsigned(check bool) *ColumnData {
	if col.DataType.UNSIGNED {
		col.Unsigned = check
		return col
	}
	return col
}

// SetZeroFill can be set by passing Boolean whether to set Column to SetZeroFill
func (col *ColumnData) SetZeroFill(check bool) *ColumnData {
	if col.DataType.ZEROFILL {
		col.ZeroFill = check
		return col
	}
	return col
}

// SetAutoIncrement can be set by passing Boolean whether to set Column to SetAutoIncrement
func (col *ColumnData) SetAutoIncrement(check bool) *ColumnData {
	if col.DataType.AutoIncrement {
		col.AutoIncremental = check
		return col
	}
	return col
}
