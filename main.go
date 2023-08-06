package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	// book cli based
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".mw-headline", func(e *colly.HTMLElement) {
		// Print link
		fmt.Printf("Link found: %q -> ", e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://en.wikipedia.org/wiki/William_Sterling_Parsons")
}
