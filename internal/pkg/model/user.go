package model

type User struct {
	Id         int    `json:"id"`
	Passport   string `json:"passport"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}
