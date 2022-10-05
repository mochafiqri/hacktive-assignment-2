package repo

import (
	"assignment2/models"
	"gorm.io/gorm"
)

type OrderRepo interface {
	Create(cmd *models.Order, trx ...*gorm.DB) error
	Get(id int, trx ...*gorm.DB) (models.Order, error)
	Update(cmd *models.Order, trx ...*gorm.DB) error
	Delete(id int, trx ...*gorm.DB) error
}

type OrderRepoContract struct {
	db *gorm.DB
}

func (o *OrderRepoContract) Delete(id int, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = o.NewDb(trx...)
	var err = newDb.Where("id = ?", id).Delete(&models.Order{}).Error

	return err
}

func (o *OrderRepoContract) Update(cmd *models.Order, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = o.NewDb(trx...)

	var err = newDb.Model(&models.Order{}).
		Where("id = ?", cmd.Id).Updates(cmd).Error

	return err
}

func (o *OrderRepoContract) Get(id int, trx ...*gorm.DB) (models.Order, error) {
	//TODO implement me
	var result = models.Order{}
	var newDb = o.NewDb(trx...)

	var err = newDb.Model(&models.Order{}).
		Where("id = ?", id).
		Preload("Item").
		Find(&result).Error
	if err != nil {
		return result, nil
	}
	return result, nil
}

func (o *OrderRepoContract) Create(cmd *models.Order, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = o.NewDb(trx...)

	var err = newDb.Model(&models.Order{}).Create(cmd).Error

	return err
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &OrderRepoContract{db: db}
}

func (o *OrderRepoContract) NewDb(trx ...*gorm.DB) *gorm.DB {
	if len(trx) > 0 {
		return trx[0]
	} else {
		return o.db
	}
}
