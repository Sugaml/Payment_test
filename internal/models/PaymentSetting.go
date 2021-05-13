package models

import (
	"github.com/jinzhu/gorm"
)

type PaymentSetting struct {
	gorm.Model
	User        *User  `gorm:"foreignkey:UserID" json:"user"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Country     string `gorm:"not null" json:"country"`
	State       string `gorm:"not null" json:"state"`
	City        string `gorm:"not null" json:"city"`
	Street      string `gorm:"not null" json:"street"`
	Postal_Code string `gorm:"not null" json:"postal_code"`
}

func (data *PaymentSetting) Save(db *gorm.DB) (*PaymentSetting, error) {
	err := db.Model(&PaymentSetting{}).Create(&data).Error
	if err != nil {
		return &PaymentSetting{}, err
	}
	return data, nil
}
func (data *PaymentSetting) FindAll(db *gorm.DB) (*[]PaymentSetting, error) {
	var err error
	datas := []PaymentSetting{}
	err = db.Model(&PaymentSetting{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]PaymentSetting{}, err
	}
	return &datas, nil
}

func (data *PaymentSetting) Find(db *gorm.DB, pid uint64) (*PaymentSetting, error) {
	var err error
	err = db.Model(&PaymentSetting{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &PaymentSetting{}, err
	}
	return data, nil
}
func (data *PaymentSetting) Update(db *gorm.DB) (*PaymentSetting, error) {
	var err error
	err = db.Model(&PaymentSetting{}).Update(&data).Error
	if err != nil {
		return &PaymentSetting{}, err
	}
	return data, nil
}
func (data *PaymentSetting) Delete(db *gorm.DB) (*PaymentSetting, error) {
	var err error
	err = db.Model(&PaymentSetting{}).Delete(&data).Error
	if err != nil {
		return &PaymentSetting{}, err
	}
	return data, nil
}
