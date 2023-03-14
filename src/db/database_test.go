package database

import (

	//"reflect"
	"testing"
//	str "beprj/src/Structs"
	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/jinzhu/gorm"

)

func TestAddFacToCourse(t *testing.T) {
	// Create a new mock database connection and a DBClient that uses it
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %v", err)
	}
	defer db.Close()

	// client := &DBClient{Db: *gorm.Open("postgres")}
	
	//client := &DBClient{Db: &gorm.DB{}}
	// Set up mock database expectations
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT \* FROM "courses" WHERE ID=\?`).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "faculty_id"}).AddRow(1, "Test Course", 2))
	mock.ExpectExec(`UPDATE "courses" SET "faculty_id"=\? WHERE "id" = \?`).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// Call the method being tested
	// if err = AddFacToCourse( 1, 1); err != nil {
	// 	t.Errorf("error was not expected while updating stats: %s", err)
	// }
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that the expected SQL queries were executed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unfulfilled expectations: %v", err)
	}

	
}

func AddFacToCourse(i1, i2 int) {
	panic("unimplemented")
}

