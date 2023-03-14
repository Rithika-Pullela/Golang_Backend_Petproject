package serverfuncs

import (
	"beprj/src/Mocks"
	pb "beprj/src/proto"
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

//add few more test cases - - - array of test cases(include edge cases and check if we are getting the expected result)
// We must also write unit tests for db functions(very imp)

func TestAddCourse(t *testing.T) {

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
			name: "valid_course_no_faculty",
			input: &pb.CourseDetails{
				Name: "test_course",
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateCourse(gomock.Any()).Return(123, nil)
			},
			expectedOutput: &pb.CourseDetails{
				ID:   123,
				Name: "test_course",
			},
			expectedError: nil,
		},
		{
			name: "valid_course_with_valid_faculty",
			input: &pb.CourseDetails{
				Name:      "test_course",
				FacultyID: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().CheckFacultyExistance(gomock.Any()).Return(true, nil)
				mockDb.EXPECT().CreateCourse(gomock.Any()).Return(123, nil)
				mockDb.EXPECT().AddFacToCourse(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedOutput: &pb.CourseDetails{
				ID:        123,
				Name:      "test_course",
				FacultyID: 1,
			},
			expectedError: nil,
		},
		{
			name: "invalid_course_missing_name",
			input: &pb.CourseDetails{
				FacultyID: 1,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Course name cannot be empty!"),
		},
		{
			name: "missing_faculty",
			input: &pb.CourseDetails{
				Name: "test_course",
				FacultyID: 200,
			},
			mockFunc: func() {
				mockDb.EXPECT().CheckFacultyExistance(gomock.Any()).Return(false, fmt.Errorf("Faculty not found!"))
				
			},
			expectedOutput: nil,
			expectedError: fmt.Errorf("Faculty not found!"),
		},
		{
			name: "Add_Fac_fails",
			input: &pb.CourseDetails{
				Name: "test_course",
				FacultyID: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().CheckFacultyExistance(gomock.Any()).Return(true, nil)
				mockDb.EXPECT().CreateCourse(gomock.Any()).Return(123, nil)
				mockDb.EXPECT().AddFacToCourse(gomock.Any(), gomock.Any()).Return(fmt.Errorf("Faculty not added!"))
			},
			expectedOutput: &pb.CourseDetails{
				ID:        123,
				Name:      "test_course",
				
			},
			expectedError: fmt.Errorf("Faculty not added!"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.AddCourse(ctx, test.input)

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

///////////////////////////////////////////////////////////////////////////////

func TestAddStudent(t *testing.T) {

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
			name: "Valid input",
			input: &pb.User{
				FirstName: "s1",
				LastName:  "l1",
				Email:     "s1@gmail.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateStudent(gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil)
			},
			expectedOutput: &pb.User{
				ID:        1,
				FirstName: "s1",
				LastName:  "l1",
				Email:     "s1@gmail.com",
			},
			expectedError: nil,
		},
		{
			name: "Missing first name",
			input: &pb.User{
				LastName: "l1",
				Email:    "s1@gmail.com",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide all the details!"),
		},
		{
			name: "Missing last name",
			input: &pb.User{
				FirstName: "s1",
				Email:     "s1@gmail.com",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide all the details!"),
		},
		{
			name: "Missing email",
			input: &pb.User{
				FirstName: "s1",
				LastName:  "l1",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide all the details!"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.AddStudent(ctx, test.input)

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

// //////////////////////////////////////////////////////////////////////////
func TestAddFaculty(t *testing.T) {
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
			name: "Valid input",
			input: &pb.User{
				FirstName: "s1",
				LastName:  "l1",
				Email:     "s1@gmail.com",
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateFaculty(gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil)
			},
			expectedOutput: &pb.User{
				ID:        1,
				FirstName: "s1",
				LastName:  "l1",
				Email:     "s1@gmail.com",
			},
			expectedError: nil,
		},
		{
			name: "Missing first name",
			input: &pb.User{
				LastName: "l1",
				Email:    "s1@gmail.com",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide all the details!"),
		},
		{
			name: "Missing last name",
			input: &pb.User{
				FirstName: "s1",
				Email:     "s1@gmail.com",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide all the details!"),
		},
		{
			name: "Missing email",
			input: &pb.User{
				FirstName: "s1",
				LastName:  "l1",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide all the details!"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.AddFaculty(ctx, test.input)

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

func TestAddEnrollment(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *pb.EnrollRequest
		mockFunc       func()
		expectedOutput *pb.Status
		expectedError  error
	}{
		{
			name: "Valid input",
			input: &pb.EnrollRequest{
				StudentID: 1,
				CourseID:  2,
			},
			mockFunc: func() {
				mockDb.EXPECT().AddCoursetoStd(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedOutput: &pb.Status{
				Des: "Enrollment Successss",
			},
			expectedError: nil,
		},
		{
			name: "Missing student ID",
			input: &pb.EnrollRequest{
				CourseID: 2,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Mention student id and and course id to enroll student in a course"),
		},
		{
			name: "Missing course ID",
			input: &pb.EnrollRequest{
				StudentID: 1,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Mention student id and and course id to enroll student in a course"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.AddEnrollment(ctx, test.input)

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

func TestAddAssignment(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *pb.Assignment
		expectedError  error
		mockFunc       func()
		expectedOutput *pb.Assignment
	}{
		{
			name: "Valid Assignment",
			input: &pb.Assignment{
				Title:       "Assignment 1",
				Description: "Description for assignment 1",
				Deadline:    "2023-03-31",
				CourseId:    1,
			},
			mockFunc: func() {
				mockDb.EXPECT().CheckCourseExistance(gomock.Any()).Return(true, nil)
				mockDb.EXPECT().CreateAssignment(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil)
			},
			expectedError: nil,
			expectedOutput: &pb.Assignment{
				ID:          1,
				Title:       "Assignment 1",
				Description: "Description for assignment 1",
				Deadline:    "2023-03-31",
				CourseId:    1,
			},
		},
		{
			name: "Missing Title",
			input: &pb.Assignment{
				Title:       "",
				Description: "Description for assignment 2",
				Deadline:    "2023-04-01",
				CourseId:    2,
			},

			mockFunc: func() {
				mockDb.EXPECT().CheckCourseExistance(gomock.Any()).Return(true, nil)
			},
			expectedError:  fmt.Errorf("Enter Assignment title!!"),
			expectedOutput: nil,
		},
		{
			name: "Invalid Course ID",
			input: &pb.Assignment{
				Title:       "Assignment 3",
				Description: "Description for assignment 3",
				Deadline:    "2023-04-02",
				CourseId:    999,
			},

			mockFunc: func() {
				mockDb.EXPECT().CheckCourseExistance(gomock.Any()).Return(false, fmt.Errorf("Course Doesn't Exists"))
			},
			expectedError:  fmt.Errorf("Course Doesn't Exists!"),
			expectedOutput: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.AddAssignment(ctx, test.input)

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
