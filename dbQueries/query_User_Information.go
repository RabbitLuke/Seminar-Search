package query

import (
	"fmt"
	"github.com/RabbitLuke/seminar-search/dbSetup"
)

type UserInfo struct{
	UserId      int `json:"UserID"`
	F_name      string `json:"First_Name"`
	L_name      string `json:"Last_Name"`
	Faculty      int `json:"Faculty"`
	Email      string `json:"Email"`
	Profile_pic      string `json:"Profile_pic"`
	Password      string `json:"Password"`
}

func InsertUser(f_name string, l_name string, faculty int, email string, profile_pic string, password string) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	// Prepare the SQL statement
	stmt, err := dbSetup.DB.Prepare("INSERT INTO user_information (First_Name, Last_Name, Faculty, Email, Profile_pic, Password) VALUES (?, ?, ?, ?, ?, ?)")
	fmt.Println(f_name, l_name, email)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the provided values
	_, err = stmt.Exec(f_name, l_name, faculty, email, profile_pic, password)
	if err != nil {
		return err
	}

	fmt.Println("User successfully created!")
	return nil
}

func DeleteUser(UserID int) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("DELETE FROM user_information WHERE UserID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(UserID)
	if err != nil {
		return err
	}

	fmt.Printf("User with ID %d deleted successfully!\n", UserID)
	return nil
}

func UpdateUser(userID int, f_name string, l_name string, faculty int, email string, profile_pic string, password string) error {
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("UPDATE user_information SET First_Name = ?, Last_Name = ?, Faculty = ?, Email = ?, Profile_pic = ?, Password = ? WHERE UserID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, f_name, l_name, faculty, email, profile_pic, password)
	if err != nil {
		return err
	}

	fmt.Printf("User with ID %d updated successfully!\n", userID)
	return nil
}

func SelectUserByID(userID int) (*UserInfo, error) {
	if dbSetup.DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	var user UserInfo
	err := dbSetup.DB.QueryRow("SELECT * FROM user_information WHERE UserID = ?", userID).Scan(&user.UserId, &user.F_name, &user.L_name, &user.Faculty, &user.Email, &user.Profile_pic, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}