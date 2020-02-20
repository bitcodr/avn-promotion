//Package rest ...
package rest

import (
	"io/ioutil"
	"net/http"

	"github.com/amiraliio/avn-promotion/config"
	"github.com/amiraliio/avn-promotion/domain/service"
	"github.com/amiraliio/avn-promotion/helper"
	"github.com/amiraliio/avn-promotion/serializer/json"
	"github.com/amiraliio/avn-promotion/serializer/msgpack"
)

type PromotionHandler interface {
	List(res http.ResponseWriter, req *http.Request)
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

func (w *promotionHandler) serializer(contentType string) service.PromotionSerializer {
	switch contentType {
	case "application/json":
		return &json.Promotion{}
	case "application/x-msgpack":
		return &msgpack.Promotion{}
	default:
		return &json.Promotion{}
	}
}

func (w *promotionHandler) Insert(res http.ResponseWriter, req *http.Request) {
	contentTypeHeader := req.Header.Get("Content-Type")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		helper.ResponseError(res, err, http.StatusNotFound, contentTypeHeader, "P-1001", config.LangConfig.GetString("MESSAGES.BODY_ERROR"))
		return
	}
	promotionModel, err := w.serializer(contentTypeHeader).Decode(body)
	if err != nil {
		helper.ResponseError(res, err, http.StatusNotFound, contentTypeHeader, "P-1002", config.LangConfig.GetString("MESSAGES.SERIALIZER_ERROR"))
		return
	}
	promotion, err := w.promotionService.Insert(promotionModel)
	if err != nil {
		helper.ResponseError(res, err, http.StatusNotFound, contentTypeHeader, "P-1003", config.LangConfig.GetString("MESSAGES.DATA_NOT_FOUND"))
		return
	}
	helper.ResponseOk(res, http.StatusOK, contentTypeHeader, promotion)
}

func (w *promotionHandler) List(res http.ResponseWriter, req *http.Request) {
	acceptHeader := req.Header.Get("Accept")
	promotions, err := w.promotionService.List()
	if err != nil {
		helper.ResponseError(res, err, http.StatusNotFound, acceptHeader, "P-1004", config.LangConfig.GetString("MESSAGES.DATA_NOT_FOUND"))
		return
	}
	helper.ResponseOk(res, http.StatusOK, acceptHeader, promotions)
}

