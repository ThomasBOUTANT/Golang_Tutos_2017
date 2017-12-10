package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}


	for i:=0; i<10; i++{
		fmt.Println(".....", i)
	}

	j:=0
	for j<10{
		fmt.Println("**", j)
		j++
	}

	x := 5
	for {
		fmt.Println("Do stuff 1 - ",x)
		x+=3
		if x > 25 {
			break
		}
	}

	for y:=5; y<25; y+=3{
		fmt.Println("Do stuff 2 - ",y)
	}

	a := 3
	for z:=5; a <25; z+=3{
		fmt.Println("Do stuff 3 - ",z)
		a+=4
	}

}