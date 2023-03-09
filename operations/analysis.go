package operations

import "gopkg.in/mgo.v2/bson"

func GetSchemaFromDocument(document bson.M) bson.M {
	schema := bson.M{}
	for key, value := range document {
		schema[key] = getSchemaFromValue(value)
	}
	return schema
}

func getSchemaFromValue(value interface{}) interface{} {
	switch value.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case float64:
		return "float64"
	case bson.M:
		return GetSchemaFromDocument(value.(bson.M))
	case []interface{}:
		return getSchemaFromList(value.([]interface{}))
	default:
		return "unknown"
	}
}

func getSchemaFromList(list []interface{}) interface{} {
	if len(list) == 0 {
		return "unknown"
	}
	schema := getSchemaFromValue(list[0])
	for _, value := range list {
		if schema != getSchemaFromValue(value) {
			return "unknown"
		}
	}
	return schema
}

func FormatSchemas(schemas map[string]string) string {
	formatted := ""
	for name, schema := range schemas {
		formatted += name + ": " + schema + "\n"
	}
	return formatted
}
