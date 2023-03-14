package serverfuncs

import (
	s "beprj/src/Structs"
	"github.com/jinzhu/gorm"
)

func InsertMockData(db *gorm.DB) error {
	users := []s.Student{
		{FirstName: "Ramana", LastName: "Pullela", Email: "john@example.com"},
		{FirstName: "Rithika1", LastName: "Pullela1", Email: "john1@example.com"},
		{FirstName: "Rithika2", LastName: "Pullela2", Email: "john2@example.com"},
		{FirstName: "Rithika3", LastName: "Pullela3", Email: "john3@example.com"},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			return result.Error
		}
	}

	users2 := []s.Faculty{
		{FirstName: "f1", LastName: "l1", Email: "fac1@example.com"},
		{FirstName: "f2", LastName: "l2", Email: "fac2@example.com"},
	}

	for _, user := range users2 {
		result := db.Create(&user)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
