package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@/todo?charset=utf8")
	if err != nil {
		log.Fatalln("Mysql Connect Error")
	}
	err = CreateTable()
	if err != nil {
		log.Fatalln("Create Table Error")
	}
}

type Todo struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetTodo(id int) (*Todo, error) {
	todo := new(Todo)
	err := Db.QueryRow("select id,title,content from t_todo where id=?", id).Scan(&todo.Id, &todo.Title, &todo.Content)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (todo *Todo) Create() error {
	stmt, err := Db.Prepare("insert into t_todo (title,content) values (?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(todo.Title, todo.Content)
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
