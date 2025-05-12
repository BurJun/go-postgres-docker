package main

import ( 
    "fmt"
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    connStr := "host=localhost port=5432 user=dev password=dev dbname=appdb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
	log.Fatal("Errors connection to the db",err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
       log.Fatal("DB don't touch",err)
    }
    fmt.Println("Successful connection to the db")

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
	    id serial primary key,
	    name text not null
    )`)
    if err != nil {
       log.Fatal("Error create table",err)
    }

    _, err = db.Exec(`INSERT INTO users (name) VALUES ($1)`, "Alice")
    if err != nil {
       log.Fatal("Error insert into table",err)
    }

    fmt.Println("Values insert")
}


