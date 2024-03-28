package query

import (
	"fmt"

	"github.com/RabbitLuke/seminar-search/dbSetup"
)

type UserInfo struct {
	UserId      int    `json:"UserID"`
	F_name      string `json:"First_Name"`
	L_name      string `json:"Last_Name"`
	Faculty     int    `json:"Faculty"`
	Email       string `json:"Email"`
	Profile_pic string `json:"Profile_pic"`
	Password    string `json:"Password"`
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
	//Please make sure that you add the ID at the END of the statement! Otherwise it won't write to the DB
	res, err := stmt.Exec(f_name, l_name, faculty, email, profile_pic, password, userID)
	if err != nil {
		fmt.Println("Error executing update query:", err)
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Rows affected: %d\n", rowsAffected)
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

func GetUserFacultyAndSeminary(email string) (*GetFacultySeminars, error) {

	//return get user faculty and seminars linked to faculty
	var facultyQuery GetFaculty
	err := dbSetup.DB.QueryRow(`SELECT 
	ui.Faculty AS 'FacultyID',
	f.Name AS 'FacultyName'
	FROM user_information ui
	JOIN faculty f ON ui.Faculty = f.facultyID
	WHERE Email = ?`, email).Scan(&facultyQuery.FacultyID, &facultyQuery.FacultyName)
	if err != nil {
		return nil, err
	}

	var seminarsQuery []GetSeminar
	rows, seminarErr := dbSetup.DB.Query(`SELECT 
	si.seminarID AS 'SeminarID',
	si.Title AS 'Title',
	si.Duration AS 'Duration',
	si.Date AS 'Date',
	si.Time AS 'Time',
	si.Location AS 'Location',
	si.no_of_seats AS 'NoOfSeats',
	si.cover_photo AS 'CoverPhoto'	
	FROM seminar_info si
	WHERE si.facultyID = ?`, facultyQuery.FacultyID)
	if seminarErr != nil {
		return nil, err
	}

	for rows.Next() {
		// Create a new User object
		var seminar GetSeminar
		// Scan the values from the current row into the User object
		if err := rows.Scan(
			&seminar.SeminarID,
			&seminar.Title,
			&seminar.Duration,
			&seminar.Date,
			&seminar.Time,
			&seminar.Location,
			&seminar.NoOfSeats,
			&seminar.CoverPhoto); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		// Append the User object to the array
		seminarsQuery = append(seminarsQuery, seminar)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	var response GetFacultySeminars
	response.Faculty = facultyQuery
	response.Seminars = seminarsQuery
	return &response, nil
}

func GetUserByEmail(email string) (*UserInfo, error) {
	if dbSetup.DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	var user UserInfo
	err := dbSetup.DB.QueryRow("SELECT * FROM user_information WHERE Email = ?", email).Scan(&user.UserId, &user.F_name, &user.L_name, &user.Faculty, &user.Email, &user.Profile_pic, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type GetFacultySeminars struct {
	Faculty  GetFaculty   `json:"faculty"`
	Seminars []GetSeminar `json:"seminars"`
}

type GetFaculty struct {
	FacultyName string
	FacultyID   int
}

type GetSeminar struct {
	SeminarID  int     `json:"seminarID"`
	Title      string  `json:"title"`
	Duration   float32 `json:"duration"`
	Date       string  `json:"date"`
	Time       string  `json:"time"`
	Location   string  `json:"location"`
	NoOfSeats  int     `json:"no_of_seats"`
	CoverPhoto string  `json:"cover_photo"`
}
