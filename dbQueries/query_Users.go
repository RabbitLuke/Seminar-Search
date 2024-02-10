// queryUsers.go
package queryUsers

import (
	"fmt"
	"github.com/RabbitLuke/seminar-search/dbSetup"
)

// InsertUser inserts a new user into the Faculty table
func InsertUser(name string) error {
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

	fmt.Println("User inserted successfully!")
	return nil
}
