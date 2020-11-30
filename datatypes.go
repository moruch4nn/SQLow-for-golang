package sqlow

// DataType is SQL DataType
type DataType struct {
	TypeName       string
	Type           string
	UNSIGNED       bool
	UndignedType   string
	ZEROFILL       bool
	MaxLength      int
	DefaultPropaty interface{}
	AutoIncrement  bool
	PrimaryKey     bool
	Default        bool
}

var (
	// TINYINT is the one that exists in SQL and can store numbers from -128 to 127. If you add UNSIGNED, you can store up to 255, but you can only use integers.
	TINYINT = DataType{TypeName: "TINYINT", Type: "INT8", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT8", AutoIncrement: true, PrimaryKey: true, Default: true}

	// SMALLINT is the one that exists in SQL and can store numbers from -32768 to 32767. With UNSIGNED, you can store up to 65535, but you can only use integers.
	SMALLINT = DataType{TypeName: "SMALLINT", Type: "INT16", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT16", AutoIncrement: true, PrimaryKey: true, Default: true}

	// MEDIUMINT is the one that exists in SQL and can store numbers from -8388608 to 8388607. With UNSIGNED, you can store up to 16777215, but you can only use integers.
	MEDIUMINT = DataType{TypeName: "MEDIUMINT", Type: "INT32", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT32", AutoIncrement: true, PrimaryKey: true, Default: true}

	// INT is the one that exists in SQL and can store numbers from -2147483648 to 2147483647. With UNSIGNED, you can store up to 4294967295, but you can only use integers.
	INT = DataType{TypeName: "INT", Type: "INT32", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT32", AutoIncrement: true, PrimaryKey: true, Default: true}

	// INTEGER is the one that exists in SQL and can store numbers from -2147483648 to 2147483647. With UNSIGNED, you can store up to 4294967295, but you can only use integers.
	INTEGER = DataType{TypeName: "INTEGER", Type: "INT32", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT32", AutoIncrement: true, PrimaryKey: true, Default: true}

	// BIGINT is the one that exists in SQL and can store numbers from -9223372036854775808 to 9223372036854775807. With UNSIGNED, you can store up to 18446744073709551615, but you can only use integers.
	BIGINT = DataType{TypeName: "BIGINT", Type: "INT64", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT64", AutoIncrement: true, PrimaryKey: true, Default: true}

	// BOOLEAN is a type in SQL that can store True or False.
	BOOLEAN = DataType{TypeName: "BOOLEAN", Type: "BOOL", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true, Default: true}

	// BOOL is a type in SQL that can store True or False.
	BOOL = DataType{TypeName: "BIGINT", Type: "BOOL", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: false, Default: true}

	// BIT is a SQL type that can store bit values.
	BIT = DataType{TypeName: "UNIT", Type: "UNIT", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true, Default: true}

	// FLOAT is a type that exists in SQL and can store accurate decimals up to the 7th decimal place.
	FLOAT = DataType{TypeName: "FLOAT", Type: "FLOAT32", UNSIGNED: true, ZEROFILL: true, AutoIncrement: true, PrimaryKey: true, Default: true}

	// DOUBLE is a type that exists in SQL and can store accurate decimals up to the 7th decimal place.
	DOUBLE = DataType{TypeName: "DOUBLE", Type: "FLOAT64", UNSIGNED: true, ZEROFILL: true, AutoIncrement: true, PrimaryKey: true, Default: true}

	// DATE is a type that exists in SQL. You can save the year, month, and day.
	DATE = DataType{TypeName: "DATE", Type: "DATE", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true, Default: true}

	// DATETIME is a type that exists in SQL. You can save the year, month, day, hour, minute, and second.
	DATETIME = DataType{TypeName: "DATETIME", Type: "DATE", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true, Default: true}

	// TIMESTAMP is a type that exists in SQL. You can save the year, month, day, hour, minute, and second. Also, if no value is explicitly assigned, the date and time will be set automatically when the value is changed.
	TIMESTAMP = DataType{TypeName: "TIMESTAMP", Type: "DATE", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true, Default: true}

	// TIME is a type that exists in SQL. You can save hours, minutes, and seconds.
	TIME = DataType{TypeName: "TIME", Type: "DATE", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true, Default: true}

	// VARCHAR is a SQL type. You can store the specified character string.
	VARCHAR = DataType{TypeName: "VARCHAR", Type: "STRING", UNSIGNED: false, ZEROFILL: false, MaxLength: 65535, DefaultPropaty: 255, AutoIncrement: false, PrimaryKey: false, Default: true}

	// TEXT is a SQL type. You can store the specified character string.
	TEXT = DataType{TypeName: "TEXT", Type: "STRING", UNSIGNED: false, ZEROFILL: false, MaxLength: 14090025, AutoIncrement: false, PrimaryKey: false, Default: false}

	// MEDIUMTEXT is a SQL type. You can store the specified character string.
	MEDIUMTEXT = DataType{TypeName: "MEDIUMTEXT", Type: "STRING", UNSIGNED: false, ZEROFILL: false, MaxLength: 3741318945, AutoIncrement: false, PrimaryKey: false, Default: false}

	// LONGTEXT is a SQL type. You can store the specified character string.
	LONGTEXT = DataType{TypeName: "LONGTEXT", Type: "STRING", UNSIGNED: false, ZEROFILL: false, MaxLength: 4294967295, AutoIncrement: false, PrimaryKey: false, Default: false}

	// ENUM is a type of SQL. Can store one of the specified string lists.
	ENUM = DataType{TypeName: "TEXT", Type: "LIST", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: false, Default: true}

	// SET is a type of SQL. Can store multiple in the specified string list.
	SET = DataType{TypeName: "SET", Type: "LIST", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: false, Default: true}
)
