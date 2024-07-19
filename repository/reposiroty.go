package repository

import (
	"block-banter/models"
	"gorm.io/gorm"
)

type TransferEvent struct {
	ID     uint          `gorm:"primaryKey"`
	From   string        `gorm:"type:varchar(42)"`
	To     string        `gorm:"type:varchar(42)"`
	Value  models.BigInt `gorm:"type:numeric"`
	TxHash string        `gorm:"type:varchar(66)"`
}

type TransferEventRepository struct {
	db *gorm.DB
}

func NewTransferEventRepository(db *gorm.DB) *TransferEventRepository {
	return &TransferEventRepository{db}
}

func (r *TransferEventRepository) Create(event *TransferEvent) error {
	return r.db.Create(event).Error
}

func (r *TransferEventRepository) GetByID(id uint) (*TransferEvent, error) {
	var event TransferEvent
	err := r.db.First(&event, id).Error
	return &event, err
}

func (r *TransferEventRepository) Update(event *TransferEvent) error {
	return r.db.Save(event).Error
}

func (r *TransferEventRepository) Delete(id uint) error {
	return r.db.Delete(&TransferEvent{}, id).Error
}

func (r *TransferEventRepository) List() ([]TransferEvent, error) {
	var events []TransferEvent
	err := r.db.Find(&events).Error
	return events, err
}
