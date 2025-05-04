package models

import (
	"errors"
	"fmt"
	"rest_api/db"
	"rest_api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password)
			  VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u User) ValidateUser() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	fmt.Println("row", row)

	var retrivedRow string
	err := row.Scan(&u.ID ,&retrivedRow)

	fmt.Println("retrievedrow", row)

	if err != nil {
		return errors.New("credential invalid")
	}

	passwordIsValid := utils.CheckMatchingHash(u.Password, retrivedRow)

	if !passwordIsValid {
		return errors.New("credential invalid")
	}

	return nil
}

// func getAllUsers()([]User ,error){
// 	query := `SELECT * FROM users`

// 	stmt, err := db.DB.Prepare(query)

// 	if err!= nil{
// 		return nil, err
// 	}

// 	result, err := stmt.Exec(query)

// 	if err!= nil{
// 		return nil, err
// 	}

// 	return result, nil
// }
