package operations

import (
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

func GetSchemaFromDocument(document bson.M) bson.M {
	schema := bson.M{}
	for key, value := range document {
		schema[key] = getTypeFromField(value)
	}
	return schema
}

func getTypeFromField(field interface{}) string {
	v := reflect.ValueOf(field)
	switch v.Type().String() {

	case "bson.M":
		// Recursively get the schema
		schema := GetSchemaFromDocument(field.(bson.M))
		return ConvertBSONToJSON(schema)

	default:
		return v.Type().String()
	}
}

func FormatSchemas(schemas map[string]string) string {
	formatted := ""
	for name, schema := range schemas {
		formatted += name + ": " + schema + "\n"
	}
	return formatted
}
