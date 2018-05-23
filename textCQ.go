package main

import (
	"database/sql"
	"log"

	_ "gopkg.in/cq.v1"
)

func main() {
	db, err := sql.Open("neo4j-cypher", "http://neo4j:Knd7r53p@localhost:7474")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`merge (a:Artist { id: "1", name: "The Persuader", realname: "Jesper Dahlback"})`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query("wefreema")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var friend string
	for rows.Next() {
		err := rows.Scan(&friend)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(friend)
	}
}
