package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

var DB *sql.DB

func initDb() error {
	var err error
	dsn := "root:@tcp(localhost:3306)/test_demo"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

func testQueryData() {
	sqlstr := "select id, name from users where id = ?;"
	row := DB.QueryRow(sqlstr, 1)
	var user User
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		fmt.Printf("query db failed, err:%v\n", err)
	}

	fmt.Printf("id=%d,name=%s\n", user.Id, user.Name)
}

func testQueryMutilRow() {
	sqlstr := "select id, name from users where id >= ?;"
	rows, err := DB.Query(sqlstr, 1)
	if err != nil {
		fmt.Printf("query db failed, err:%v\n", err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			fmt.Printf("query db failed, err:%v\n", err)
		}

		fmt.Printf("id=%d,name=%s\n", user.Id, user.Name)
	}

}

func testPrepareData() {
	sqlstr := "select id, name from users where id >= ?;"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	rows, err := stmt.Query(1)
	if err != nil {
		fmt.Printf("stmt failed, err:%v\n", err)
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			fmt.Printf("query db failed, err:%v\n", err)
		}

		fmt.Printf("id=%d,name=%s\n", user.Id, user.Name)
	}
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
	}
	defer DB.Close()

	testQueryData()
	testQueryMutilRow()
	testPrepareData()
}
