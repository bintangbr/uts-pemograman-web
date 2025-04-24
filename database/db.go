package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "uts:utsHahaha@tcp(103.190.112.100:3306)/uts")
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal konek kee database:", err)
	}
	fmt.Println("Alhamdulillah Berhasil terhubung ke database dan jalan diport 8080")
}
