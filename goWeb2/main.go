package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// { // for creating users table
	// 	query := `
	//         CREATE TABLE users (
	//             id INT AUTO_INCREMENT,
	//             username TEXT NOT NULL,
	//             password TEXT NOT NULL,
	//             created_at DATETIME,
	//             PRIMARY KEY (id)
	//         );`

	// 	if _, err := db.Exec(query); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	{
		username := "ali"
		password := "password"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			fmt.Println("-1")
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{
		var (
			id        int
			username  string
			password  string
			createdAt []uint8
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			fmt.Println("0")
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			fmt.Println("1")
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var user user

			err := rows.Scan(&user.id, &user.username, &user.password, &user.createdAt)
			if err != nil {
				fmt.Println("2")
				log.Fatal(err)
			}
			users = append(users, user)
		}
		if err := rows.Err(); err != nil {
			fmt.Println("3")
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		_, err := db.Exec(`DELETE FROM users WHERE name = ?`, "ali")
		if err != nil {
			fmt.Println("4")
			log.Fatal(err)
		}
	}
}
