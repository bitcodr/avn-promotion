package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppInterface interface {
	Environment()
	DB() *mongo.Database
}

type App struct {
	CurrentTime time.Time
}

var (
	LangConfig *viper.Viper
)

func (app *App) Init() {
	app.langConfig()
}

func (app *App) langConfig() {
	LangConfig = viper.New()
	LangConfig.SetConfigType("yaml")
	LangConfig.SetConfigName("lang")
	LangConfig.AddConfigPath(os.Getenv("PROMOTION_SERVICE_ROOT_DIR") + "/resource/lang")
	err := LangConfig.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
