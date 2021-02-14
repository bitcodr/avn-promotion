//Package mongo ...
package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/bitcodr/avn-promotion/config"
	"github.com/bitcodr/avn-promotion/domain/model"
	"github.com/bitcodr/avn-promotion/domain/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PROMOTION_COLLECTION = "promotion"
const RECEIVERS_COLLECTION = "receivers"

type promotionRepo struct {
	app *config.App
}

func NewMongoPromotionRepository(app *config.App) service.PromotionRepository {
	return &promotionRepo{
		app,
	}
}

func (w *promotionRepo) GetByPromotionCode(promotionCode string) (promotion *model.Promotion, err error) {
	db := w.app.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.Collection(PROMOTION_COLLECTION).FindOne(ctx, primitive.M{"promotionCode": promotionCode}).Decode(&promotion); err != nil {
		return nil, err
	}
	return promotion, nil
}

func (w *promotionRepo) Insert(promotion *model.Promotion) (*model.Promotion, error) {
	db := w.app.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	promotion.CreatedAt = w.app.CurrentTime
	promotion.ID = primitive.NewObjectID()
	document, err := db.Collection(PROMOTION_COLLECTION).InsertOne(ctx, promotion)
	if err != nil {
		return nil, err
	}
	promotion.ID = document.InsertedID.(primitive.ObjectID)
	return promotion, nil
}

func (w *promotionRepo) List() ([]*model.Promotion, error) {
	db := w.app.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.Collection(PROMOTION_COLLECTION).Find(ctx, primitive.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var promotions []*model.Promotion
	for cursor.Next(ctx) {
		promotion := new(model.Promotion)
		if err := cursor.Decode(&promotion); err != nil {
			return nil, err
		}
		promotions = append(promotions, promotion)
	}
	return promotions, nil
}

func (w *promotionRepo) Receivers(promotionCode string) ([]*model.Receiver, error) {
	db := w.app.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.Collection(RECEIVERS_COLLECTION).Find(ctx, primitive.M{"promotion.promotionCode": promotionCode})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var receivers []*model.Receiver
	for cursor.Next(ctx) {
		receiver := new(model.Receiver)
		if err := cursor.Decode(&receiver); err != nil {
			return nil, err
		}
		receivers = append(receivers, receiver)
	}
	return receivers, nil
}

func (w *promotionRepo) PromotionReceiversCount(usableTimes uint32, promotionCode string) error {
	db := w.app.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err := db.Collection(RECEIVERS_COLLECTION).CountDocuments(ctx, primitive.M{"promotion.promotionCode": promotionCode})
	if err != nil || uint32(count) == usableTimes {
		return errors.New("repository.PromotionReceiversCount")
	}
	return nil
}

func (w *promotionRepo) InsertReceiver(receiver *model.Receiver) (*model.Receiver, error) {
	db := w.app.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	receiver.CreatedAt = w.app.CurrentTime
	receiver.ID = primitive.NewObjectID()
	document, err := db.Collection(RECEIVERS_COLLECTION).InsertOne(ctx, receiver)
	if err != nil {
		return nil, err
	}
	receiver.ID = document.InsertedID.(primitive.ObjectID)
	return receiver, nil
}
