package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	user   string = "postgres"
	pass   string = "123456"
	host   string = "localhost"
	port   int    = 5450
	dbname string = "inventariado_alv"
)

func init() {
	CreateConectionDB()
}

func CreateConectionDB() {
	if database, err := sql.Open("postgres", GenerateConection()); err != nil {
		panic(err)
	} else {
		db = database
		fmt.Printf("Conexion Exitosa!\n")
	}
}

//CloseConectionDB is that
func CloseConectionDB() {
	db.Close()
}

//Ping is that
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func createTable(schema string, tableName string) {
	if existTable(tableName) {
		return
	}
	Exec(schema)
}

func existTable(tableName string) bool {
	sql := fmt.Sprintf("SELECT * FROM pg_catalog.pg_tables WHERE tablename = '%s';", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

func Query(query string, arg ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, arg...)
	if err != nil {
		log.Print(err)
	}
	return rows, err
}

func Exec(query string, arg ...interface{}) (sql.Result, error) {
	Result, err := db.Exec(query, arg...)
	if err != nil {
		log.Println(err)
	}

	return Result, err
}

func GenerateConection() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port,
		user, pass, dbname)
}
