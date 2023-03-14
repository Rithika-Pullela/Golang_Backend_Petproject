package database

import (
	str "beprj/src/Structs"
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Database interface {

	//implement db level func

	CreateCourse(string) (int, error)
	AddFacToCourse(int32, int) error

	CreateStudent(string, string, string) (int, error)
	CreateFaculty(string, string, string) (int, error)

	AddCoursetoStd(int32, int) error
	CreateAssignment(string, string, string, int32) (int, error)

	GetStudent(string) (str.Student, error)
	GetFaculty(string) (str.Faculty, error)
	GetCoursesEnrolledByStudentEmail(string) ([]str.Course, error)
	GetCoursesByFacultyEmail(string) ([]str.Course, error)

	GetCourseByCourseId(int32) (str.Course, error)
	GetAllCourses() ([]str.Course, error)
	GetFacultyByCourseId(int32) (str.Faculty, error)
	GetStudentsEnrolledByCourseId(int32) ([]str.Student, error)
	GetAssignmentsByCourseId(int32) ([]str.Assignment, error)

	CheckFacultyExistance(int32) (bool, error)
	CheckCourseExistance(int32) (bool, error)

	UpdateCourse(int32, string, int32) (str.Course, error)
	UpdateAssignment(str.Assignment, int32) (str.Assignment, error)
}

type DBClient struct {
	Db *gorm.DB
}

func (s *DBClient) CreateAssignment(title string, des string, Deadline string, cid int32) (int, error) {
	layout := "01/02/2006"
	dateString := Deadline
	date, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return 0, fmt.Errorf(err.Error())
	}

	var a str.Assignment
	a.CourseID = uint(cid)
	a.Deadline = date
	a.Description = des
	a.Title = title

	if a.Deadline.Before(time.Now()) {
		return 0, fmt.Errorf("the assignment deadline is in the past")
	}
	err1 := s.Db.Create(&a).Error
	if err1 != nil {
		return 0, fmt.Errorf(err1.Error())
	}
	return int(a.ID), nil
}

func (s *DBClient) AddCoursetoStd(sId int32, cId int) error {
	var c str.Course
	err := s.Db.Preload("Students").Preload("Assignments").Where("ID=?", cId).First(&c).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	var std str.Student
	err = s.Db.Preload("Courses").Where("ID=?", sId).First(&std).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	std.Courses = append(std.Courses, c)
	err = s.Db.Save(&std).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (s *DBClient) CreateStudent(fname string, lname string, email string) (int, error) {
	var std str.Student
	std.FirstName = fname
	std.LastName = lname
	std.Email = email

	err := s.Db.Create(&std).Error
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return int(std.ID), nil
}

func (s *DBClient) CreateFaculty(fname string, lname string, email string) (int, error) {

	var fac str.Faculty
	fac.FirstName = fname
	fac.LastName = lname
	fac.Email = email

	err := s.Db.Create(&fac).Error
	if err != nil {
		return -1, fmt.Errorf(err.Error())
	}
	return int(fac.ID), nil
}

func (s *DBClient) CreateCourse(name string) (int, error) {
	var c str.Course
	c.Name = name
	err := s.Db.Create(&c).Error
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}

	return int(c.ID), nil

}

func (s *DBClient) AddFacToCourse(fId int32, cId int) error {
	var c str.Course
	err := s.Db.Where("ID=?", cId).First(&c).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	c.FacultyID = uint(fId)
	err = s.Db.Save(&c).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
