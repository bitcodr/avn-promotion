package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Promotion struct {
	ID            primitive.ObjectID `json:"id" bson:"_id" msgpack:"id" validate:"required"`
	PromotionCode string             `json:"promotionCode" bson:"promotionCode" msgpack:"promotionCode" validate:"required"`
	UsableTimes   uint32             `json:"usableTimes" bson:"usableTimes" msgpack:"usableTimes" validate:"required"`
	ExpireDate    primitive.DateTime `json:"expireDate" bson:"expireDate" msgpack:"expireDate" validate:"required"`
	CreatedAt     primitive.DateTime `json:"createdAt" bson:"createdAt" msgpack:"createdAt"`
	UpdatedAt     primitive.DateTime `json:"updatedAt" bson:"updatedAt" msgpack:"updatedAt"`
}
