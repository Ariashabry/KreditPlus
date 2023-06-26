package models

import "gorm.io/gorm"

type Konsumen struct {
	IdKonsumen   uint       `json:"IdKonsumen" gorm:"column:id_konsumen;primaryKey;autoIncrement;"`
	NIK          string     `json:"NIK" gorm:"column:nik;"`
	Gender       string     `json:"Gender" gorm:"column:gender;"`
	FullName     string     `json:"FullName" gorm:"column:full_name;"`
	LegalName    string     `json:"LegalName" gorm:"column:legal_name;"`
	TempatLahir  string     `json:"TempatLahir" gorm:"column:tempat_lahir;"`
	TanggalLahir string     `json:"TanggalLahir" gorm:"column:tanggal_lahir;"`
	Gaji         uint       `json:"Gaji" gorm:"column:gaji;"`
	FotoKTP      string     `json:"FotoKTP" gorm:"foto_ktp"`
	FotoSelfie   string     `json:"FotoSelfie" gorm:"foto_selfie"`
	TotalDebt    int        `json:"-" gorm:"-"`
	Pinjamans    []Pinjaman `json:"Pinjaman" gorm:"foreignKey:UserID"`
}

func (Konsumen) TableName() string {
	return "konsumen"
}

type Konsumens []Konsumen

func (m *Konsumen) Create(db *gorm.DB) error {
	return db.Model(Konsumen{}).Create(m).Error
}

func (m *Konsumen) GetById(db *gorm.DB, Id int) error {
	return db.Model(Konsumen{}).Preload("Pinjamans").Where("id_konsumen = ?", Id).First(m).Error
}

func (m *Konsumen) Update(db *gorm.DB) error {
	return db.Model(Konsumen{}).Omit("id_konsumen").Where("id_konsumen = ?", m.IdKonsumen).Updates(m).Error
}

func (m *Konsumen) Delete(db *gorm.DB) error {
	return db.Model(Konsumen{}).Where("id_konsumen = ?", m.IdKonsumen).Delete(m).Error
}

func (m *Konsumens) GetAll(db *gorm.DB) error {
	return db.Model(Konsumen{}).Find(m).Error
}
