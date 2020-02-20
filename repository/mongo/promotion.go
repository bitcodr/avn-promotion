//Package mongo ...
package mongo

import (
	"context"
	"time"

	"github.com/amiraliio/avn-promotion/config"
	"github.com/amiraliio/avn-promotion/domain/model"
	"github.com/amiraliio/avn-promotion/domain/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PROMOTION_COLLECTION = "promotion"

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
