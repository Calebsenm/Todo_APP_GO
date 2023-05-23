package db 

import (
    _ "github.com/lib/pq"
    "database/sql"
    "fmt"
    "log"
)

const (
    DB_USER = "caleb"
    DB_PASSWORD = "12345"
    DB_NAME = "test"
    DB_HOST = "localhost"
    DB_PORT = 5432
    
)


func Connet() (*sql.DB, error ){
    
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST,DB_PORT, DB_USER , DB_PASSWORD , DB_NAME);
    //open the conexion 
    db, err := sql.Open("postgres",connStr );
    
    if err != nil {
        log.Fatal("Error to Open the database " , err);
    }
    // return the data 
    return db , err ;  
}
