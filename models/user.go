package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID    uint `gorm:"not null" json:"user_id"`
	SegmentID uint `gorm:"not null" json:"segment_id"`
}

func AddUserToSegment(segmentsToAdd [](Segment), segmentsToDelete [](Segment), userID uint) error {
	var err error
	for _, segmentToAdd := range segmentsToAdd {
		//err = DB.Create(&User{UserID: userID, SegmentID: segmentToAdd.ID}).Error
		err = DB.FirstOrCreate(&User{}, &User{UserID: userID, SegmentID: segmentToAdd.ID}).Error
		if err != nil {
			return err
		}
	}

	for _, segmentToDelete := range segmentsToDelete {
		err = DB.Where("user_id = ? AND segment_id = ?", userID, segmentToDelete.ID).Delete(&User{}).Error
		if err != nil {
			return err
		}

	}
	return nil
}

func ActiveSegments(userId uint) ([]Segment, error) {
	var segments_users []User
	var err error
	err = DB.Where(&User{UserID: userId}).Find(&segments_users).Error
	var segments []Segment

	for _, userSegment := range segments_users {
		var segment Segment
		err = DB.Where(&Segment{
			Model: gorm.Model{ID: userSegment.SegmentID},
		}).Find(&segment).Error
		segments = append(segments, segment)
	}

	if err != nil {
		return segments, err
	}
	return segments, nil

}
