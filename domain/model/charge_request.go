package model

type ChargeRequest struct {
	PromotionCode string    `json:"promotionCode" msgpack:"promotionCode" validate:"required"`
	Cellphone     uint64    `json:"cellphone" msgpack:"cellphone" validate:"required"`
	Fullname      string    `json:"fullname" msgpack:"fullname"`
	_             struct{}  //to prevent unkeyed literals
	_             [0]func() //to prevent struct compration
}
