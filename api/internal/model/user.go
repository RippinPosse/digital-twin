package model

type User struct {
	ID         string  `json:"id"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Patronymic *string `json:"patronymic"`
}
