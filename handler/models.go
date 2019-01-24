package handler

import (
	"github.com/jinzhu/gorm"
)

// Employees - Table Emloyees
type Employees struct {
	gorm.Model
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Phone   int    `json:"Phone"`
	Hobby   string `json:"Hobby"`
}

// CreateTable - Employees
func (t Employees) CreateTable(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&Employees{}).Error

	return
}

// InsertData -
func (t *Employees) InsertData(db *gorm.DB) (err error) {
	err = db.Create(&t).Error

	return
}

// UpdateData -
func (t Employees) UpdateData(db *gorm.DB, inputdata *Employees) (res *Employees, err error) {
	data := new(Employees)

	db.Where("name = ?", t.Name).First(&data)

	data.Name = inputdata.Name
	data.Address = inputdata.Address

	err = db.Save(&data).Error
	res = data
	return
}

// DeleteData -
func (t Employees) DeleteData(db *gorm.DB) (err error) {
	db.Where("name = ?", t.Name).Delete(&Employees{})
	// err = db.Delete(&t).Error

	return
}

// GetDataByName -
func (t Employees) GetDataByName(db *gorm.DB) (res Employees, err error) {
	err = db.Where("name = ?", t.Name).First(&res).Error

	return
}

// GetAllData -
func (t Employees) GetAllData(db *gorm.DB) (res []Employees, err error) {
	err = db.Find(&res).Error

	return
}
