package operations

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

func ConvertJSONToBSON(json string) bson.M {
	var data bson.M
	bson.UnmarshalJSON([]byte(json), &data)
	return data
}

func ConvertJSONArrayToBSON(json string) []bson.M {
	var data []bson.M
	bson.UnmarshalJSON([]byte(json), &data)
	return data
}

func ConvertBSONToJSON(bson bson.M) string {
	json, _ := json.MarshalIndent(bson, "", " ")
	return string(json)

}
