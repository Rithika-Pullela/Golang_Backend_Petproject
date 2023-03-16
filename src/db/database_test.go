package database

import (
	"testing"
	_ "time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func newMockDBClient() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		panic("Got an unexpected error while opening mock sql connection.")
	}
	return gormDB, mock
}

func TestCreateUser(t *testing.T) {
	// Create a new mock database connection and a DBClient that uses it
	db, mock := newMockDBClient()
	defer db.Close()
	defer mock.ExpectClose()

	DBClient := DBClient{
		Db: db,
	}

	
	// t.Run("Adding student", func(t *testing.T) {
		//mock.ExpectBegin()
	// 	mock.ExpectQuery(`INSERT INTO "students"(.+)`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	// 	mock.ExpectCommit()

	// 	id, err := DBClient.CreateStudent("s1", "l1", "@.com")
	// 	assert.Nil(t, err)

	// 	if id != 1 {
	// 		t.Errorf("Unable to add student %v", err)
	// 	}

	// 	// we make sure that all expectations were met
	// 	if err := mock.ExpectationsWereMet(); err != nil {
	// 		t.Errorf("there were unfulfilled expectations: %s", err)
	// 	}

	// })

	// t.Run("Adding faculty", func(t *testing.T) {
	//	mock.ExpectBegin()
	// 	mock.ExpectQuery(`INSERT INTO "faculties"(.+)`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	// 	mock.ExpectCommit()

	// 	id,err:=DBClient.CreateFaculty("f1","l1","f1@.com")
	// 	assert.Nil(t,err)

	// 	if id!=1{
	// 		t.Errorf("Unable to add faculty %v",err)
	// 	}

	// 	if err := mock.ExpectationsWereMet(); err != nil {
	// 		t.Errorf("there were unfulfilled expectations: %s", err)
	// 	}

	// })

	t.Run("Adding Course", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "courses"(.+)`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		id,_,err:=DBClient.CreateCourse("oops","a")
		assert.Nil(t,err)

		if id!=1{
			t.Errorf("Unable to add course %v",err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("Adding Faculty to course", func(t *testing.T) {
		result := sqlmock.NewResult(1, 1)
	
	
		mock.ExpectQuery(`SELECT (.+) FROM "courses".*$`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "faculty_id"}).AddRow(1, "aka", 1))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE (.+) FROM "courses".*$`).WithArgs(1,1).WillReturnResult(result)
		 mock.ExpectCommit()

		err := DBClient.AddFacToCourse(1, 1)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		assert.Nil(t, err)

	})


}
