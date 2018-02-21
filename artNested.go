package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Artists struct {
	XMLName xml.Name `xml:"artists"`
	Artist  []Artist `xml:"artist"`
}

type Artist struct {
	XMLName xml.Name `xml:"artist"`
	ArtistID uint32 `xml:" id,omitempty"`
	ArtistName string `xml:" name,omitempty"`
	Profile string `xml:" profile,omitempty"`
	RealName string `xml:" realname,omitempty"`
	Members MembersID `xml:"members,omitempty"`
}

type MembersID struct {
	MemberID []uint32 `xml:"id,omitempty"`
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

	for i := 0; i < len(artists.Artist); i++ {
		fmt.Println("ArtistID: " + fmt.Sprint(artists.Artist[i].ArtistID))
		fmt.Println("Name: " + artists.Artist[i].ArtistName)
		fmt.Println("Real Name: " + artists.Artist[i].RealName)
		fmt.Println("Profile: " + artists.Artist[i].Profile)
		fmt.Println("")
		fmt.Printf("%v\n",artists.Artist[i].Members)
		fmt.Println("")
	}
}
