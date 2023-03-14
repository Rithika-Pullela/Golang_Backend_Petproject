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

func TestUpdateCourse(t *testing.T) {
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
				ID:        1,
				Name:      "Oops",
				FacultyID: 2,
			},
			mockFunc: func() {
				mockDb.EXPECT().UpdateCourse(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					str.Course{
						Model:     gorm.Model{ID: 1},
						Name:      "Oops",
						FacultyID: 2,
					}, nil)
			},
			expectedOutput: &pb.CourseDetails{
				ID:        1,
				Name:      "Oops",
				FacultyID: 2,
			},
			expectedError: nil,
		},
		{
			name: "Empty input",
			input: &pb.CourseDetails{
				ID:        1,
				Name:      "",
				FacultyID: 0,
			},
			mockFunc: func() {

			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Enter valid input to properly update the course"),
		},
		{
			name: "Invalid Faculty ID",
			input: &pb.CourseDetails{
				ID:        1,
				Name:      "Oops",
				FacultyID: 0,
			},
			mockFunc: func() {

			},

			expectedOutput: nil,
			expectedError:  fmt.Errorf("Enter valid input to properly update the course"),
		},
		{
			name: "Invalid Course ID",
			input: &pb.CourseDetails{
				ID:        100,
				Name:      "Oops",
				FacultyID: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().UpdateCourse(gomock.Any(), gomock.Any(), gomock.Any()).Return(str.Course{}, fmt.Errorf("Enter valid course id"))
			},

			expectedOutput: nil,
			expectedError:  fmt.Errorf("Enter valid course id"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.UpdateCourse(ctx, test.input)

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

func TestUpdateAssignment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	myServer := &UniversityServer{Db: mockDb}
	ctx := context.Background()

	layout := "01/02/2006"
	dateString := "11/11/2023"
	ds, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	tests := []struct {
		name           string
		input          *pb.Assignment
		mockFunc       func()
		expectedOutput *pb.Assignment
		expectedError  error
	}{
		{
			name: "Valid input",
			input: &pb.Assignment{
				ID:          1,
				Title:       "UpdatedAssignmenttitle",
				Description: "updateddes",
				Deadline:    "11/11/2023",
			},
			mockFunc: func() {
				mockDb.EXPECT().UpdateAssignment(gomock.Any(), gomock.Any()).Return(
					str.Assignment{
						Model:       gorm.Model{ID: 1},
						Title:       "UpdatedAssignmenttitle",
						Description: "updateddes",
						Deadline:    ds,
					}, nil)
			},
			expectedOutput: &pb.Assignment{
				ID:          1,
				Title:       "UpdatedAssignmenttitle",
				Description: "updateddes",
				Deadline:    "11/11/2023",
			},
			expectedError: nil,
		},
		{
			name: "Empty input",
			input: &pb.Assignment{
				ID:          1,
				Title:       "",
				Description: "",
				Deadline:    "",
			},
			mockFunc: func() {

			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Enter valid input to properly update the assignment"),
		},
		{
			name: "Missing Title",
			input: &pb.Assignment{
				ID:          1,
				Title:       "",
				Description: "updateddes",
				Deadline:    "11/11/2023",
			},
			mockFunc: func() {

			},

			expectedOutput: nil,
			expectedError:  fmt.Errorf("Enter valid input to properly update the assignment"),
		},
		{
			name: "Missing Description",
			input: &pb.Assignment{
				ID:          1,
				Title:       "UpdatedAssignmenttitle",
				Description: "",
				Deadline:    "11/11/2023",
			},
			mockFunc: func() {

			},

			expectedOutput: nil,
			expectedError:  fmt.Errorf("Enter valid input to properly update the assignment"),
		},
		{
			name: "Missing Deadline",
			input: &pb.Assignment{
				ID:          1,
				Title:       "UpdatedAssignmenttitle",
				Description: "updateddes",
				Deadline:    "",
			},
			mockFunc: func() {

			},

			expectedOutput: nil,
			expectedError:  fmt.Errorf("Enter valid input to properly update the assignment"),
		},
		{
			name: "Out_of_time_Deadline",
			input: &pb.Assignment{
				ID:          1,
				Title:       "UpdatedAssignmenttitle",
				Description: "updateddes",
				Deadline:    "01/01/2023",
			},
			mockFunc: func() {
				mockDb.EXPECT().UpdateAssignment(gomock.Any(), gomock.All()).Return(str.Assignment{}, fmt.Errorf("the assignment deadline is in the past"))
			},

			expectedOutput: nil,
			expectedError:  fmt.Errorf("the assignment deadline is in the past"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mock functions
			test.mockFunc()

			// call function being tested
			output, err := myServer.UpdateAssignment(ctx, test.input)

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
