package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UserDetail struct {
	Name     string
	Username string
	Password string //you should use various cryptographic types to store passwords.
}

type UserList []UserDetail

var db *sql.DB
var err error

func CommitUser(user UserDetail) bool {
	// createConnection()
	db, err := sql.Open("mysql", "root:root@/golangwebapp")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO users values(?,?,?);")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.Name, user.Username, user.Password)
	numberOfRows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if numberOfRows > 0 {
		return true
	}
	return false
}

func ValidatePassword(user UserDetail) bool {
	db, err := sql.Open("mysql", "root:root@/golangwebapp")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT password from users where username=? and password = ?`, user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	if rows.Next() {
		return true
	} else {
		return false
	}
	// var password string
	// for rows.Next() {
	// 	err = rows.Scan(&password)
	// 	if err != nil{
	// 		fmt.Println(err)
	// 	}
	// }

}

func GetAllUsers() UserList {
	db, err := sql.Open("mysql", "root:root@/golangwebapp")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT name, username from users;`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var name, username string
	var users UserList
	for rows.Next() {
		err := rows.Scan(&name, &username)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, UserDetail{name, username, "***"})
	}
	return users
}
