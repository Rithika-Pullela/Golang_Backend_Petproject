package serverfuncs

import (
	s "beprj/src/Structs"
	"github.com/jinzhu/gorm"
)

func InsertMockData(db *gorm.DB) error {
	users := []s.Student{
		{FirstName: "Ramana", LastName: "Pullela", Email: "john@example.com",Password:"123"},
		{FirstName: "Rithika1", LastName: "Pullela1", Email: "john1@example.com",Password:"123"},
		{FirstName: "Rithika2", LastName: "Pullela2", Email: "john2@example.com",Password:"123"},
		{FirstName: "Rithika3", LastName: "Pullela3", Email: "john3@example.com",Password:"123"},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			return result.Error
		}
	}

	users2 := []s.Faculty{
		{FirstName: "f1", LastName: "l1", Email: "fac1@example.com",Password:"123"},
		{FirstName: "f2", LastName: "l2", Email: "fac2@example.com",Password:"123"},
		{FirstName: "f3", LastName: "l3", Email: "fac3@example.com",Password:"123"},
	}

	for _, user := range users2 {
		result := db.Create(&user)
		if result.Error != nil {
			return result.Error
		}
	}

	courses :=[]s.Course{
		{
			Name: "c1",
			FacultyID: 1,
			Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQbk6-cfNgAGDMOQpWDx_m0yWbDC0xXqY201Q&usqp=CAU",
		},
		{
			Name: "c2",
			FacultyID: 1,
			Image: "https://studyopedia.com/wp-content/uploads/2017/04/computer-networking.png",
		},
		{
			Name: "c3",
			FacultyID: 2,
			Image:"https://cdn.educba.com/academy/wp-content/uploads/2018/12/Computer-Architecture.jpg" ,
		},
		{
			Name: "c4",
			FacultyID: 2,
			Image:"https://cdn.educba.com/academy/wp-content/uploads/2018/12/Computer-Architecture.jpg" ,
		},
		{
			Name: "newwwwww",
			FacultyID: 2,
			Image:"https://cdn.educba.com/academy/wp-content/uploads/2018/12/Computer-Architecture.jpg" ,
		},
	}

	for _, course := range courses  {
		result := db.Create(&course)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
