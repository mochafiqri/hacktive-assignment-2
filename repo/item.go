package repo

import (
	"assignment2/models"
	"gorm.io/gorm"
)

type ItemRepo interface {
	Creates(cmd *[]models.OrderItem, trx ...*gorm.DB) error
	Updates(save *[]models.OrderItem, trx ...*gorm.DB) error
	DeleteByOrder(orderId int, trx ...*gorm.DB) error
}

type ItemRepoContract struct {
	db *gorm.DB
}

func (i *ItemRepoContract) DeleteByOrder(orderId int, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = i.NewDb(trx...)

	var err = newDb.Where("order_id = ?", orderId).Delete(&models.OrderItem{}).Error
	return err
}

func (i *ItemRepoContract) Updates(save *[]models.OrderItem, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = i.NewDb(trx...)

	var err = newDb.Save(save).Error

	return err
}

func (i *ItemRepoContract) Creates(cmd *[]models.OrderItem, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = i.NewDb(trx...)

	var err = newDb.Create(cmd).Error

	return err
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return &ItemRepoContract{db: db}
}

func (i *ItemRepoContract) NewDb(trx ...*gorm.DB) *gorm.DB {
	if len(trx) > 0 {
		return trx[0]
	} else {
		return i.db
	}
}
