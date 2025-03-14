package models

type StudentCredentials struct {
	Email    string   `json:"email"`
	Student  Student  `json:"student"`
}