package handlers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func FreeTechBooksHandler(searchQuery string) ([]string, error) {
	var response []string
	modifiedQuery := strings.Replace(searchQuery, " ", "+", -1)
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML("h2 a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Print("Book found: -> https://amazon.in" + link + " ")
		response = append(response, "https://amazon.in + link")

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(" Visiting -> ", r.URL.String())
	})

	c.Visit("https://amazon.in/s?k=" + modifiedQuery)
	fmt.Print("https://amazon.in/s?k=" + modifiedQuery)
	return response, nil
}
