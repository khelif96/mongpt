package operations

import (
	"os"
	"testing"
)

const testDir = "./bin/test"

func TestMain(m *testing.M) {
	// Create the bin/cache directory if it doesn't exist
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		os.MkdirAll(testDir, 0755)
	} else {
		os.RemoveAll(testDir)
		os.MkdirAll(testDir, 0755)
	}
	code := m.Run()
	os.Exit(code)
}

func TestWriteJSONSchemaToFile(t *testing.T) {
	err := WriteJSONSchemaToFile(testDir+"/test.json", `{"test": "test"}`)
	if err != nil {
		t.Error(err)
	}

	// Check if the file exists
	_, err = os.Stat(testDir + "/test.json")

	if os.IsNotExist(err) {
		t.Error("File doesn't exist")
	}

}
