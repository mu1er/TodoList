package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []*User

func GetAllUser() (*Users, error) {
	users := &Users{}
	rows, err := Db.Query("select id,username,password,email from t_user")
	if err != nil {
		return nil, err
	}
	rows.Close()
	for rows.Next() {
		user := new(User)
		_ = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users, nil
}

func GetUser(id int) (*User, error) {
	user := new(User)
	err := Db.QueryRow("select id,username,password,email from t_user where id=?", id).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
