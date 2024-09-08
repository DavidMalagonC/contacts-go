package config

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "os"
)

func NewConnection() (*sql.DB, error) {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    database := os.Getenv("DB_NAME")

    if user == "" || password == "" || host == "" || database == "" {
        return nil, fmt.Errorf("missing database environment variables")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
