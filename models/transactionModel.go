package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	IDTransaction uint      `gorm:"column:id_transaction;primarykey;autoIncrement;"`
	LoanID        uint      `gorm:"column:loan_id;not null"`
	UserID        uint      `gorm:"column:user_id;not null"`
	Amount        int       `gorm:"column:amount;not null"`
	PaymentDate   time.Time `gorm:"column:payment_date;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Transaction) TableName() string {
	return "transaction"
}

func (t *Transaction) Create(db *gorm.DB) error {
	return db.Model(Transaction{}).Create(t).Error
}
