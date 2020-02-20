package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
 
type Promotion struct {
	ID            primitive.ObjectID `json:"id" bson:"_id" msgpack:"id"`
	PromotionCode string             `json:"promotionCode" bson:"promotionCode" msgpack:"promotionCode" validate:"required"`
	UsableTimes   uint32             `json:"usableTimes" bson:"usableTimes" msgpack:"usableTimes" validate:"required"`
	ExpireDate    uint64             `json:"expireDate" bson:"expireDate" msgpack:"expireDate" validate:"required"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt" msgpack:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt" msgpack:"updatedAt"`
}
