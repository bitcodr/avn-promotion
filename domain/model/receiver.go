package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Receiver struct {
	ID        primitive.ObjectID `json:"id" bson:"_id" msgpack:"id"`
	Fullname  string             `json:"fullname" bson:"fullname" msgpack:"fullname" validate:"required"`
	Cellphone uint64             `json:"cellphone" bson:"cellphone" msgpack:"cellphone" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" msgpack:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" msgpack:"updatedAt"`
}
