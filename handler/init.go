//Package promotion ...
package handler

import (
	"net/http"

	"github.com/amiraliio/avn-promotion/config"
	"github.com/amiraliio/avn-promotion/domain/service"
	"github.com/amiraliio/avn-promotion/handler/rest"
	"github.com/amiraliio/avn-promotion/repository/mongo"
	"github.com/gorilla/mux"
)

const (
	REST_GET_PROMOTION          = "REST_GET_PROMOTION"
	REST_INSERT_PROMOTION       = "REST_INSERT_PROMOTION"
)

func choosePromotionRepo(connection string, app *config.App) service.PromotionRepository {
	switch connection {
	case "mongo":
		return mongo.NewMongoPromotionRepository(app)
	default:
		return nil
	}
}

func HTTP(app *config.App, router *mux.Router) {

	promotionRepo := choosePromotionRepo("mongo", app)

	promotionService := service.NewPromotionService(promotionRepo)

	promotionRestHandler := rest.NewRestPromotionHandler(promotionService)

	// router.HandleFunc("/promotion/{cellphone}", promotionRestHandler.Get).Methods(http.MethodGet).Name(REST_GET_PROMOTION)
	router.HandleFunc("/promotion", promotionRestHandler.Insert).Methods(http.MethodPost).Name(REST_INSERT_PROMOTION)

}

func GRPC(app *config.App) {
	//implement grpc handler route here
}
