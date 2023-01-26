package datatype

func GenerateDataType(sqlType string) (goType string) {

	switch sqlType {
	case "int":
		goType = "int"
	case "int4":
		goType = "int"
	case "int8":
		goType = "int"
	case "varchar":
		goType = "string"
	case "datetime":
		goType = "time.Time"
	case "timestamp":
		goType = "time.Time"
	case "text":
		goType = "string"
	case "tinyint":
		goType = "int"
	case "bigint":
		goType = "int"
	case "decimal":
		goType = "float64"
	case "double":
		goType = "float64"
	case "float":
		goType = "float64"
	case "date":
		goType = "time.Time"
	case "char":
		goType = "string"
	case "mediumtext":
		goType = "string"
	case "longtext":
		goType = "string"
	case "tinytext":
		goType = "string"
	case "json":
		goType = "string"
	case "jsonb":
		goType = "[]bytes"
	case "enum":
		goType = "string"
	case "blob":
		goType = "string"
	case "mediumblob":
		goType = "string"
	case "longblob":
		goType = "string"
	case "tinyblob":
		goType = "string"
	case "binary":
		goType = "string"
	case "varbinary":
		goType = "string"
	case "bit":
		goType = "string"
	case "bool":
		goType = "bool"
	case "boolean":
		goType = "bool"
	case "uuid":
		goType = "string"
	default:
		goType = "interface{}"
	}

	return goType
}
