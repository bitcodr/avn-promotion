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
	List() ([]*model.Promotion, error)
	Insert(promotion *model.Promotion) (*model.Promotion, error)
}

type PromotionRepository interface {
	List() ([]*model.Promotion, error)
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

func (p *promotionService) List() ([]*model.Promotion, error) {
	return p.promotionRepo.List()
}

func (p *promotionService) Insert(promotion *model.Promotion) (*model.Promotion, error) {
	if err := helper.ValidateModel(promotion); err != nil {
		return nil, err
	}
	if _, err := p.promotionRepo.GetByPromotionCode(promotion.PromotionCode); err == nil {
		return nil, errors.New("service.insert.dupplicate.promotion")
	}
	return p.promotionRepo.Insert(promotion)
}
