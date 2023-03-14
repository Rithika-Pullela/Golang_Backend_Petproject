package serverfuncs

import (
	pb "beprj/src/proto"
	"context"
	_ "github.com/lib/pq"
	
)

func (u *UniversityServer) GetCourseByCourseId(ctx context.Context, in *pb.CourseDetails) (*pb.CourseDetails, error) {
	cid := in.GetID()
	res, err := u.Db.GetCourseByCourseId(cid)
	if err != nil {
		return nil, err
	}
	result := &pb.CourseDetails{
		ID:        cid,
		Name:      res.Name,
		FacultyID: int32(res.FacultyID),
		Faculty: &pb.User{
			FirstName: res.Faculty.FirstName,
			LastName:  res.Faculty.LastName,
			Email:     res.Faculty.Email,
		},
	}
	return result, nil
}

func (u *UniversityServer) GetStudentByEmail(ctx context.Context, in *pb.User) (*pb.User, error) {
	mail := in.GetEmail()
	res, err := u.Db.GetStudent(mail)
	if err != nil {
		return nil, err
	}
	result := &pb.User{
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
	}
	return result, nil
}

func (u *UniversityServer) GetFacultyByEmail(ctx context.Context, in *pb.User) (*pb.User, error) {
	email := in.GetEmail()
	res, err := u.Db.GetFaculty(email)
	if err != nil {
		return nil, err
	}
	result := &pb.User{
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
	}
	return result, nil
}

func (u *UniversityServer) GetAllCourses(ctx context.Context, in *pb.Empty) (*pb.CourseList, error) {
	res, err := u.Db.GetAllCourses()
	if err != nil {
		return nil, err
	}
	result := &pb.CourseList{}
	for _, course := range res {
		x := &pb.CourseDetails{
			FacultyID: int32(course.FacultyID),
			Name:      course.Name,
			ID:        int32(course.ID),
			Faculty: &pb.User{
				ID:        int32(course.Faculty.ID),
				FirstName: course.Faculty.FirstName,
				LastName:  course.Faculty.LastName,
				Email:     course.Faculty.Email,
			},
		}
		result.Courses = append(result.Courses, x)

	}
	return result, nil
}

func (u *UniversityServer) GetCoursesEnrolledByStudentEmail(ctx context.Context, in *pb.User) (*pb.CourseList, error) {
	email := in.GetEmail()
	res, err := u.Db.GetCoursesEnrolledByStudentEmail(email)
	if err != nil {
		return nil, err
	}
	result := &pb.CourseList{}
	for _, course := range res {
		x := &pb.CourseDetails{
			FacultyID: int32(course.FacultyID),
			Name:      course.Name,
			ID:        int32(course.ID),
			Faculty: &pb.User{
				ID:        int32(course.Faculty.ID),
				FirstName: course.Faculty.FirstName,
				LastName:  course.Faculty.LastName,
				Email:     course.Faculty.Email,
			},
		}
		result.Courses = append(result.Courses, x)
	}
	return result, nil
}

func (u *UniversityServer) GetCoursesByFacultyEmail(ctx context.Context, in *pb.User) (*pb.CourseList, error) {
	email := in.GetEmail()
	res, err := u.Db.GetCoursesByFacultyEmail(email)
	if err != nil {
		return nil, err
	}
	result := &pb.CourseList{}
	for _, course := range res {
		x := &pb.CourseDetails{
			FacultyID: int32(course.FacultyID),
			Name:      course.Name,
			ID:        int32(course.ID),
			
		}
		result.Courses = append(result.Courses, x)
	}
	return result, nil
}

func (u *UniversityServer) GetFacultyByCourseId(ctx context.Context, in *pb.CourseDetails) (*pb.User, error) {
	cId := in.GetID()
	res, err := u.Db.GetFacultyByCourseId(cId)
	if err != nil {
		return nil, err
	}
	ans := &pb.User{
		ID: int32(res.ID),
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
	}
	return ans, nil
}

func (u *UniversityServer) GetStudentsEnrolledByCourseId(ctx context.Context, in *pb.CourseDetails) (*pb.UserList, error) {
	cId := in.GetID()
	res, err := u.Db.GetStudentsEnrolledByCourseId(cId)
	if err != nil {
		return nil, err
	}
	response := &pb.UserList{}
	for _, student := range res {
		x := &pb.User{
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		}
		response.Users = append(response.Users, x)
	}
	return response, nil
}

func (u *UniversityServer) GetAssignmentsByCourseId(ctx context.Context, in *pb.CourseDetails) (*pb.AssignmentList, error) {

	cId := in.GetID()
	res, err := u.Db.GetAssignmentsByCourseId(cId)
	if err != nil {
		return nil, err
	}
	response := &pb.AssignmentList{}
	for _, assignment := range res {
		x := &pb.Assignment{
			Title:       assignment.Title,
			Description: assignment.Description,
			Deadline:    assignment.Deadline.Format("01/02/2006"),
		}
		response.Assignments = append(response.Assignments, x)
	}
	return response, nil
}
