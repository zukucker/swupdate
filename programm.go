package main

import "fmt"
import "io/ioutil"
import "net/http"
import "encoding/xml"
import "regexp"


  type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"atom,attr"`
	Channel struct {
		Text string `xml:",chardata"`
		Link struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Title         string `xml:"title"`
		Description   string `xml:"description"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Guid        string `xml:"guid"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			Category    string `xml:"category"`
			PubDate     string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
} 


func main() {
  resp, err := http.Get("https://www.shopware.com/de/changelog-sw5/?sRss=1")
  if err != nil {
    fmt.Println(err)
    // handle err
  }
  defer resp.Body.Close()

  byteValue, _ := ioutil.ReadAll(resp.Body)
  var rss Rss
  xml.Unmarshal(byteValue, &rss)

  fmt.Println(rss.Channel.Title)
  fmt.Println(rss.Channel.LastBuildDate)
  fmt.Println("=======================")
  fmt.Println("Author: Alexander Froehling")
  fmt.Println("Contact: alexander.froehling@googlemail.com")
  fmt.Println("=======================")

  for _, s := range rss.Channel.Item {
    fmt.Println("\n", "🆕Version " + s.Title, "\n")

    var re = regexp.MustCompile(`SW-`)
    s := re.ReplaceAllString(s.Description,  "\n \t" + `✔️  SW-` )

    fmt.Println("\t" + "Patches")
    fmt.Println(s)
    }
}

