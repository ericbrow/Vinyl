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
	XMLName  xml.Name  `xml:"artist"`
	ArtistID uint32    `xml:"id"`
	Name     string    `xml:"name"`
	RealName string    `xml:"realname"`
	Profile  string    `xml:"profile"`
	Members  []Members `xml:"members"`
}

type Members struct {
	XMLName    xml.Name `xml:"members"`
	MemberID   uint32   `xml:"id"`
	MemberName string   `xml:"name"`
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
		fmt.Println("Name: " + artists.Artist[i].Name)
		fmt.Println("Real Name: " + artists.Artist[i].RealName)
		fmt.Println("Profile: " + artists.Artist[i].Profile)
		fmt.Println("Members: " + fmt.Sprint(len(artists.Artist[i].Members)))
		//  if len(artists.Artist.member.Members) > 0 {
		 //
		// 		for j := 0; j < len(artists.Artist.member.Members); j++ {
		// 		fmt.Println("    Member: " + fmt.Sprint(artists.Artist[i].member.Member[j].MemberID))
		// 		fmt.Println("    Name:   " + fmt.Println(artists.Artist[i].member.Member[j].MemberName))
		// 		}
		// 	}
	}
}
