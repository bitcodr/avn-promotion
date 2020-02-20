package service

import (
	"errors"

	"github.com/amiraliio/avn-promotion/domain/model"
	"github.com/amiraliio/avn-promotion/helper"
)

var (
	ErrPromotionNotFound = errors.New("Promotion Not Found")
	ErrPromotionInvalid  = errors.New("Promotion Invalid")
)

type PromotionService interface {
	// Get(cellphone uint64) (*model.Promotion, error)
	Insert(promotion *model.Promotion) (*model.Promotion, error)
}

type PromotionRepository interface {
	// Get(cellphone uint64) (*model.Promotion, error)
	Insert(promotion *model.Promotion) (*model.Promotion, error)
	GetByPromotionCode(promotionCode string) (*model.Promotion, error)
}

type PromotionSerializer interface {
	Encode(input *model.Promotion) ([]byte, error)
	Decode(input []byte) (*model.Promotion, error)
}

type promotionService struct {
	promotionRepo PromotionRepository
}

func NewPromotionService(promotionRepo PromotionRepository) PromotionService {
	return &promotionService{
		promotionRepo,
	}
}

// func (w *promotionService) Get(cellphone uint64) (*model.Promotion, error) {
// 	return w.promotionRepo.Get(cellphone)
// }

func (w *promotionService) Insert(promotion *model.Promotion) (*model.Promotion, error) {
	if err := helper.ValidateModel(promotion); err != nil {
		return nil, err
	}
	if _, err := w.promotionRepo.GetByPromotionCode(promotion.PromotionCode); err == nil {
		return nil, errors.New("service.insert.dupplicate.promotion")
	}
	return w.promotionRepo.Insert(promotion)
}
