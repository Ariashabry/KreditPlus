package models

import (
	"gorm.io/gorm"
	"time"
)

type Pinjaman struct {
	IDPinjaman     uint      `json:"IDPinjaman" gorm:"column:id_pinjaman;primarykey;autoIncrement"`
	UserID         uint      `json:"UserID" gorm:"column:user_id;not null"`
	Amount         int       `json:"Amount" gorm:"column:amount;not null"`
	Tenor          int       `json:"Tenor" gorm:"column:tenor;not null"`
	Status         string    `json:"Status" gorm:"column:status;not null"`
	TotalDebt      int       `json:"TotalDebt" gorm:"column:totaldebt;not null"`
	InterestRate   float64   `json:"InterestRate" gorm:"column:interest_rate;not null"`
	MonthlyPayment int       `json:"MonthlyPayment" gorm:"-"`
	CreatedAt      time.Time `json:"CreatedAt" gorm:"column:created_at;autoCreateTime"`
}

func (Pinjaman) TableName() string {
	return "pinjaman"
}

func (p *Pinjaman) Create(db *gorm.DB) error {
	return db.Model(Pinjaman{}).Create(p).Error
}

func (p *Pinjaman) GetById(db *gorm.DB, Id int) error {
	return db.Model(Pinjaman{}).Where("id_pinjaman = ?", Id).First(p).Error
}

func (p *Pinjaman) Update(db *gorm.DB) error {
	return db.Model(Pinjaman{}).Omit("id_pinjaman").Where("id_pinjaman = ?", p.IDPinjaman).Updates(p).Error
}
