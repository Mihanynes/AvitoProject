package models

import "github.com/jinzhu/gorm"

type Segment struct {
	gorm.Model
	Slug string `gorm:"not null" json:"slug"`
}

func (segment *Segment) Create() (*Segment, error) {
	err := DB.FirstOrCreate(&Segment{}, &segment).Error
	if err != nil {
		return segment, err
	}
	return segment, nil
}

func (segment *Segment) Delete() (*Segment, error) {
	err := DB.Where("slug = ?", segment.Slug).Delete(&Segment{}).Error
	if err != nil {
		return segment, err
	}
	return segment, nil
}
