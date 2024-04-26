package dbs

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// ConnectMongodbURL returns a connection to a mongodb instance through a connection string in the configuration file
func ConnectMongodbURL() (*mongo.Client, error) {
	env := viper.GetString("env")

	connectURI := viper.GetString(env + ".db.mongo.url")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
