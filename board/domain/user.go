package domain

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Grade    string `json:"grade"`
}
