//Package mysql ...
package mysql

import (
	"github.com/amiraliio/avn-promotion/config"
	"github.com/amiraliio/avn-promotion/domain/model"
	"github.com/amiraliio/avn-promotion/domain/service"
)

type promotionRepo struct {
	app *config.App
}

func NewMysqlPromotionRepository(app *config.App) service.PromotionRepository {
	return &promotionRepo{
		app,
	}
}

func (w *promotionRepo) Get(cellphone uint64) (*model.Promotion, error) {
	// db := w.app.DB()
	return nil, nil
}

func (w *promotionRepo) Insert(promotion *model.Promotion) (*model.Promotion, error) {
	// db := w.app.DB()
	return nil, nil
}
