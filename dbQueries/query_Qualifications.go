package query

import (
	"fmt"
	"github.com/RabbitLuke/seminar-search/dbSetup"
)

type Qualifications struct{
	QualID int `json:"QualID"`
	Name string `json:"Name"`
}

func SelectQualifications() ([]Qualifications, error) {
	if dbSetup.DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	rows, err := dbSetup.DB.Query("SELECT * FROM qualifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualifications []Qualifications

	for rows.Next() {
		var qualification Qualifications
		if err := rows.Scan(&qualification.QualID, &qualification.Name); err != nil {
			return nil, err
		}
		qualifications = append(qualifications, qualification)
	}

	return qualifications, nil
}