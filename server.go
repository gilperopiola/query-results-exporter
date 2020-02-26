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

	rows, err := db.DB.Query("SELECT name, location, idKind, idStatus FROM servers WHERE physical = 1")
	defer rows.Close()
	if err != nil {
		log.Println("error 1: " + err.Error())
		return
	}

	out := ""
	for rows.Next() {
		name := ""
		location := ""
		idKind := 0
		idStatus := 0

		rows.Scan(&name, &location, &idKind, &idStatus)

		kind := getKindFromID(idKind)
		status := getStatusFromID(idStatus)

		out += name + "\t" + status + "\t" + kind + location + "\t" + "\n"
	}

	frutils.CreateFile("query.txt", out)
}

func getKindFromID(idKind int) string {
	if idKind == 0 {
		return "Desconocido"
	}
	if idKind == 1 {
		return "Appliance"
	}
	if idKind == 2 {
		return "Blade"
	}
	if idKind == 3 {
		return "Hoja"
	}
	if idKind == 4 {
		return "Librería"
	}
	if idKind == 5 {
		return "VM"
	}
	if idKind == 6 {
		return "PC"
	}
	if idKind == 7 {
		return "Servidor"
	}
	if idKind == 8 {
		return "Almacenamiento"
	}
	if idKind == 9 {
		return "Switch Blade"
	}
	if idKind == 10 {
		return "Switch FC"
	}
	return ""
}

func getStatusFromID(idStatus int) string {
	if idStatus == 0 {
		return "Desconocido"
	}
	if idStatus == 1 {
		return "Activo"
	}
	if idStatus == 2 {
		return "Desguasado"
	}
	if idStatus == 3 {
		return "Apagado"
	}
	if idStatus == 4 {
		return "Inactivo"
	}
	return ""
}
