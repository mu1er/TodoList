package model

const t_todo = `
CREATE TABLE IF NOT EXISTS t_todo(
	id int primary key auto_increment,
	title varchar(255) not null,
	content text
);
`

var Tables = [...]string{t_todo}

func CreateTable() error {
	tables := Tables
	for _, table := range tables {
		_, err := Db.Exec(table)
		if err != nil {
			return err
		}
	}
	return nil
}
