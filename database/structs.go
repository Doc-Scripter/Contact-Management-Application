package database

type Contacts struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Phone int `json:"phone"`
}