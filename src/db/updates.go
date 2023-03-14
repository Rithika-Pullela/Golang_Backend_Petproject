package database

import (
	str "beprj/src/Structs"
	"fmt"
	"time"
	//"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func (s *DBClient) UpdateAssignment(x str.Assignment, aid int32) (str.Assignment, error) {
	var a str.Assignment
	err := s.Db.Where("ID=?", aid).First(&a).Error
	if err != nil {
		return str.Assignment{}, err
	}
	if x.Deadline.Before(time.Now()) {
		return str.Assignment{}, fmt.Errorf("the assignment deadline is in the past")
	}
	err1 := s.Db.Model(&a).Updates(str.Assignment{Title: x.Title, Description: x.Description, Deadline: x.Deadline}).Error
	if err1 != nil {
		return str.Assignment{}, err1
	}
	return a, nil
}

func (s *DBClient) UpdateCourse(fid int32, c_name string, cid int32) (str.Course, error) {
	var c str.Course
	err := s.Db.Where("ID=?", cid).First(&c).Error
	if err != nil {
		return str.Course{}, err
	}
	err1 := s.Db.Model(&c).Updates(str.Course{Name: c_name, FacultyID: uint(fid)}).Error
	if err1 != nil {
		return str.Course{}, err1
	}
	return c, nil
}
