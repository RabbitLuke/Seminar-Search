// queryUsers.go
package query

import (
	"fmt"
	"github.com/RabbitLuke/seminar-search/dbSetup"
)

// You also need to define a User struct at the beginning of the queryUsers.go file
type Faculty struct {
	FacultyID int    `json:"facultyID"`
	Name      string `json:"name"`
}

// InsertUser inserts a new user into the Faculty table
func InsertFaculty(name string) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	// Prepare the SQL statement
	stmt, err := dbSetup.DB.Prepare("INSERT INTO Faculty (Name) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the provided values
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}

	fmt.Println("Faculty inserted successfully!")
	return nil
}

func DeleteFaculty(facultyID int) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	// Prepare the SQL statement
	stmt, err := dbSetup.DB.Prepare("DELETE FROM Faculty WHERE facultyID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the provided facultyID
	_, err = stmt.Exec(facultyID)
	if err != nil {
		return err
	}

	fmt.Printf("Faculty with facultyID %d deleted successfully!\n", facultyID)
	return nil
}

func UpdateFaculty(facultyID int, name string) error {
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("UPDATE Faculty SET Name = ? WHERE facultyID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, facultyID)
	if err != nil {
		return err
	}

	fmt.Printf("Faculty with facultyID %d updated successfully!\n", facultyID)
	return nil
}

func SelectFaculties() ([]Faculty, error) {
	if dbSetup.DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	rows, err := dbSetup.DB.Query("SELECT * FROM Faculty")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Faculty

	for rows.Next() {
		var user Faculty
		if err := rows.Scan(&user.FacultyID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func SelectFacultyByID(facultyID int) (*Faculty, error) {
	if dbSetup.DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	var user Faculty
	err := dbSetup.DB.QueryRow("SELECT * FROM Faculty WHERE facultyID = ?", facultyID).Scan(&user.FacultyID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
