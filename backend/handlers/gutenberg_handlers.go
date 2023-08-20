package handlers

import (
	"fmt"

	"github.com/gocolly/colly"
)

type GutenbergResponse struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Image  string `json:"image_link"`
	Author string `json:"author"`
}

func GutenbergHandler(searchQuery string) ([]GutenbergResponse, error) {
	var response []GutenbergResponse
	c := colly.NewCollector()
	c.OnHTML(".booklink", func(e *colly.HTMLElement) {
		image_link := "https://www.gutenberg.org" + e.ChildAttr("img", "src")
		author := e.ChildText(".subtitle")
		link := "https://www.gutenberg.org" + e.ChildAttr("a[href]", "href")
		title := e.ChildText(".title")
		eachResponse := GutenbergResponse{
			Title:  title,
			Link:   link,
			Image:  image_link,
			Author: author,
		}
		response = append(response, eachResponse)

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(" Visiting -> ", r.URL.String())
	})
	c.Visit("https://www.gutenberg.org/ebooks/search/?query=" + searchQuery)
	return response, nil
}
