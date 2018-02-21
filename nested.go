package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Artists struct {
	//XMLName xml.Name `xml:"artists"`
	Artist  []struct {
    ArtistID uint32 `xml:" id,omitempty"`
    ArtistName string `xml:" name,omitempty"`
    Profile string `xml:" profile,omitempty"`
    RealName string `xml:" realname,omitempty"`
    Members []struct {
      	MemberID uint32 `xml:"id,omitempty"`
        MemberName string `xml:"name,omitempty"`
    }
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
	//var imembers IndMembers
 	xml.Unmarshal(byteValue, &artists)
  fmt.Println(artists.Artist)
	}
