package operations

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
	"gotest.tools/v3/assert"
)

func TestGetSchemaFromDocument(t *testing.T) {
	document := bson.M{
		"test":  "test",
		"test2": 1,
		"test3": 1.0,
	}
	schema := GetSchemaFromDocument(document)
	assert.Equal(t, schema["test"], "string")
	assert.Equal(t, schema["test2"], "int")
	assert.Equal(t, schema["test3"], "float64")

	document = bson.M{
		"test": bson.M{
			"test": "test",
		},
		"test2": bson.M{
			"test": 1,
		},
		"test3": bson.M{
			"test": 1.0,
		},
	}
	schema = GetSchemaFromDocument(document)
	assert.Equal(t, schema["test"].(bson.M)["test"], "string")
	assert.Equal(t, schema["test2"].(bson.M)["test"], "int")
	assert.Equal(t, schema["test3"].(bson.M)["test"], "float64")

}
