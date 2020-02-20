package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Receivers struct {
	ID        primitive.ObjectID `json:"id" bson:"_id" msgpack:"id" validate:"required"`
	Fullname  string             `json:"fullname" bson:"fullname" msgpack:"fullname" validate:"required"`
	Cellphone uint64             `json:"cellphone" bson:"cellphone" msgpack:"cellphone" validate:"required"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt" msgpack:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt" msgpack:"updatedAt"`
}
