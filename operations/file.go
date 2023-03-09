package operations

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSONSchemaToFile(filename string, data string) error {
	// Convert data to bytes
	bytes := []byte(data)
	return os.WriteFile(filename, bytes, 0644)
}

func ReadJSONSchemaFromFile(filename string) string {
	fmt.Println("Opening file: ", filename)
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// Read the file
	data := ""
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var v interface{}
		err := decoder.Decode(&v)
		if err != nil {
			panic(err)
		}
		data += fmt.Sprintf("%v", v)
	}

	return data
}

func ReadJSONSchemasFromDir(dir string) map[string]string {

	schemas := map[string]string{}
	// Get all files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// Read all the files and convert them to bson.M
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		// Remove the .json extension
		fmt.Println(dir + file.Name())
		schemas[file.Name()[:len(file.Name())-5]] = ReadJSONSchemaFromFile(dir + file.Name())
	}

	return schemas
}

func Init() {
	// Create the bin/cache directory if it doesn't exist
	if _, err := os.Stat("./bin/cache"); !os.IsNotExist(err) {
		os.RemoveAll("./bin/cache")
	}
	os.MkdirAll("./bin/cache", 0755)

}
