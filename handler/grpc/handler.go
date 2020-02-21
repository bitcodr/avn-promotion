package grpc

import (
	"context"

	"github.com/amiraliio/avn-grpc-promotion-proto/proto"
	"github.com/amiraliio/avn-promotion/domain/service"
)

type PromotionHandler interface {
	Verify(ctx context.Context, req *proto.Request) (*proto.Response, error)
}

type promotionGRPCHandler struct {
	promotionService service.PromotionService
}

func NewGRPCPromotionHandler(promotionService service.PromotionService) PromotionHandler {
	return &promotionGRPCHandler{
		promotionService,
	}
}

func (p *promotionGRPCHandler) Verify(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	promotion, err := p.promotionService.Verify(req.GetPromotionCode())
	if err != nil {
		return nil, err
	}
	response := new(proto.Response)
	response.Charge = promotion.Charge
	return response, nil
}
