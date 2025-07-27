package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConectarDB() {

	var err error

	DB, err = sql.Open("sqlite3", "clientes.db")

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	result, err := DB.Exec("CREATE TABLE IF NOT EXISTS pedidos (cliente TEXT NOT NULL, pedido TEXT NOT NULL, telefono TEXT NOT NULL);")

	if err != nil {
		log.Fatal(result)
		log.Fatal(err)
	} else {
		fmt.Println(result)
		fmt.Println("Exito al conectar la base de datos")
	}
}
