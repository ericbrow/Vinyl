package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	//"bytes"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
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
		panic(e)
	}
}

func main() {

	//connect to text file
	xmlFile, err := os.Open("short_artists.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened artists file")
	defer xmlFile.Close()

	//connect to database
	myDriver := bolt.NewDriver()
	conn, _ := myDriver.OpenNeo("bolt://localhost:7687")
	defer conn.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var artists Artists
	xml.Unmarshal(byteValue, &artists)

	//create string for cypher creating artists nodes

	//go through file and create queries

	for i := 0; i < len(artists.Artist); i++ {
		//start node
		//update to merging - http://neo4j.com/docs/developer-manual/current/cypher/clauses/merge/#merge-merge-single-node-with-properties
		//https://stackoverflow.com/questions/35255540/neo4j-add-update-properties-if-node-exists/35255802
		myQuery := "merge (a:Artist { id: \"" + fmt.Sprint(artists.Artist[i].ArtistID)
		myQuery += "\", name: \"" + (artists.Artist[i].ArtistName) + "\""

		//write real name if any
		if len(artists.Artist[i].RealName) > 0 {
			myQuery += ", realname: \"" + artists.Artist[i].RealName + "\""
		}

		//complete writing create node query
		myQuery += "})\n"
		fmt.Println(myQuery)

		check(err)

		//clear buffer
		myQuery = ""

		//write query for relationships
		if len(artists.Artist[i].Members.MemberID) > 0 {
			for j := 0; j < len(artists.Artist[i].Members.MemberID); j++ {
				memberQuery := "merge (a:Artist {id: \"" + fmt.Sprint(artists.Artist[i].Members.MemberID[j])
				memberQuery += "\", name:\"" + fmt.Sprint(artists.Artist[i].Members.MemberName[j]) + "\"})\n"

				relateQuery := "match (a:Artist)(b:artist) where a.id = "
				relateQuery += fmt.Sprint(artists.Artist[i].Members.MemberID[j]) + " and b.id = "
				relateQuery += fmt.Sprint(artists.Artist[i].ArtistID) + "\n"
				relateQuery += "merge (a)-[r:memberof]->(b)\n"

				fmt.Println(memberQuery)
				fmt.Println(relateQuery)
				memberQuery = ""
				relateQuery = ""
				check(err)
			}
		}
	}
}
