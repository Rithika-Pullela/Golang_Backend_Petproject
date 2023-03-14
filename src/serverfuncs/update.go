package serverfuncs

import (
	str "beprj/src/Structs"
	pb "beprj/src/proto"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func (u *UniversityServer) UpdateCourse(ctx context.Context, in *pb.CourseDetails) (*pb.CourseDetails, error) {

	if in.GetFacultyID() == 0 || in.GetName() == "" {
		return nil, fmt.Errorf("Enter valid input to properly update the course")
	}
	fid := in.GetFacultyID()
	c_name := in.GetName()
	cid := in.GetID()
	res, err := u.Db.UpdateCourse(fid, c_name, cid)
	if err != nil {
		return nil, err
	}
	return &pb.CourseDetails{
		ID:        int32(res.ID),
		Name:      res.Name,
		FacultyID: int32(res.FacultyID),
	}, nil
}

func (u *UniversityServer) UpdateAssignment(ctx context.Context, in *pb.Assignment) (*pb.Assignment, error) {
	aid := in.GetID()
	deadline := in.GetDeadline()
	des := in.GetDescription()
	title := in.GetTitle()
	if deadline == "" || des == "" || title == "" {
		return nil, fmt.Errorf("Enter valid input to properly update the assignment")
	}
	layout := "01/02/2006"
	s, err := time.Parse(layout, deadline)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, fmt.Errorf(err.Error())
	}
	x := str.Assignment{
		Title:       title,
		Deadline:    s,
		Description: des,
	}
	res, err := u.Db.UpdateAssignment(x, aid)
	if err != nil {
		return nil, err
	}
	return &pb.Assignment{
		ID:          int32(res.ID),
		Title:       res.Title,
		Description: res.Description,
		Deadline:    res.Deadline.Format("01/02/2006"),
	}, nil
}
