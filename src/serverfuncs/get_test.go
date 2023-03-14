package serverfuncs

import (
	"beprj/src/Mocks"
	str "beprj/src/Structs"
	pb "beprj/src/proto"
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
)

func TestGetCourseByCourseId(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *pb.CourseDetails
		mockFunc       func()
		expectedOutput *pb.CourseDetails
		expectedError  error
	}{
		{
			name: "Valid input",
			input: &pb.CourseDetails{
				ID: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetCourseByCourseId(gomock.Any()).Return(str.Course{
					Name:      "c1",
					FacultyID: 2,
					Faculty: str.Faculty{
						FirstName: "s",
						LastName:  "b",
						Email:     "kaj@gmail.com",
					},
				}, nil)
			},
			expectedOutput: &pb.CourseDetails{
				ID:        1,
				Name:      "c1",
				FacultyID: 2,
				Faculty: &pb.User{
					FirstName: "s",
					LastName:  "b",
					Email:     "kaj@gmail.com",
				},
			},

			expectedError: nil,
		},
		{
			name: "Missing course ID",
			input: &pb.CourseDetails{
				ID: 2,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetCourseByCourseId(gomock.Any()).Return(str.Course{}, fmt.Errorf("record not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("record not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.GetCourseByCourseId(ctx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}

}

////////////////////

func TestGetStudentByEmail(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *pb.User
		mockFunc       func()
		expectedOutput *pb.User
		expectedError  error
	}{
		{
			name: "valid_email",
			input: &pb.User{
				Email: "a@.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetStudent(gomock.Any()).Return(str.Student{
					FirstName: "f",
					LastName:  "l",
					Email:     "a@.com",
				}, nil)
			},
			expectedOutput: &pb.User{
				FirstName: "f",
				LastName:  "l",
				Email:     "a@.com",
			},
			expectedError: nil,
		},
		{
			name: "invalid_email",
			input: &pb.User{
				Email: "a@.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetStudent(gomock.Any()).Return(str.Student{}, fmt.Errorf("student not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("student not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.GetStudentByEmail(ctx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}

}

//////

func TestGetFacultyByEmail(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *pb.User
		mockFunc       func()
		expectedOutput *pb.User
		expectedError  error
	}{
		{
			name: "valid_email",
			input: &pb.User{
				Email: "a@.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetFaculty(gomock.Any()).Return(str.Faculty{
					FirstName: "f",
					LastName:  "l",
					Email:     "a@.com",
				}, nil)
			},
			expectedOutput: &pb.User{
				FirstName: "f",
				LastName:  "l",
				Email:     "a@.com",
			},
			expectedError: nil,
		},
		{
			name: "invalid_email",
			input: &pb.User{
				Email: "a@.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetFaculty(gomock.Any()).Return(str.Faculty{}, fmt.Errorf("Fac not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Fac not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.GetFacultyByEmail(ctx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}

}

func TestGetAllCourses(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	input := &pb.Empty{}
	course := str.Course{
		Name:      "c1",
		FacultyID: 1,
	}
	course_array := []str.Course{
		course,
	}
	mockDb.EXPECT().GetAllCourses().Return(course_array, nil)

	expected := &pb.CourseList{}
	c := &pb.CourseDetails{
		Name:      "c1",
		FacultyID: 1,
		Faculty:   &pb.User{},
	}
	expected.Courses = append(expected.Courses, c)

	output, err := myServer.GetAllCourses(ctx, input)
	if err != nil {
		t.Errorf("GetAllCourses returned an error: %v", err)
	}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", expected, output)
	}
}

func TestGetCoursesEnrolledByStudentEmail(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	result := &pb.CourseList{}

	x := &pb.CourseDetails{
		FacultyID: 2,
		Name:      "c",
		ID:        1,
		Faculty: &pb.User{
			ID:        2,
			FirstName: "f",
			LastName:  "l",
			Email:     "c@gmail.com",
		},
	}
	result.Courses = append(result.Courses, x)

	tests := []struct {
		name           string
		input          *pb.User
		mockFunc       func()
		expectedOutput *pb.CourseList
		expectedError  error
	}{
		{
			name: "valid_email",
			input: &pb.User{
				Email: "a@.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetCoursesEnrolledByStudentEmail(gomock.Any()).Return([]str.Course{
					{
						Model: gorm.Model{
							ID: 1,
						},
						FacultyID: 2,
						Name:      "c",
						Faculty: str.Faculty{
							Model: gorm.Model{
								ID: 2,
							},
							FirstName: "f",
							LastName:  "l",
							Email:     "c@gmail.com",
						},
					},
				}, nil)
			},

			expectedOutput: result,

			expectedError: nil,
		},
		{
			name: "invalid_email",
			input: &pb.User{
				Email: "a@.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetCoursesEnrolledByStudentEmail(gomock.Any()).Return([]str.Course{}, fmt.Errorf("Email not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Email not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.GetCoursesEnrolledByStudentEmail(ctx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}
}

func TestGetCoursesByFacultyEmail(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	result := &pb.CourseList{}

	x := &pb.CourseDetails{
		FacultyID: 2,
		Name:      "c",
		ID:        1,
	}
	result.Courses = append(result.Courses, x)

	tests := []struct {
		name           string
		input          *pb.User
		mockFunc       func()
		expectedOutput *pb.CourseList
		expectedError  error
	}{
		{
			name: "valid_email",
			input: &pb.User{
				Email: "a@.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetCoursesByFacultyEmail(gomock.Any()).Return([]str.Course{
					{
						Model: gorm.Model{
							ID: 1,
						},
						FacultyID: 2,
						Name:      "c",
						Faculty: str.Faculty{
							Model: gorm.Model{
								ID: 2,
							},
							FirstName: "f",
							LastName:  "l",
							Email:     "c@gmail.com",
						},
					},
				}, nil)
			},

			expectedOutput: result,

			expectedError: nil,
		},
		{
			name: "invalid_email",
			input: &pb.User{
				Email: "a@.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetCoursesByFacultyEmail(gomock.Any()).Return([]str.Course{}, fmt.Errorf("Email not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Email not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.GetCoursesByFacultyEmail(ctx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}
}

func TestGetFacultyByCourseId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *pb.CourseDetails
		mockFunc       func()
		expectedOutput *pb.User
		expectedError  error
	}{
		{
			name: "valid_email",
			input: &pb.CourseDetails{
				ID: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetFacultyByCourseId(gomock.Any()).Return(str.Faculty{

					Model: gorm.Model{
						ID: 2,
					},
					FirstName: "f",
					LastName:  "l",
					Email:     "ak@.com",
				}, nil)
			},

			expectedOutput: &pb.User{
				ID:        2,
				FirstName: "f",
				LastName:  "l",
				Email:     "ak@.com",
			},

			expectedError: nil,
		},
		{
			name: "invalid_email",
			input: &pb.CourseDetails{
				ID: 3,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetFacultyByCourseId(gomock.Any()).Return(str.Faculty{}, fmt.Errorf("ID not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("ID not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.GetFacultyByCourseId(ctx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}
}

func TestGetStudentsEnrolledByCourseId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()
	result := &pb.UserList{}

	x := &pb.User{
		FirstName: "f",
		LastName:  "l",
		Email:     "@.com",
	}
	result.Users = append(result.Users, x)

	tests := []struct {
		name           string
		input          *pb.CourseDetails
		mockFunc       func()
		expectedOutput *pb.UserList
		expectedError  error
	}{
		{
			name: "valid_email",
			input: &pb.CourseDetails{
				ID: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetStudentsEnrolledByCourseId(gomock.Any()).Return([]str.Student{

					{Model: gorm.Model{
						ID: 1,
					},
						FirstName: "f",
						LastName:  "l",
						Email:     "@.com",
					},
				}, nil)
			},

			expectedOutput: result,

			expectedError: nil,
		},
		{
			name: "invalid_email",
			input: &pb.CourseDetails{
				ID: 3,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetStudentsEnrolledByCourseId(gomock.Any()).Return([]str.Student{}, fmt.Errorf("ID not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("ID not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.GetStudentsEnrolledByCourseId(ctx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}
}

func TestGetAssignmentsByCourseId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	result := &pb.AssignmentList{}

	x := &pb.Assignment{
		Title:       "Assignment 3",
		Description: "Description for assignment 3",
		Deadline:    "11/11/2023",
	}
	result.Assignments = append(result.Assignments, x)

	layout := "01/02/2006"
	dateString := "11/11/2023"
	ds, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	tests := []struct {
		name           string
		input          *pb.CourseDetails
		mockFunc       func()
		expectedOutput *pb.AssignmentList
		expectedError  error
	}{
		{
			name: "valid_id",
			input: &pb.CourseDetails{
				ID: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetAssignmentsByCourseId(gomock.Any()).Return([]str.Assignment{
					{
						Title:       "Assignment 3",
						Description: "Description for assignment 3",
						Deadline:    ds,
						CourseID:    1,
					},
				}, nil)
			},

			expectedOutput: result,

			expectedError: nil,
		},
		{
			name: "invalid_id",
			input: &pb.CourseDetails{
				ID: 3,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetAssignmentsByCourseId(gomock.Any()).Return([]str.Assignment{}, fmt.Errorf("ID not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("ID not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.GetAssignmentsByCourseId(ctx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}
}
