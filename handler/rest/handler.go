//Package rest ...
package rest

import (
	"net/http"
	"strconv"

	"github.com/amiraliio/avn-promotion/config"
	"github.com/amiraliio/avn-promotion/domain/model"
	"github.com/amiraliio/avn-promotion/domain/service"
	"github.com/amiraliio/avn-promotion/helper"
	"github.com/amiraliio/avn-promotion/serializer/json"
	"github.com/amiraliio/avn-promotion/serializer/msgpack"
	"github.com/gorilla/mux"
)

type PromotionHandler interface {
	Get(res http.ResponseWriter, req *http.Request)
	Insert(res http.ResponseWriter, req *http.Request)
}

type promotionHandler struct {
	promotionService service.PromotionService
}

func NewRestPromotionHandler(promotionService service.PromotionService) PromotionHandler {
	return &promotionHandler{
		promotionService,
	}
}

func (h *promotionHandler) serializer(contentType string) service.PromotionSerializer {
	switch contentType {
	case "application/json":
		return &json.Promotion{}
	case "application/x-msgpack":
		return &msgpack.Promotion{}
	default:
		return &json.Promotion{}
	}
}

func (w *promotionHandler) Get(res http.ResponseWriter, req *http.Request) {
	// acceptHeader := req.Header.Get("Accept")
	// params := mux.Vars(req)
	// if params == nil {
	// 	helper.ResponseError(res, nil, http.StatusUnprocessableEntity, acceptHeader, "W-1000", config.LangConfig.GetString("MESSAGES.PARAM_EMPTY"))
	// 	return
	// }
	// cellphone, err := strconv.ParseUint(params["cellphone"], 10, 64)
	// if err != nil {
	// 	helper.ResponseError(res, err, http.StatusInternalServerError, acceptHeader, "W-1001", config.LangConfig.GetString("MESSAGES.PARSE_CELLPHONE"))
	// 	return
	// }
	// promotion, err := w.promotionService.Get(cellphone)
	// if err != nil {
	// 	helper.ResponseError(res, err, http.StatusNotFound, acceptHeader, "W-1002", config.LangConfig.GetString("MESSAGES.DATA_NOT_FOUND"))
	// 	return
	// }
	// helper.ResponseOk(res, http.StatusOK, acceptHeader, promotion)
}

func (w *promotionHandler) Insert(res http.ResponseWriter, req *http.Request) {
	// acceptHeader := req.Header.Get("Accept")
	// promotionCode := req.FormValue("promotionCode")
	// if promotionCode == "" {
	// 	helper.ResponseError(res, nil, http.StatusUnprocessableEntity, acceptHeader, "W-1003", config.LangConfig.GetString("MESSAGES.PROMOTION_CODE_IS_REQUIRED"))
	// 	return
	// }
	// //TODO get promotion code is verified from promotion server with grpc
	// promotionModel := new(model.Promotion)
	// promotion, err := w.promotionService.Insert(promotionModel)
	// if err != nil {
	// 	helper.ResponseError(res, err, http.StatusNotFound, acceptHeader, "W-1004", config.LangConfig.GetString("MESSAGES.DATA_NOT_FOUND"))
	// 	return
	// }
	// //TODO send an event to promotion server who get the promotion
	// //TODO waite for acknowledge from broker
	// helper.ResponseOk(res, http.StatusOK, acceptHeader, promotion)
}

