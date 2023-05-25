package mac_scraping

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseRedditDevops() {
	res, err := http.Get("http://old.reddit.com/r/devops")
	if err != nil {
		if res.StatusCode == 429 {
			//Implement exponential backoff
			log.Fatal((err))
		}
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	fmt.Print(res.Header.Get("User-Agent"))

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".title").Each(func(i int, s *goquery.Selection) {
		post := s.Text()
		link, linkExists := s.Attr("href")
		macRelated := strings.EqualFold(post, "mac") || strings.EqualFold(post, "apple")
		if linkExists && macRelated {
			fmt.Printf("%s: %s\n", post, link)
		}
	})
}
