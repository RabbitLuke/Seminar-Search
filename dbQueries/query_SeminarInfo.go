package query

import (
	"fmt"
	"time"

	"github.com/RabbitLuke/seminar-search/dbSetup"
)

type SeminarInfo struct {
	SeminarID  int     `json:"seminarID"`
	Title      string  `json:"Title"`
	FacultyID  int     `json:"facultyID"`
	Duration   float32 `json:"Duration"`
	Date       string  `json:"Date"`
	Time       string  `json:"Time"`
	Location   string  `json:"Location"`
	NoOfSeats  int     `json:"no_of_seats"`
	CoverPhoto string  `json:"cover_photo"`
}

func InsertSeminar(title string, facultyID int, duration float32, date string, time string, location string, noOfSeats int, coverPhoto string) error {
	// Ensure that the database is initialized
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	stmt, err := dbSetup.DB.Prepare("INSERT INTO seminar_info (Title, facultyID, Duration, Date, Time, Location, no_of_seats, cover_photo) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
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

func UpdateSeminar(seminarID int, title string, facultyID int, duration float32, dateString string, timeString string, location string, noOfSeats int, coverPhoto string) error {
	if dbSetup.DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	// Parse the date string into a time.Time value
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return fmt.Errorf("error parsing date: %v", err)
	}

	// Parse the time string into a time.Time value
	timeOfDay, err := time.Parse("15:04:05", timeString)
	if err != nil {
		return fmt.Errorf("error parsing time: %v", err)
	}

	// Combine the date and time into a single time.Time value
	dateTime := time.Date(date.Year(), date.Month(), date.Day(), timeOfDay.Hour(), timeOfDay.Minute(), timeOfDay.Second(), 0, time.UTC)

	stmt, err := dbSetup.DB.Prepare("UPDATE seminar_info SET Title = ?, facultyID = ?, Duration = ?, Date = ?, Time = ?, Location = ?, no_of_seats = ?, cover_photo = ? WHERE seminarID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, facultyID, duration, dateTime, location, noOfSeats, coverPhoto, seminarID)
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
		if err := rows.Scan(&seminar.SeminarID, &seminar.Title, &seminar.Duration, &seminar.FacultyID, &seminar.Date, &seminar.Time, &seminar.Location, &seminar.NoOfSeats, &seminar.CoverPhoto); err != nil {
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
	err := dbSetup.DB.QueryRow("SELECT * FROM seminar_info WHERE seminarID = ?", seminarID).Scan(&seminar.SeminarID, &seminar.Title, &seminar.Duration, &seminar.FacultyID, &seminar.Date, &seminar.Time, &seminar.Location, &seminar.NoOfSeats, &seminar.CoverPhoto)
	if err != nil {
		return nil, err
	}

	return &seminar, nil
}
