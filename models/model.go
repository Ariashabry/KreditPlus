package models

import "gorm.io/gorm"

func MigrateModel(db *gorm.DB) error {
	return db.AutoMigrate(&Konsumen{}, &Pinjaman{}, &Transaction{}, &TransactionKonsumen{})
}
