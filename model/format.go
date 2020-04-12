package model

import "github.com/jinzhu/gorm"

// Format is struct
type Format struct {
	gorm.Model
	Name string
}

// NewFormat is constructor
func NewFormat(name string) *Format {
	return &Format{Name: name}
}

// SetName is setter of Name
func (f *Format) SetName(name string) {
	f.Name = name
}

// FindByID is
func (f *Format) FindByID(db *gorm.DB, id int) (*Format, error) {
	var format Format
	if error := db.Find(&format).Error; error != nil {
		return nil, error
	}
	return &format, nil
}

// FindAll is
func (f *Format) FindAll(db *gorm.DB) (*[]Format, error) {
	var formats []Format
	if error := db.Find(&formats).Error; error != nil {
		return nil, error
	}
	return &formats, nil
}

// Create is
func (f *Format) Create(db *gorm.DB) (*Format, error) {
	if error := db.Create(f).Error; error != nil {
		return nil, error
	}
	return f, nil
}