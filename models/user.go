package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Could not prepare statement.")
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic("Could not hash password.")
	}
	results, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		panic("Could not execute statement.")
	}
	id, err := results.LastInsertId()
	u.ID = id
	return err

}
func (u *User) ValidateCredentials() (bool, error) {
	query := "SELECT password FROM users WHERE email = ?"
	var hashedPassword string
	row := db.DB.QueryRow(query, u.Email)
	err := row.Scan(&hashedPassword)
	if err != nil {
		return false, err
	}
	return utils.CheckPasswordHash(u.Password, hashedPassword), nil
}
