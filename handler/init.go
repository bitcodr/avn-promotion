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
	REST_GET_PROMOTION_RECEIVERS = "REST_GET_PROMOTION_RECEIVERS"
	REST_INSERT_PROMOTION        = "REST_INSERT_PROMOTION"
	REST_LIST_PROMOTIONS = "LIST_OF_PROMOTIONS"
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

	// router.HandleFunc("/promotions/{promotionCode}/receivers", promotionRestHandler.receivers).Methods(http.MethodGet).Name(REST_GET_PROMOTION_RECEIVERS)
	router.HandleFunc("/promotions", promotionRestHandler.Insert).Methods(http.MethodPost).Name(REST_INSERT_PROMOTION)
	router.HandleFunc("/promotions", promotionRestHandler.List).Methods(http.MethodGet).Name(REST_LIST_PROMOTIONS)

}

func GRPC(app *config.App) {
	//implement grpc handler route here
}
