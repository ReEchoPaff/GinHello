package model

import (
	"GinHello/initDB"
	"database/sql"
	"log"
)

type UserModel struct {
	Id       int64          `json:"id" form:"id"`
	Email    string         `json:"email" form:"email" binding:"email"`
	Password string         `json:"password" form:"password" `
	Avatar   sql.NullString `json:"avatar" form:"avatar"`
}

// user保存
func (user *UserModel) Save() int64 {
	result, e := initDB.Db.Exec("insert into ginhello.user (email, password) values (?,?);", user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

// 通过email返回用户信息
func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where email = ?;", user.Email)
	e := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Panicln(e)
	}
	return u
}

// 通过id返回用户信息
func (user *UserModel) QueryById(id int) (UserModel, error) {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where id = ?;", id)
	e := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Panicln(e)
	}
	return u, e
}

// 更新密码，头像，id
func (user *UserModel) Update(id int) error {
	var stmt, e = initDB.Db.Prepare("update user set password=?,avatar=?  where id=? ")
	if e != nil {
		log.Panicln("发生了错误", e.Error())
	}
	_, e = stmt.Exec(user.Password, user.Avatar.String, user.Id)
	if e != nil {
		log.Panicln("错误 e", e.Error())
	}

	return e
}
