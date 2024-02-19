// package models

// import (
// 	"errors"

// 	"example.com/rest-api/db"
// 	"example.com/rest-api/utils"
// )

// type User struct {
// 	ID       int64
// 	Email    string `binding:"required"`
// 	Password string `binding:"required"`
// }

// func (u User) Save() error {
//     query := "INSERT INTO users(email, password) VALUES ($1, $2) RETURNING id"
//     stmt, err := db.DB.Prepare(query)
//     if err != nil {
//         return err
//     }
//     defer stmt.Close()

//     hashedPassword, err := utils.HashPassword(u.Password)
//     if err != nil {
//         return err
//     }

//     var userId int64
//     err = stmt.QueryRow(u.Email, hashedPassword).Scan(&userId)
//     if err != nil {
//         return err
//     }

//     u.ID = userId
//     return nil
// }

// func (u *User) ValidateCredentials() error {
// 	query := "SELECT id, password FROM users WHERE email = ?"
// 	row := db.DB.QueryRow(query, u.Email)

// 	var retrievedPassword string
// 	err := row.Scan(&u.ID, &retrievedPassword)

// 	if err != nil {
// 		return errors.New("Credentials invalid")
// 	}

// 	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

// 	if !passwordIsValid {
// 		return errors.New("Credentials invalid")
// 	}

// 	return nil
// }

package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
    ID       int64
    Email    string `binding:"required"`
    Password string `binding:"required"`
}

func (u *User) Save() error {
    query := "INSERT INTO users(email, password) VALUES ($1, $2) RETURNING id"
    stmt, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    hashedPassword, err := utils.HashPassword(u.Password)
    if err != nil {
        return err
    }

    var userId int64
    err = stmt.QueryRow(u.Email, hashedPassword).Scan(&userId)
    if err != nil {
        return err
    }

    u.ID = userId
    return nil
}

func (u *User) ValidateCredentials() error {
    query := "SELECT id, password FROM users WHERE email = $1"
    row := db.DB.QueryRow(query, u.Email)

    var retrievedPassword string
    var userID int64
    err := row.Scan(&userID, &retrievedPassword)

    if err != nil {
        return errors.New("Credentials invalid")
    }

    passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

    if !passwordIsValid {
        return errors.New("Credentials invalid")
    }

    u.ID = userID
    return nil
}
