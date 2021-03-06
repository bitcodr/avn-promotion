package service

import (
	"errors"
	"time"

	"github.com/bitcodr/avn-promotion/domain/model"
	"github.com/bitcodr/avn-promotion/helper"
)

var (
	ErrPromotionNotFound = errors.New("Promotion Not Found")
	ErrPromotionInvalid  = errors.New("Promotion Invalid")
)

type PromotionService interface {
	GetByPromotionCode(promotionCode string) (*model.Promotion, error)
	List() ([]*model.Promotion, error)
	Insert(promotion *model.Promotion) (*model.Promotion, error)
	Receivers(promotionCode string) ([]*model.Receiver, error)
	Verify(promotionCode string) (*model.Promotion, error)
	InsertReceiver(receiver *model.Receiver) (*model.Receiver, error)
}

type PromotionRepository interface {
	List() ([]*model.Promotion, error)
	Insert(promotion *model.Promotion) (*model.Promotion, error)
	GetByPromotionCode(promotionCode string) (*model.Promotion, error)
	Receivers(promotionCode string) ([]*model.Receiver, error)
	PromotionReceiversCount(usableTimes uint32, promotionCode string) error
	InsertReceiver(receiver *model.Receiver) (*model.Receiver, error)
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

func (p *promotionService) Receivers(promotionCode string) ([]*model.Receiver, error) {
	return p.promotionRepo.Receivers(promotionCode)
}

func (p *promotionService) Verify(promotionCode string) (*model.Promotion, error) {
	promotion, err := p.promotionRepo.GetByPromotionCode(promotionCode)
	if err != nil {
		return nil, errors.New("service.verify.promotion.notExist")
	}
	if promotion.ExpireDate < uint64(time.Now().Unix()) {
		return nil, errors.New("service.verify.promotion.expireTime")
	}
	if err := p.promotionRepo.PromotionReceiversCount(promotion.UsableTimes, promotionCode); err != nil {
		return nil, errors.New("service.verify.promotion.receivers")
	}
	return promotion, nil
}

func (p *promotionService) GetByPromotionCode(promotionCode string) (*model.Promotion, error) {
	return p.promotionRepo.GetByPromotionCode(promotionCode)
}

func (p *promotionService) InsertReceiver(receiver *model.Receiver) (*model.Receiver, error) {
	if err := helper.ValidateModel(receiver); err != nil {
		return nil, err
	}
	return p.promotionRepo.InsertReceiver(receiver)
}
