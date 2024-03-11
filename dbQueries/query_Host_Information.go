package query

import (
	"fmt"

	"github.com/RabbitLuke/seminar-search/dbSetup"
)

type HostInfo struct {
	HostID              int    `json:"HostID"`
	F_name              string `json:"First_Name"`
	L_name              string `json:"Last_Name"`
	Faculty             int    `json:"Faculty"`
	Qualifications int `json:"Qualifications"`
	Years_of_Experience int    `json:"Years_of_Experience"`
	Email               string `json:"Email"`
	Profile_pic         string `json:"Profile_pic"`
	Password            string `json:"Password"`
}

func InsertHost(f_name string, l_name string, faculty int, Qualifications int, years_of_experience int, email string, profile_pic string, password string) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	// Prepare the SQL statement
	stmt, err := dbSetup.DB.Prepare("INSERT INTO host_information (First_Name, Last_Name, Qualifications, Faculty, Years_of_Experience, Email, Profile_pic, Password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	fmt.Println(f_name, l_name, email)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the provided values
	_, err = stmt.Exec(f_name, l_name, faculty, Qualifications, years_of_experience, email, profile_pic, password)
	if err != nil {
		return err
	}

	fmt.Println("Host successfully created!")
	return nil
}

func DeleteHost(HostID int) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("DELETE FROM host_information WHERE HostID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(HostID)
	if err != nil {
		return err
	}

	fmt.Printf("Host with ID %d deleted successfully!\n", HostID)
	return nil
}

func UpdateHost(hostID int, f_name string, l_name string, faculty int, Qualifications int, years_of_experience int, email string, profile_pic string, password string) error {
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("UPDATE host_information SET First_Name = ?, Last_Name = ?, Faculty = ?, Qualifications = ?, Years_of_Experience = ?, Email = ?, Profile_pic = ?, Password = ? WHERE HostID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(hostID, f_name, l_name, faculty, Qualifications, years_of_experience, email, profile_pic, password)
	if err != nil {
		return err
	}

	fmt.Printf("Host with ID %d updated successfully!\n", hostID)
	return nil
}

func SelectHostByID(hostID int) (*HostInfo, error) {
	if dbSetup.DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	var host HostInfo
	err := dbSetup.DB.QueryRow("SELECT * FROM host_information WHERE HostID = ?", hostID).Scan(&host.F_name, &host.L_name, &host.Faculty, &host.Qualifications, &host.Years_of_Experience, &host.Email, &host.Profile_pic, &host.Password)
	if err != nil {
		return nil, err
	}

	return &host, nil
}
