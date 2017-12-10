package main

import (
  "encoding/xml"
  "fmt"
)

var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>test ! Hello everyone !!</loc>
   </sitemap>
</sitemapindex>
`)

type Sitemapindex struct {
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (e Location) String () string {
  //return fmt.Sprintf(e.Loc)
	return fmt.Sprintln(e.Loc)
}

func main() {
  bytes := washPostXML
  var s Sitemapindex
  xml.Unmarshal(bytes, &s)
  fmt.Println(s.Locations)
}