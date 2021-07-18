package model

import (
	"personal-site/log"
	"personal-site/utils"
)

// User表结构
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type ResponseUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

// 转换数据封装
func converUser(user User) ResponseUser {
	response := ResponseUser{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Token: user.Token,
	}

	return response
}

// 依据邮箱和密码验证用户
func QueryUserOfEmailAndPasswd(email string, password string) (User, error) {
	user := User{}

	// 查询并填充数据
	sql := "SELECT * FROM user_t WHERE email = $1 AND password = $2"
	err := Db.QueryRow(sql, email, utils.EncodeMd5(password)).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Token)

	if err != nil {
		log.Warning(err.Error())
	}

	return user, err
}

// 通过token查询用户信息
func QueryUserOfToken(token string) (ResponseUser, error) {
	user := User{}
	// 查询并填充数据
	err := Db.QueryRow("SELECT id, email, name, token FROM user_t WHERE token = $1", token).
		Scan(&user.Id, &user.Email, &user.Name, &user.Token)

	if err != nil {
		log.Warning(err.Error())
	}

	responseUser := converUser(user)

	return responseUser, err
}

// 插入新用户
func InsertUser(email string, name string, password string) error {
	stmt, err := Db.Prepare("INSERT INTO user_t (name, email, password, token) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Warning(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, utils.EncodeMd5(password), utils.EncodeToken())
	if err != nil {
		log.Warning(err.Error())
	}

	return err
}

// 更新用户信息
func UpdateUser(id string, email string, name string) error {
	stmt, err := Db.Prepare("UPDATE user_t SET email = $1, name = $2 WHERE id = $3")
	if err != nil {
		log.Warning(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, name, id)
	if err != nil {
		log.Warning(err.Error())
	}

	return err
}