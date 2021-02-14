//Package msgpack ...
package msgpack

import (
	"github.com/bitcodr/avn-promotion/domain/model"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack/v4"
)

type Promotion struct{}

func (m *Promotion) Encode(input *model.Promotion) ([]byte, error) {
	rawPromotion, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Promotion.Encode")
	}
	return rawPromotion, nil
}

func (m *Promotion) Decode(input []byte) (*model.Promotion, error) {
	promotionModel := new(model.Promotion)
	if err := msgpack.Unmarshal(input, promotionModel); err != nil {
		return nil, errors.Wrap(err, "serializer.Promotion.Decode")
	}
	return promotionModel, nil
}
