package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/evergreen-ci/utility"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var client *mongo.Client
var database *mongo.Database
var databases []string
var collections []string
var ctx = context.TODO()

func Init() {
	fmt.Println("Connecting to MongoDB...")
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING"))
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

}
func GetDatabases() ([]string, error) {
	var err error
	databases, err = client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrapf(err, "Error getting databases")
	}
	return databases, nil
}

func ChooseDatabase(name string) error {
	if len(databases) == 0 {
		return fmt.Errorf("No databases found! Did you call GetDatabases() first?")
	}
	if !utility.StringSliceContains(databases, name) {
		fmt.Println(databases)
		fmt.Println(name)
		return fmt.Errorf("Database %s not found!", name)
	}
	database = client.Database(name)
	return nil
}

func GetCollections() []string {
	var err error
	collections, err = database.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return collections
}

func ChooseCollection(name string) (*mongo.Collection, error) {
	if !utility.StringSliceContains(collections, name) {
		return nil, fmt.Errorf("Collection %s not found!", name)
	}
	collection := database.Collection(name)
	return collection, nil
}

func GetSample(collection *mongo.Collection) (bson.M, error) {
	var result bson.M
	err := collection.FindOne(ctx, bson.M{}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CollectDocumentSamplesFromCollection(collection string) (bson.M, error) {
	c, err := ChooseCollection(collection)
	if err != nil {
		return nil, errors.Wrapf(err, "Error choosing collection %s", collection)
	}
	sample, err := GetSample(c)
	if err != nil {
		return nil, errors.Wrapf(err, "Error getting sample from collection %s", collection)
	}
	fmt.Println(sample)
	return sample, nil
}

func PerformAggregation(collection *mongo.Collection, pipeline []bson.M) ([]bson.M, error) {
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrapf(err, "Error performing aggregation")
	}
	results := []bson.M{}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrapf(err, "Error decoding aggregation results")
	}
	return results, nil
}
