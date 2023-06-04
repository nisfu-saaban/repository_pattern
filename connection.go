package repositorypattern

import (
	"database/sql"
	"time"
)

func GetConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "root:tryHackme07@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
