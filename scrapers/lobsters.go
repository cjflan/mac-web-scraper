package mac_scraping

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseLobsters() {
	res, err := http.Get("http://lobste.rs")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".u-url").Each(func(i int, s *goquery.Selection) {
		post := s.Text()
		link, linkExists := s.Attr("href")
		macRelated := strings.Contains(strings.ToLower(post), "mac") || strings.Contains(strings.ToLower(post), "apple")
		if linkExists && macRelated {
			fmt.Printf("%s: %s\n", post, link)
		}
	})
}
