package main

import (
	"os"
	"fmt"
	"encoding/xml"
	"io/ioutil"
)

type Query struct {
	Series Artist
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	ArtistList []Artist `xml:"Artist>"`
}

type Artist struct {
	// Have to specify where to find the series title since
	// the field of this struct doesn't match the xml tag
	ArtistID int `xml:"id>"`
	Name string
	Profile string
	RealName string
}


func (s Artist) String() string {
	return fmt.Printf(s.ArtistID, s.Name, s.RealName)
}

//func (p Profile) String() string {
//	return fmt.Sprintf(p.Profile)
//}

func main() {
	xmlFile, err := os.Open("short_artists.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(xmlFile, &q)

	fmt.Println(q.Series)
	for _, episode := range q.ArtistList {
		fmt.Printf("\t%s\n", episode)
	}
}
