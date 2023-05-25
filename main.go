package main

import (
	mac_scraping "github.com/cjflan/mac-web-scraper/scrapers"
)

func main() {
	mac_scraping.ParseRedditDevops()
	mac_scraping.ParseLobsters()
	mac_scraping.ParseHackerNews()
}
