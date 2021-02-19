//Package config ...
package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB config
func (app *App) DB() *mongo.Database {
	var config string
	if os.Getenv("MONGO_DB_USERNAME") != "" && os.Getenv("MONGO_DB_PASSWORD") != "" {
		config = "mongodb://" + os.Getenv("MONGO_DB_USERNAME") + ":" + os.Getenv("MONGO_DB_PASSWORD") + "@" + os.Getenv("MONGO_DB_HOST") + ":" + "/?authSource=" + os.Getenv("MONGO_DB_NAME")
	} else {
		config = "mongodb://" + os.Getenv("MONGO_DB_HOST")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config))
	if err != nil {
		log.Fatal(err.Error())
	}
	return client.Database(os.Getenv("MONGO_DB_NAME"))
}
