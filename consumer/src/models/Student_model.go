package models

type Student struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Age         uint8  `json:"age"`
	PhoneNumber uint64 `json:"phoneNumber"`
}