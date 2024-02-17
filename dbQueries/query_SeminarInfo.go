package query


import (
	"fmt"
	"time"
	"github.com/RabbitLuke/seminar-search/dbSetup"
)

type SeminarInfo struct {
	SeminarID int `json:"seminarID"`
	Title string `json:"Title"`
	FacultyID int `json:"Faculty"`
	Duration float32 `json:"Duration"`
	Date time.Time `json:"Date"`
	Time time.Time `json:"Time"`
	Location string `json:"Location"`
	NoOfSeats int `json:"no_of_seats"`
	CoverPhoto string `json:"cover_photo"`
}

func InsertSeminar(title string, facultyID int, duration float32, date time.Time, time time.Time, location string, noOfSeats string, coverPhoto string) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("INSERT INTO seminar_info (Title, Faculty, Duration, Date, Time, Location, no_of_seats, cover_photo) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, facultyID, duration, date, time, location, noOfSeats, coverPhoto)
	if err != nil {
		return err
	}

	fmt.Println("Seminar inserted successfully!")
	return nil
}

func DeleteSeminar(seminarID int) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("DELETE FROM seminar_info WHERE seminarID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(seminarID)
	if err != nil {
		return err
	}

	fmt.Printf("Seminar with ID %d deleted successfully!\n", seminarID)
	return nil
}

func UpdateSeminar(seminarID int, title string, facultyID int, duration float32, date time.Time, time time.Time, location string, noOfSeats string, coverPhoto string) error {
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("UPDATE seminar_info SET Title = ?, Faculty = ?, Duration = ?, Date = ?, Time = ?, Location = ?, no_of_seats = ?, cover_photo = ? WHERE seminarID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, facultyID, duration, date, time, location, noOfSeats, coverPhoto, seminarID)
	if err != nil {
		return err
	}

	fmt.Printf("Seminar with seminarID %d updated successfully!\n", seminarID)
	return nil
}

func SelectSeminars() ([]SeminarInfo, error) {
	if dbSetup.DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	rows, err := dbSetup.DB.Query("SELECT * FROM seminar_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seminars []SeminarInfo

	for rows.Next() {
		var seminar SeminarInfo
		if err := rows.Scan(&seminar.SeminarID); err != nil {
			return nil, err
		}
		seminars = append(seminars, seminar)
	}

	return seminars, nil
}

func SelectSeminarByID(seminarID int) (*SeminarInfo, error) {
	if dbSetup.DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	var seminar SeminarInfo
	err := dbSetup.DB.QueryRow("SELECT * FROM seminar_info WHERE seminarID = ?", seminarID).Scan(&seminar.SeminarID)
	if err != nil {
		return nil, err
	}

	return &seminar, nil
}