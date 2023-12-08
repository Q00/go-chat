package dto

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"userType"` // STUDENT or TEACHER
}
