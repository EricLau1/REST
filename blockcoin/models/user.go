package models

import (
	"api/utils"
)

type User struct {
	UUID 	  uint32 `json:"_id"`
	Nickname  string `json:"nickname"`
	Email 	  string `json:"email"`
	Password  string `json:"password"`
	Status 	  byte 	 `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUser(user User) (bool, error) {
	con := Connect()
	defer con.Close()
	tx, err := con.Begin()
	if err != nil {
		return false, err
	}
	sql := "insert into users (nickname, email, password) values ($1, $2, $3) returning uuid"
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		hashedPassword, err := utils.Bcrypt(user.Password)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		err = stmt.QueryRow(user.Nickname, user.Email, hashedPassword).Scan(&user.UUID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	sql = "insert into wallet (public_key, usr) values ($1, $2)"
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		var wallet = Wallet{User:user}
		wallet.generatePublicKey()
		_, err = stmt.Exec(wallet.PublicKey, wallet.User.UUID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	return true, tx.Commit()
}

func GetUsers() ([]User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users order by uuid asc"
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var users []User
	for rs.Next() {
		var user User
		err := rs.Scan(&user.UUID, &user.Nickname, &user.Email, &user.Password,
		&user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByNickname(nickname string) (User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users where nickname = $1 limit 1"
	rs, err := con.Query(sql, nickname)
	if err != nil {
		return User{}, err
	}
	defer rs.Close()
	var user User
	if rs.Next() {
		err := rs.Scan(&user.UUID, &user.Nickname, &user.Email, &user.Password,
			&user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users where email = $1 limit 1"
	rs, err := con.Query(sql, email)
	if err != nil {
		return User{}, err
	}
	defer rs.Close()
	var user User
	if rs.Next() {
		err := rs.Scan(&user.UUID, &user.Nickname, &user.Email, &user.Password,
			&user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func GetUserById(id uint32) (User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users where uuid = $1 limit 1"
	rs, err := con.Query(sql, id)
	if err != nil {
		return User{}, err
	}
	defer rs.Close()
	var user User
	if rs.Next() {
		err := rs.Scan(&user.UUID, &user.Nickname, &user.Email, &user.Password,
			&user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func UpdateUser(user User) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "update users set nickname = $1, email = $2 where uuid = $3"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(user.Nickname, user.Email, user.UUID)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func DeleteUser(id uint32) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "delete from users where uuid = $1"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func ConfirmAccount(user User) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "update users set status = $1 where nickname = $2"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(1, user.Nickname)
	if err != nil {
		return 0, err
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		return 0, err
	} 
	return rows, nil
}