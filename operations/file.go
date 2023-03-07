package operations

import (
	"encoding/json"
	"os"

	"gopkg.in/mgo.v2/bson"
)

func WriteSchemaToFile(filename string, data bson.M) error {
	// Create file and write to it
	file, _ := json.MarshalIndent(data, "", " ")
	return os.WriteFile(filename, file, 0644)

}
