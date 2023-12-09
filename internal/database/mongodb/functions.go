package mongodb

import (
	"context"
	"time"

	"github.com/fonteeboa/go-smtp-hub/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertSMTPConfig inserts a base log into the specified MongoDB collection.
//
// Client is the MongoDB client used to connect to the database.
// log is the base log to be inserted.
// Returns an error if the insertion fails.
func InsertSMTPConfig(Client *mongo.Client, log domain.SMTPConfig) error {
	collection := getCollection(Client, "smtpConfig")

	_, err := collection.InsertOne(context.Background(), log)
	return err
}

// GetDataMongo retrieves SMTP configurations from MongoDB.
//
// client: the MongoDB client used to connect to the database.
// Returns a slice of SMTPConfig structs and an error, if any.
func GetDataMongo(client *mongo.Client) ([]*domain.SMTPConfig, error) {
	collection := getCollection(client, "smtpConfig")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var configs []*domain.SMTPConfig

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var config *domain.SMTPConfig
		if err := cur.Decode(&config); err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return configs, nil
}
