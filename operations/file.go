package operations

import (
	"encoding/json"
	"os"
)

func WriteJSONSchemaToFile(filename string, data string) error {
	// Convert data to bytes
	bytes := []byte(data)
	return os.WriteFile(filename, bytes, 0644)
}

func ReadJSONSchemaFromFile(filename string) string {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// Read the file
	data := ""
	decoder := json.NewDecoder(file)
	decoder.Decode(&data)

	return data
}

func ReadJSONSchemasFromDir(dir string) []string {
	// Get all files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// Read all the files and convert them to bson.M
	schemas := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		schemas = append(schemas, ReadJSONSchemaFromFile(dir+"/"+file.Name()))
	}

	return schemas
}

func Init() {
	// Create the bin/cache directory if it doesn't exist
	if _, err := os.Stat("./bin/cache"); os.IsNotExist(err) {
		os.MkdirAll("./bin/cache", 0755)
	}
}
