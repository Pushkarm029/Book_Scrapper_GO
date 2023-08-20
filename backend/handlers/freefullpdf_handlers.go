package handlers

import (
	"fmt"

	"github.com/gocolly/colly"
)

type FreeFullPDFResponse struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Image  string `json:"image_link"`
	Author string `json:"author"`
}

// later

func FreeFullPDFHandler(searchQuery string) ([]FreeFullPDFResponse, error) {
	var response []FreeFullPDFResponse
	c := colly.NewCollector()
	c.OnHTML("article", func(e *colly.HTMLElement) {
		image_link := e.ChildAttr(".entry-image[src]", "data-src")
		author := e.ChildText(".postmetainfo")
		link := e.ChildAttr("a[href]", "href")
		title := e.ChildText(".entry-title-link")
		eachResponse := FreeFullPDFResponse{
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
	c.Visit("https://www.freefullpdf.com/search_gcse/?q=" + searchQuery)
	return response, nil
}
