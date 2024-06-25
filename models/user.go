package models

type User struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Role_ID int    `json:"role_id"`
}
