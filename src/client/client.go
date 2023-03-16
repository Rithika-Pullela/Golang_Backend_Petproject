package main

import (
	pb "beprj/src/proto"
	"context"
	"fmt"
	_ "math/rand"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	//"log"
)

func main() {
	fmt.Println("client application here!")

	conn, err := grpc.Dial("localhost:7005", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic("did not connect ")
	}
	defer conn.Close()

	c := pb.NewUniversityFuncClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("Connection Successfull")
	fmt.Println()

	cou := &pb.CourseDetails{
		Name:      "cooooppsss",
		FacultyID: 1,
		Image: "https://cdn.educba.com/academy/wp-content/uploads/2018/12/Computer-Architecture.jpg",
	}
	res, err := c.AddCourse(ctx, cou)
	fmt.Println("AddCourse ", res,err)
	fmt.Println()

	// cou2 := &pb.CourseDetails{
	// 	FacultyID:   1,
	// 	Name: "test_course",
	// }
	// res1, err := c.AddCourse(ctx, cou2)
	// fmt.Println("AddCoursetestttt", res1,err)
	// fmt.Println()


	// cou3 := &pb.CourseDetails{
	// 	Name:      "",
	// 	FacultyID: 2,
	// }
	// res3, err := c.AddCourse(ctx, cou3)
	// fmt.Println("AddCourse", res3,err)
	// fmt.Println()
	std := &pb.User{
		FirstName: "John",
		LastName: "ammm",
		Email:     "john.doe@example.com",
	}
	stdres, err := c.AddStudent(ctx, std)
	fmt.Println("AddStudent", stdres,err)
	fmt.Println()

	// // fac := &pb.User{
	// // 	FirstName: "facccc",
	// // 	LastName:  "snrrrr",
	// // 	Email:     "facacc@.com",
	// // }
	// // facres, err := c.AddFaculty(ctx, fac)
	// // fmt.Println("AddFac", facres)
	// // fmt.Println()

	// enroll := &pb.EnrollRequest{
	// 	CourseID: 2,
	// }
	// enrollres, err := c.AddEnrollment(ctx, enroll)
	// fmt.Println("AddEnrollment", enrollres,err)
	// fmt.Println()

	//  a := &pb.Assignment{
	// 			Title:         "Assignment 3",
	// 			Description:   "Description for assignment 3",
	// 			Deadline:      "2023-04-02",
	// 			CourseId:      999,
	// 		}
	// response, err := c.AddAssignment(ctx, a)
	// fmt.Println("AddAssignment", response,err)
	// fmt.Println()

	// course := &pb.CourseDetails{
	// 	ID: 100,
	// }
	// resp, err := c.GetCourseByCourseId(ctx, course)
	// fmt.Println("GetCourseByCourseId", resp,err)
	// fmt.Println()

	// student := &pb.User{
	// 	ID: 1,
	// }
	// r, err := c.GetStudentById(ctx, student)
	// fmt.Println("GetStudentById", r)
	// fmt.Println()

	// faculty := &pb.User{
	// 	ID: 2,
	// }
	// r1, err := c.GetFacultyById(ctx, faculty)
	// fmt.Println("GetFacultyById", r1)
	// fmt.Println()

	empty := &pb.Empty{}
	bigres, err := c.GetAllCourses(ctx, empty)
	fmt.Println("GetAllCourses", bigres)
	fmt.Println()

	input := &pb.User{
		Email:"john1@example.com",
	}
	output,err :=c.GetCoursesEnrolledByStudentEmail(ctx,input)
	fmt.Println("Gettttttttttttttt",output,err)
	fmt.Println()
	// xx := &pb.User{
	// 	ID: 2,
	// }
	// ans, err := c.GetCoursesEnrolledByStudentId(ctx, xx)
	// fmt.Println("GetCoursesEnrolledByStudentId", ans)
	// fmt.Println()

	// params := &pb.User{
	// 	ID: 2,
	// }
	// sol, err := c.GetCoursesByFacultyId(ctx, params)
	// fmt.Println("GetCoursesByFacultyId", sol)
	// fmt.Println()

	// params1 := &pb.CourseDetails{
	// 	ID: 2,
	// }
	// sol1, err := c.GetFacultyByCourseId(ctx, params1)
	// fmt.Println("GetFacultyByCourseId", sol1)
	// fmt.Println()

	// params2 := &pb.CourseDetails{
	// 	ID: 1,
	// }
	// sol2, err := c.GetStudentsEnrolledByCourseId(ctx, params2)
	// fmt.Println("GetStudentsEnrolledByCourseId", sol2)
	// fmt.Println()

	// params3 := &pb.CourseDetails{
	// 	ID: 1,
	// }
	// sol3, err := c.GetAssignmentsByCourseId(ctx, params3)
	// fmt.Println("GetAssignmentsByCourseId", sol3)
	// fmt.Println()

	//  params4 := &pb.CourseDetails{
	// 	ID:        1,
	// 	Name:      "",
	// 	FacultyID: 0,
	// }
	// sol4,err := c.UpdateCourse(ctx,params4)
	// fmt.Println("UpdateCourse", sol4,err)
	// fmt.Println()

	// params5 := &pb.Assignment{
	// 	ID: 1,
	// 	Title: "UpdatedAssignmenttitle",
	// 	Description: "updateddes",
	// 	Deadline: "01/01/2023",
	// }
	// sol5,err := c.UpdateAssignment(ctx,params5)
	// fmt.Println("UpdateAssignment", sol5,err)
	// fmt.Println()
}
