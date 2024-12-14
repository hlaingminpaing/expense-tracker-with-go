// package config

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB

// func ConnectDB() {
// 	var err error
// 	DB, err = sql.Open("mysql", "goadmin:Mygoexpensetracker@tcp(expenses.cvrlgpla3e5u.ap-southeast-1.rds.amazonaws.com:3306)/expense_tracker")
// 	if err != nil {
// 		log.Fatalf("Could not connect to the database: %v", err)
// 	}

// 	if err = DB.Ping(); err != nil {
// 		log.Fatalf("Could not ping the database: %v", err)
// 	}

// 	log.Println("Database connected successfully")
// }


package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "goadmin:Mygoexpensetracker@tcp(expenses.cvrlgpla3e5u.ap-southeast-1.rds.amazonaws.com:3306)/expense_tracker")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connected successfully")
}
