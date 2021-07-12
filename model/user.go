package model

import "personal-site/log"

// User表结构
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// 依据邮箱和密码验证用户
func AuthUser(email string, password string) (User, error) {
	user := User{}

	// 查询并填充数据
	sql := "SELECT * FROM user_t WHERE email = $1 AND password = $2"
	err := Db.QueryRow(sql, email, password).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Token)

	if err != nil {
		log.Info(err.Error())
	}

	return user, err
}

// 验证token
func AuthToken(token string) (string, error) {
	name := ""
	// 查询并填充数据
	err := Db.QueryRow("SELECT name FROM user_t WHERE token = $1", token).Scan(&name)

	if err != nil {
		log.Info(err.Error())
	}

	return name, err
}
