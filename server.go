package main

import (
	"flag"
	"log"

	"github.com/gilperopiola/frutils"

	"github.com/gilperopiola/query-results-exporter/config"
)

var cfg config.MyConfig
var db MyDatabase

func main() {
	env := flag.String("env", "local", "local")
	flag.Parse()

	cfg.Setup(*env)
	db.Setup(cfg)
	defer db.Close()

	log.Println("starting query")

	rows, err := db.DB.Query("SELECT name, location, idStatus FROM servers WHERE physical = 1")
	defer rows.Close()
	if err != nil {
		log.Println("error 1: " + err.Error())
		return
	}

	out := ""
	for rows.Next() {
		name := ""
		location := ""
		idStatus := 0

		rows.Scan(&name, &location, &idStatus)

		out += name + "\t" + location + "\t" + frutils.ToString(idStatus) + "\n"
	}

	frutils.CreateFile("query.txt", out)
}
