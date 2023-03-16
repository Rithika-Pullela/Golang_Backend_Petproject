package serverfuncs

import (
	dataBase "beprj/src/db"
	pb "beprj/src/proto"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type UniversityServer struct {
	pb.UnimplementedUniversityFuncServer
	Db dataBase.Database
	DB *gorm.DB
}

func (u *UniversityServer) AddCourse(ctx context.Context, in *pb.CourseDetails) (*pb.CourseDetails, error) {
	// cname ledu = error
	// fid==0 create course
	//fid invalid---dont create course
	// fid!=0 create course and add fac to course
	
	var fid = in.GetFacultyID()
	var c_name = in.GetName()
	var c_image = in.GetImage()
	if c_name == ""||c_image==""{
		return nil, fmt.Errorf("Course fields cannot be empty!")
	}

	if fid!=0{

		fac_exists, fac_err := u.Db.CheckFacultyExistance(fid)
		if !fac_exists {
			return nil, fac_err
		}
	}
	id,image, err := u.Db.CreateCourse(c_name,c_image)
	if err != nil {
		return nil, err
	}
	if fid!=0{

		err = u.Db.AddFacToCourse(fid, id)
		if err != nil {
			return nil, err
		}
	}
	
	
	res := &pb.CourseDetails{
		ID:   int32(id),
		Name: c_name,
		Image: image,
	}
	
	res.FacultyID = fid
	return res, nil
}

func (u *UniversityServer) AddStudent(ctx context.Context, in *pb.User) (*pb.User, error) {
	fname := in.GetFirstName()
	lname := in.GetLastName()
	email := in.GetEmail()
	if fname == "" || lname == "" || email == "" {
		return nil, fmt.Errorf("Provide all the details!")
	}
	id, err := u.Db.CreateStudent(fname, lname, email)
	if err != nil {
		return nil, err
	}
	res := &pb.User{
		ID:        int32(id),
		FirstName: fname,
		LastName:  lname,
		Email:     email,
	}
	return res, nil
}

func (u *UniversityServer) AddFaculty(ctx context.Context, in *pb.User) (*pb.User, error) {

	fname := in.GetFirstName()
	lname := in.GetLastName()
	email := in.GetEmail()
	if fname == "" || lname == "" || email == "" {
		return nil, fmt.Errorf("Provide all the details!")
	}
	id, err := u.Db.CreateFaculty(fname, lname, email)
	if err != nil {
		return nil, err
	}
	res := &pb.User{
		ID:        int32(id),
		FirstName: fname,
		LastName:  lname,
		Email:     email,
	}
	return res, nil
}

func (u *UniversityServer) AddEnrollment(ctx context.Context, in *pb.EnrollRequest) (*pb.Status, error) {

	sid := in.GetStudentID()
	cid := in.GetCourseID()
	if sid == 0 || cid == 0 {
		return nil, fmt.Errorf("Mention student id and and course id to enroll student in a course")
	}
	res := &pb.Status{
		Des: "Enrollment Failed",
	}
	err := u.Db.AddCoursetoStd(int32(sid), int(cid))
	if err != nil {
		res.Error = err.Error()
		return res, nil
	}
	suc := &pb.Status{
		Des: "Enrollment Successss",
	}
	return suc, nil
}

func (u *UniversityServer) AddAssignment(ctx context.Context, in *pb.Assignment) (*pb.Assignment, error) {
	_, err2 := u.Db.CheckCourseExistance(int32(in.GetCourseId()))
	if err2 != nil {
		return nil, fmt.Errorf("Course Doesn't Exists!")
	}
	title := in.GetTitle()
	des := in.GetDescription()
	deadline := in.GetDeadline()
	cid := in.GetCourseId()
	if title == "" {
		return nil, fmt.Errorf("Enter Assignment title!!")
	}
	id, err := u.Db.CreateAssignment(title, des, deadline, cid)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	res := &pb.Assignment{
		ID:          int32(id),
		Title:       title,
		Description: des,
		Deadline:    deadline,
		CourseId:    cid,
	}
	return res, nil

}
