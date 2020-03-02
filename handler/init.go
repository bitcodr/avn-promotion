//Package promotion ...
package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amiraliio/avn-grpc-promotion-proto/proto"
	"github.com/amiraliio/avn-promotion/config"
	"github.com/amiraliio/avn-promotion/domain/model"
	"github.com/amiraliio/avn-promotion/domain/service"
	grpcHandler "github.com/amiraliio/avn-promotion/handler/grpc"
	restHandler "github.com/amiraliio/avn-promotion/handler/rest"
	"github.com/amiraliio/avn-promotion/repository/mongo"
	"github.com/gorilla/mux"
	nats "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

const (
	REST_GET_PROMOTION_RECEIVERS = "REST_GET_PROMOTION_RECEIVERS"
	REST_INSERT_PROMOTION        = "REST_INSERT_PROMOTION"
	REST_LIST_PROMOTIONS         = "LIST_OF_PROMOTIONS"
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

	promotionRestHandler := restHandler.NewRestPromotionHandler(promotionService)

	router.HandleFunc("/promotions/{promotionCode}/receivers", promotionRestHandler.Receivers).Methods(http.MethodGet).Name(REST_GET_PROMOTION_RECEIVERS)
	router.HandleFunc("/promotions", promotionRestHandler.Insert).Methods(http.MethodPost).Name(REST_INSERT_PROMOTION)
	router.HandleFunc("/promotions", promotionRestHandler.List).Methods(http.MethodGet).Name(REST_LIST_PROMOTIONS)
}

func GRPC(app *config.App, server *grpc.Server) {

	promotionRepo := choosePromotionRepo("mongo", app)

	promotionService := service.NewPromotionService(promotionRepo)

	promotionRestHandler := grpcHandler.NewGRPCPromotionHandler(promotionService)

	proto.RegisterPromotionServer(server, promotionRestHandler)
}

func NATS(app *config.App, c *nats.EncodedConn) {

	c.Subscribe("promotion.*", func(m *nats.Msg) {
		request := new(model.ChargeRequest)
		if err := json.Unmarshal(m.Data, request); err != nil {
			log.Println(err)
			return
		}
		promotionRepo := choosePromotionRepo("mongo", app)
		promotionService := service.NewPromotionService(promotionRepo)
		promotion, err := promotionService.GetByPromotionCode(request.PromotionCode)
		if err != nil {
			log.Println(err)
			return
		}
		if err := m.Sub.AutoUnsubscribe(int(promotion.UsableTimes)); err != nil {
			log.Println(err)
			return
		}
		receiver := new(model.Receiver)
		receiver.Promotion = promotion
		receiver.Cellphone = request.Cellphone
		receiver.Fullname = request.Fullname
		_, err = promotionService.InsertReceiver(receiver)
		if err != nil {
			log.Println(err)
			return
		}
	})
}
