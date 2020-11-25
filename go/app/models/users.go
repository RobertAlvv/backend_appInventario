package models

import (
	"errors"
)

type Users []User

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	TokenJWT string `json:"token_jwt"`
}

func GetUsers() Users {

	users := Users{}
	sql := `SELECT id, username, password FROM users`
	rows, _ := Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}
	return users
}

func GetUserById(id int64) *User {
	user := new(User)
	sql := `SELECT id, username, password FROM users where id = $1`
	row, _ := Query(sql, id)
	for row.Next() {
		row.Scan(&user.Id, &user.Username, &user.Password)
	}
	return user
}

func (this *User) ExistUsername() (bool, error) {
	var cantidad int
	sql := `SELECT COUNT(*) as cantidad FROM users WHERE username = $1`
	row, _ := Query(sql, this.Username)
	for row.Next() {
		row.Scan(&cantidad)
	}
	if cantidad > 0 {
		return true, errors.New("This username already exists")
	}
	return false, nil
}

func (this *User) GetUserID() {
	sql := `SELECT id, password FROM users WHERE username = $1 AND password = MD5($2)`
	row, _ := Query(sql, this.Username, this.Password)
	for row.Next() {
		row.Scan(&this.Id, &this.Password)
	}
}

func DeleteUserById(id int64) int64 {
	sql := `DELETE FROM users WHERE id = $1`
	result, _ := Exec(sql, id)
	rowAffect, _ := result.RowsAffected()
	return rowAffect
}

func (this *User) Save() {
	if this.Id == 0 {
		this.insert()
	} else {
		this.update()
	}
}

func (this *User) insert() {
	sql := `INSERT INTO users (username, password) VALUES ($1, MD5($2)) RETURNING id, password`
	row, _ := Query(sql, &this.Username, &this.Password)

	for row.Next() {
		row.Scan(&this.Id, &this.Password)
	}
}

func (this *User) update() {
	sql := `UPDATE users SET username = $1, password = MD5($2) where id = $3 RETURNING password`
	row, _ := Query(sql, &this.Username, &this.Password, &this.Id)
	for row.Next() {
		row.Scan(&this.Password)
	}
}
