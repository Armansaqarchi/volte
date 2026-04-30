package databases

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log/slog"
)

var (
	username = flag.String("mongo_username", "", "MongoDB username.")
	password = flag.String("mongo_password", "", "MongoDB password.")
	host     = flag.String("mongo_host", "", "MongoDB host.")
)

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClientWithConfig(opts *options.ClientOptions) *MongoClient {
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	// Sends a ping to confirm a successful connection.
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	slog.Info("Successfully connected to MongoDB!")
	return &MongoClient{
		client: client,
	}
}

func NewMongoClient() *MongoClient {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://%s:%s@%s", *username, *password, *host),
	).SetServerAPIOptions(serverAPI)
	fmt.Println(fmt.Sprintf(
		"mongodb://%s:%s@%s", *username, *password, *host),
	)
	return NewMongoClientWithConfig(opts)
}

func (m *MongoClient) CreateCollectionSchema(database string, collection string) {

	if err := m.client.Database(database).CreateCollection(context.Background(), collection); err != nil {
		var we mongo.CommandError
		if errors.As(err, &we) && we.Code == 48 {
			slog.Info("Collection already exists, skipping creation phase.")
		} else {
			panic(err)
		}
	}
	slog.Info(fmt.Sprintf("Created collection %s", collection))
}

func (m *MongoClient) GetClient() *mongo.Client {
	return m.client
}

//func (m *MongoClient) Get(ctx context.Context, keyID string, out interface{}) error {
//	return m.client.Database(database).Collection(collection).FindOne(ctx, bson.M{"_id": keyID}).Decode(out)
//}
//
//func (m *MongoClient) Set(ctx context.Context, collection string, keyID string, value interface{}) error {
//	return m.client.Database(*database).Collection(collection).FindOneAndUpdate(
//		ctx, bson.M{"_id": keyID}, bson.M{"$set": value}, options.FindOneAndUpdate().SetUpsert(true),
//	).Err()
//}
