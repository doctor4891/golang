/*MYSQL CRUD*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //_ для подключки драйвера перед использованием
)

func main() {
	/*подключка к удаленной базе данных mysql*/
	db, err := sql.Open("mysql", "admin_golang:password@tcp(188.225.72.32:3306)/admin_golang")
	if err != nil {
		fmt.Println("error connection:", err)
	}

	//create(db)

	//read(db)

	//update(db)

	deleteRow(db)

}

//////////////////////			INSERT			//////////////////////
func create(db *sql.DB) {
	res, err := db.Exec("INSERT into goTypes (name) value ('string')")
	if err != nil {
		fmt.Println("error insert:", err)
	}

	fmt.Println(res.LastInsertId()) //2 <nil>
}

//////////////////////			SELECT			//////////////////////
func read(db *sql.DB) {
	res, err := db.Query("SELECT * from goTypes where id<?", 2)
	if err != nil {
		fmt.Println("error select:", err)
	}

	/*поля таблицы*/
	var name string
	var id int

	/*итерация по всем строкам результата*/
	for res.Next() {
		err := res.Scan(&name, &id)
		fmt.Println(err)
		fmt.Println(name, id)
	}
}

func update(db *sql.DB) {
	_, err := db.Query("update goTypes set name='integer' where id=?", 1)
	if err != nil {
		fmt.Println("error udpate:", err)
	}
}

func deleteRow(db *sql.DB) {
	_, err := db.Query("delete from goTypes where id=?", 1)
	if err != nil {
		fmt.Println("error delete:", err)
	}
}
