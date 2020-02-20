//Package json ...
package json

import (
	"encoding/json"

	"github.com/amiraliio/avn-promotion/domain/model"
	"github.com/pkg/errors"
)

type Promotion struct{}

func (m *Promotion) Encode(input *model.Promotion) ([]byte, error) {
	rawPromotion, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Promotion.Encode")
	}
	return rawPromotion, nil
}

func (m *Promotion) Decode(input []byte) (*model.Promotion, error) {
	promotionModel := new(model.Promotion)
	if err := json.Unmarshal(input, promotionModel); err != nil {
		return nil, errors.Wrap(err, "serializer.Promotion.Decode")
	}
	return promotionModel, nil
}
