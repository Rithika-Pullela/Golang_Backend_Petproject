package structs

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Course struct {
	gorm.Model
	Name        string `gorm:"unique"`
	FacultyID   uint   // foreign key to Faculty model
	Faculty     Faculty
	Assignments []Assignment // has many relationship with Assignment model
	Students    []Student    `gorm:"many2many:student_courses"` // many-to-many relationship with Student model
	Image       string
}

type Faculty struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string
	Email     string   `gorm:"unique;not null"`
	Courses   []Course // has many relationship with Course model
	Password  string
}

type Student struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string
	Email     string   `gorm:"unique;not null"`
	Courses   []Course `gorm:"many2many:student_courses"` // many-to-many relationship with Course model
	Password  string
}

type Assignment struct {
	gorm.Model
	Title       string
	Description string
	Deadline    time.Time
	CourseID    uint `gorm:"notnull"` // foreign key to Course model
}
