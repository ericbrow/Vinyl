package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Artists struct {
	XMLName xml.Name `xml:"artists"`
	Artist  []struct {
		ArtistID uint32 `xml:" id,omitempty"`
		ArtistName string `xml:" name,omitempty"`
		Profile string `xml:" profile,omitempty"`
		RealName string `xml:" realname,omitempty"`
		Members []struct {
			MemberId uint32 `xml:" members>id,omitempty"`
			MemberName string `xml:" members>name,omitempty"`
		}
	}
}

// type Artist struct {
// 	ArtistID uint32 `xml:" id,omitempty"`
// 	Members []*Members `xml:" members,omitempty"`
// 	ArtistName string `xml:" name,omitempty"`
// 	Profile string `xml:" profile,omitempty"`
// 	RealName string `xml:" realname,omitempty"`
// }
//
// type Members struct {
// 	MemberId uint32 `xml:" members>id,omitempty"`
// 	MemberName string `xml:" members>name,omitempty"`
// }

func main() {

	xmlFile, err := os.Open("short_artists.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened artists file")
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var artists Artists
	//var imembers IndMembers
 	xml.Unmarshal(byteValue, &artists)
	//xml.Unmarshal(byteValue, &imembers)
fmt.Println(artists.Artist)
	// for i := 0; i < len(artists.Artist); i++ {
	// 	fmt.Println("ArtistID: " + fmt.Sprint(artists.Artist[i].ArtistID))
	// 	fmt.Println("Name: " + artists.Artist[i].ArtistName)
	// 	fmt.Println("Real Name: " + artists.Artist[i].RealName)
	// 	fmt.Println("Profile: " + artists.Artist[i].Profile)
	// 	fmt.Println("")
	// 	fmt.Printf("%v\n",len(artists.Artist[i].Members))
	// 	fmt.Println("")
	//
	// }
		defer xmlFile.Close()
}
