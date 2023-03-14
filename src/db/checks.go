package database

import (
	str "beprj/src/Structs"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func (s *DBClient) CheckCourseExistance(cid int32) (bool, error) {
	var c str.Course
	err := s.Db.Where("ID=?", cid).First(&c).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		} //no need of this if ---just for understanding---remove this later
		return false, err
	}
	return true, nil
}

func (s *DBClient) CheckFacultyExistance(fid int32) (bool, error) {
	var fac str.Faculty
	err := s.Db.Where("ID=?", fid).First(&fac).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		} //no need of this if ---just for understanding---remove this later
		return false, err
	}
	return true, nil
}
