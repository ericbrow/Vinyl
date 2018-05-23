/* This works great.  However manually writing each indiviual node and relationships
takes far to long.  Going to change this file to time it, and then try a few other
options and try those. */

package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	_ "gopkg.in/cq.v1"
)

type Artists struct {
	XMLName xml.Name `xml:"artists"`
	Artist  []Artist `xml:"artist"`
}

type Artist struct {
	XMLName    xml.Name  `xml:"artist"`
	ArtistID   uint32    `xml:" id,omitempty"`
	ArtistName string    `xml:" name,omitempty"`
	Profile    string    `xml:" profile,omitempty"`
	RealName   string    `xml:" realname,omitempty"`
	Members    MembersID `xml:"members,omitempty"`
}

type MembersID struct {
	MemberID   []uint32 `xml:"id,omitempty"`
	MemberName []string `xml:"name,omitempty"`
}

func check(e error) {
	if e != nil {
		fmt.Print(e)
		panic(e)
	}
}

func main() {

	start := time.Now()
	//r := new(big.Int)
	//fmt.Println(r.Binomial(1000, 10))

	//connect to text file
	xmlFile, err := os.Open("short_artists.xml")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Successfully opened artists file")
	//defer xmlFile.Close()

	//connect to database
	db, err := sql.Open("neo4j-cypher", "http://neo4j:Knd7r53p@localhost:7474")

	if err != nil {
		log.Println("error connecting to neo4j:", err)
	}
	defer db.Close()

	//read and unmarshal the XML
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var artists Artists
	xml.Unmarshal(byteValue, &artists)

	//create string for cypher creating artists nodes
	//go through file and create queries

	for i := 0; i < len(artists.Artist); i++ {
		//fmt.Println()
		//fmt.Println("Now printing Artist " + fmt.Sprint(i))
		//fmt.Println()
		//start node
		//update to merging - http://neo4j.com/docs/developer-manual/current/cypher/clauses/merge/#merge-merge-single-node-with-properties
		//https://stackoverflow.com/questions/35255540/neo4j-add-update-properties-if-node-exists/35255802
		myQuery := "merge (a:Artist { artId: \"" + fmt.Sprint(artists.Artist[i].ArtistID)
		myQuery += "\", name: \"" + strings.Replace(artists.Artist[i].ArtistName, `"`, `\"`, -1) + "\""

		//write real name if any
		if len(artists.Artist[i].RealName) > 0 {
			myQuery += ", realname: \"" + strings.Replace(artists.Artist[i].RealName, `"`, `\"`, -1) + "\""
		}

		//complete writing create node query
		myQuery += "})\n"
		//write it out to see what is being written
		//fmt.Println(myQuery)

		//write to database
		stmt, err := db.Prepare(myQuery)
		if err != nil {
			log.Println("Error at Preparing myQuery statement")
		}

		//no idea what this next part does
		rows, err := stmt.Query("wefreema")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		// var friend string
		// for rows.Next() {
		// 	err := rows.Scan(&friend)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	log.Println(friend)
		// }

		//blank out query
		myQuery = ""

		//write query for relationships
		if len(artists.Artist[i].Members.MemberID) > 0 {
			for j := 0; j < len(artists.Artist[i].Members.MemberID); j++ {
				memberQuery := "merge (a:Artist {artId: \"" + fmt.Sprint(artists.Artist[i].Members.MemberID[j])
				memberQuery += "\", name:\"" + strings.Replace(fmt.Sprint(artists.Artist[i].Members.MemberName[j]), `"`, `\"`, -1) + "\"})\n"

				relateQuery := "match (a:Artist), (b:Artist) where a.artId=\""
				relateQuery += fmt.Sprint(artists.Artist[i].Members.MemberID[j]) + "\" and b.artId=\""
				relateQuery += fmt.Sprint(artists.Artist[i].ArtistID) + "\"\n"
				relateQuery += "merge (a)-[r:memberof]->(b)\n"

				//fmt.Println(memberQuery)
				stmt, err := db.Prepare(memberQuery)
				if err != nil {
					log.Println("Error at Preparing memberQuery statement")
				}

				//no idea what this next part does
				rows, err := stmt.Query("wefreema")
				if err != nil {
					log.Fatal(err)
				}
				defer rows.Close()

				//fmt.Println(relateQuery)
				stmt2, err := db.Prepare(relateQuery)
				if err != nil {
					log.Println("Error at Preparing releateQuery statement")
				}

				//no idea what this next part does
				rows2, err := stmt2.Query("technige")
				if err != nil {
					log.Fatal(err)
				}
				defer rows2.Close()

				memberQuery = ""
				relateQuery = ""

			}
		}
	}

	xmlFile.Close()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
