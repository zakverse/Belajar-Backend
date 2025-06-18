package db

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB

func ConnectDatabase() {
    // Format DSN: username:password@tcp(host:port)/dbname
    dsn := "root:DZAKI2006@tcp(127.0.0.1:3306)/produk"
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Gagal koneksi database:", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatal("Database tidak responsif:", err)
    }

    log.Println("Database terkoneksi.")
}
