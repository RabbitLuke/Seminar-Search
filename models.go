package main

type SeminarInfo struct {
	SeminarID int `json:"seminarID"`
	Title string `json:"Title"`
	FacultyID int `json:"facultyID"`
	Duration float32 `json:"Duration"`
	Date string `json:"Date"`
	Time string `json:"Time"`
	Location string `json:"Location"`
	NoOfSeats int `json:"no_of_seats"`
	CoverPhoto string `json:"cover_photo"`
}

type UserInfo struct {
	UserID int `json:"UserID"`
	FirstName string `json:"First_Name"`
	LastName string `json:"Last_Name"`
	Faculty int `json:"Faculty"`
	Email string `json:"Email"`
	ProfilePic string `json:"Profile_pic"`
	Password string `json:"Password"`
}

type TokenResponse struct {
	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}