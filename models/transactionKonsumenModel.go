package models

import "time"

type TransactionKonsumen struct {
	ID            uint      `gorm:"column:id;primarykey;autoIncrement;"`
	KonsumenID    uint      `gorm:"column:konsumen_id;not null"`
	OTR           int       `gorm:"column:amount;not null"`
	AdminFee      time.Time `gorm:"column:payment_date;not null"`
	JumlahCicilan time.Time `gorm:"column:created_at;autoCreateTime"`
	JumlahBunga   float64   `gorm:"column:jumlah_bunga"`
	NamaBarang    string    `gorm:"column:nama_barang"`
}

func (TransactionKonsumen) TableName() string {
	return "transaction_konsumen"
}
