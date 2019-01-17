package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

 const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "test123"
  dbname   = "postgres"
)


func main() {

 



  psqlInfo := fmt.Sprintf("host=%s  port=%d  user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

  // **************Inserted data into Table*****************

  // sqlStatement := `
  //   INSERT INTO users (age, email, first_name, last_name)
  //   VALUES ($1, $2, $3, $4)
  //   RETURNING id`
  // id := 0
  // err = db.QueryRow(sqlStatement, 21, "aniket@gmail.com", "Aniket", "Amin").Scan(&id)
  // if err != nil {
  //   panic(err)
  // }
  // fmt.Println("New record ID is:", id)

  // rows, err := db.Query("select * from users")
  // if err != nil {
  //   fmt.Println(err)
  // }

  // fmt.Println(rows)

  //*************Delete data from table*******************

  // sqlStatement := `
  //   DELETE FROM users
  //   WHERE id = $1;`
  //   _, err = db.Exec(sqlStatement, 15)
  // if err != nil {
  //   panic(err)
  // }

  // fmt.Println("Successfully Deleted")

  //****************update***************

  // sqlStatement = `
  //   UPDATE users
  //   SET first_name = $2, last_name = $3
  //   WHERE id = $1;`
  // res, err2 := db.Exec(sqlStatement, 16, "Mayank", "Prajapati")
  // if err2 != nil {
  //   panic(err)
  // }
  // count, err := res.RowsAffected()
  // if err != nil {
  //   panic(err)
  // }
  // fmt.Println(count)
  // fmt.Println("Successfully update")

  //*********************select*************
  
  rows, err := db.Query("SELECT * FROM users")
        if err !=nil {
          panic(err)
        }

        for rows.Next() {
            var id int
            var age string
            var email string
            var first_name string
            var last_name string
            err = rows.Scan(&id, &age, &first_name, &last_name, &email)
            if err != nil {
              panic(err)
            }
            fmt.Println("ID | AGE | EMAIL | FIRST_NAME | LAST_NAME ")
            fmt.Printf("%3v | %8v | %6v | %6v | %6v\n", id, age, first_name, last_name, email)
        }

}