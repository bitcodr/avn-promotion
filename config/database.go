//Package config ...
package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB config
func (app *App) DB() *mongo.Database {
	var config string
	if AppConfig.GetString("DATABASES.MONGO.USERNAME") != "" && AppConfig.GetString("DATABASES.MONGO.PASSWORD") != "" {
		config = "mongodb://" + AppConfig.GetString("DATABASES.MONGO.USERNAME") + ":" + AppConfig.GetString("DATABASES.MONGO.PASSWORD") + "@" + AppConfig.GetString("DATABASES.MONGO.HOST") + ":" + AppConfig.GetString("DATABASES.MONGO.PORT") + "/?authSource=" + AppConfig.GetString("DATABASES.MONGO.DATABASE")
	} else {
		config = "mongodb://" + AppConfig.GetString("DATABASES.MONGO.HOST") + ":" + AppConfig.GetString("DATABASES.MONGO.PORT")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config))
	if err != nil {
		log.Fatal(err.Error())
	}
	return client.Database(AppConfig.GetString("DATABASES.MONGO.DATABASE"))
}
