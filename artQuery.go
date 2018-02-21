package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"bytes"
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
	MemberID []uint32 `xml:"id,omitempty"`
	//MemberName []string `xml:"name,omitempty"`
}

func check(e error) {
		if e != nil {
			panic(e)
		}
}

func main() {

	xmlFile, err := os.Open("short_artists.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened artists file")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var artists Artists
	xml.Unmarshal(byteValue, &artists)
	var buffer bytes.Buffer
	//create file for cypher creating artists nodes
	node, err := os.Create("artNode.txt")
	check(err)
	defer node.Close()
	//create file for cypher creating relationships to artists
	rel, err := os.Create("artRelate.txt")
	check(err)
	defer rel.Close()
	//go through file and create queries
	for i := 0; i < len(artists.Artist); i++ {
		//start node
		buffer.WriteString("create (a:Artist { id: " + fmt.Sprint(artists.Artist[i].ArtistID))
		buffer.WriteString(", name: \"" + (artists.Artist[i].ArtistName) + "\"")
		//write real name if any
		if(len(artists.Artist[i].RealName) > 0) {
		buffer.WriteString(", realname: \"" + artists.Artist[i].RealName + "\"") }
		//write profile if any - removed due to ditry strings in profile
		//if(len(artists.Artist[i].Profile) > 0) {
		//	buffer.WriteString(", profile: \"" + artists.Artist[i].Profile + "\"") }
		//complete writing create node query
		buffer.WriteString("})\n")
		_ , err := buffer.WriteTo(node)
	  check(err)
		fmt.Println("Wrote artist ID:" + fmt.Sprint(artists.Artist[i].ArtistID))
		//clear buffer
		buffer.Reset()
		//write query for relationships
		if(len(artists.Artist[i].Members.MemberID) > 0) {
			for j :=0; j < len(artists.Artist[i].Members.MemberID); j++ {
			buffer.WriteString("match (a:Artist)(b:artist) where a.id = ")
			buffer.WriteString(fmt.Sprint(artists.Artist[i].Members.MemberID[j]) + " and b.id = ")
			buffer.WriteString(fmt.Sprint(artists.Artist[i].ArtistID) + "\n")
			buffer.WriteString("create (a)-[r:memberof]->(b)\n")
			_ , err := buffer.WriteTo(rel)
			check(err)
			fmt.Println("Wrote relationship from" + fmt.Sprint(artists.Artist[i].ArtistID) + " to " + fmt.Sprint(artists.Artist[i].Members.MemberID[j]))
			buffer.Reset()
		}
		}
	}
}
