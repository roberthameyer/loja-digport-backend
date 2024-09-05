package main

import (
	_ "github.com/lib/pq"
	"github.com/roberthameyer/loja-digport-backend/db"
)

func main() {
	db.InitDB()
	StartServer()
}
