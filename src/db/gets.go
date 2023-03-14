package database

import (
	str "beprj/src/Structs"
	"fmt"
	_ "github.com/lib/pq"
)

func (s *DBClient) GetAssignmentsByCourseId(cId int32) ([]str.Assignment, error) {
	var c str.Course
	err := s.Db.Preload("Assignments").Where("ID=?", cId).Find(&c).Error
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return c.Assignments, nil
}

func (s *DBClient) GetStudentsEnrolledByCourseId(cId int32) ([]str.Student, error) {
	var c str.Course
	err := s.Db.Preload("Students").Where("ID=?", cId).Find(&c).Error
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return c.Students, nil
}

func (s *DBClient) GetFacultyByCourseId(cId int32) (str.Faculty, error) {

	var c str.Course
	err := s.Db.Preload("Faculty").Where("id=?", cId).Find(&c).Error
	if err != nil {
		return str.Faculty{}, fmt.Errorf(err.Error())
	}
	res := str.Faculty{
		FirstName: c.Faculty.FirstName,
		LastName:  c.Faculty.LastName,
		Email:     c.Faculty.Email,
	}
	return res, nil
}

func (s *DBClient) GetCoursesByFacultyEmail(email string) ([]str.Course, error) {
	var fac str.Faculty
	err := s.Db.Preload("Courses").Where("Email=?", email).Find(&fac).Error
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return fac.Courses, nil
}

func (s *DBClient) GetCoursesEnrolledByStudentEmail(email string) ([]str.Course, error) {
	var std str.Student
	err := s.Db.Preload("Courses").Preload("Courses.Faculty").Where("Email=?", email).Find(&std).Error
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	
	return std.Courses, nil
}

func (s *DBClient) GetAllCourses() ([]str.Course, error) {
	var courses []str.Course
	err := s.Db.Preload("Faculty").Preload("Students").Find(&courses).Error
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return courses, nil
}

func (s *DBClient) GetStudent(email string) (str.Student, error) {
	var std str.Student
	err := s.Db.Model(str.Student{}).Preload("Courses").Where("Email=?", email).First(&std).Error
	x := str.Student{
		FirstName: std.FirstName,
		LastName:  std.LastName,
		Email:     std.Email,
	}
	if err != nil {
		return str.Student{}, fmt.Errorf(err.Error())
	}
	return x, nil
}

func (s *DBClient) GetFaculty(email string) (str.Faculty, error) {
	var fac str.Faculty
	err := s.Db.Model(str.Faculty{}).Preload("Courses").Where("Email=?", email).First(&fac).Error
	x := str.Faculty{
		FirstName: fac.FirstName,
		LastName:  fac.LastName,
		Email:     fac.Email,
	}
	if err != nil {
		return str.Faculty{}, fmt.Errorf(err.Error())
	}
	return x, nil
}

func (s *DBClient) GetCourseByCourseId(cId int32) (str.Course, error) {
	c := str.Course{}
	err := s.Db.Model(str.Course{}).Preload("Faculty").Preload("Assignments").Where("ID=?", cId).First(&c).Error
	x := str.Course{
		Name:      c.Name,
		FacultyID: c.FacultyID,
		Faculty:   c.Faculty,
	}
	if err != nil {
		return str.Course{}, fmt.Errorf(err.Error())
	}

	return x, nil
}