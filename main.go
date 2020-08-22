package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type career struct {
	id string
	name string
	job_title string
	age int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:sammidev@tcp(127.0.0.1:3306)/db_belajargolang")
	if err != nil {
		return nil, err
	}

	return db, nil
}
func sqlQuery()  {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 19
	rows, err := db.Query("select id, name, job_title, age from career_level where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []career
	for rows.Next() {
		var each = career{}
		var err = rows.Scan(&each.id, &each.name, &each.job_title, &each.age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.id)
		fmt.Println(each.name)
		fmt.Println(each.job_title)
		fmt.Println(each.age)
	}
}
func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = career{}
	var id = "A1"
	err = db.
		QueryRow("select name, job_title from career_level where id = ?", id).
		Scan(&result.name, &result.job_title)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Name : ", result.name, "\nJob Title : ", result.job_title)
}
func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, job_title from career_level where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = career{}
	stmt.QueryRow("A2").Scan(&result1.name, &result1.job_title)
	fmt.Println("Name : ", result1.name, "\nJob Title : ", result1.job_title)

	var result2 = career{}
	stmt.QueryRow("A3").Scan(&result2.name, &result2.job_title)
	fmt.Println("Name : ", result2.name, "\nJob Title : ", result2.job_title)
}
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into career_level values (?, ?, ?, ?)", "A10", "Unknown", "Android Developer", 19)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update career_level set name = ? where id = ?", "Abdul Rauf", "A3")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from career_level where id = ?", "A10")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")
}
func sqlExecPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into career_level values (?, ?, ?, ?)")
	stmt.Exec("A10", "Otong Surotong", "IOS Engineer", 19)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("insert success!")
}

func main() {
	//sqlQuery()
	//sqlQueryRow()
	//sqlPrepare()
	//sqlExec()
	sqlExecPrepare()
}