package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	var searchQuery string
	fmt.Print("Type Book You Want to Search -> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		searchQuery = scanner.Text()
	}
	modifiedQuery := strings.Replace(searchQuery, " ", "+", -1)
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML("h2 a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Print("Book found: -> https://amazon.in" + link + " ")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(" Visiting -> ", r.URL.String())
	})

	c.Visit("https://amazon.in/s?k=" + modifiedQuery)
	fmt.Print("https://amazon.in/s?k=" + modifiedQuery)
}
