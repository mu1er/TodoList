package model

type Todo struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User_id int    `json:"user_id"`
}

type Todos []*Todo

func GetAllTodo() (*Todos, error) {
	todos := &Todos{}
	rows, err := Db.Query("select id,title,content,user_id from t_todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		todo := new(Todo)
		_ = rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.User_id)
		todos = append(todos, todo)
	}
	return todos, nil
}
func GetTodo(id int) (*Todo, error) {
	todo := new(Todo)
	err := Db.QueryRow("select id,title,content,user_id from t_todo where id=?", id).Scan(&todo.Id, &todo.Title, &todo.Content, &todo.User_id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (user *User) Create() error {
	stmt, err := Db.Prepare("insert into t_todo (title,content,user_id) values (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(todo.Title, todo.Content, user.Id)
	if err != nil {
		return err
	}
	return nil
}
func (todo *Todo) Update() error {
	stmt, err := Db.Prepare("update t_todo set title=?,content=? where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(todo.Title, todo.Content, todo.Id)
	if err != nil {
		return err
	}
	return nil
}

func (todo *Todo) Delete() error {
	stmt, err := Db.Prepare("delete from t_todo where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(todo.Id)
	if err != nil {
		return err
	}
	return nil
}
