package handlers

import (
	"fmt"

	"github.com/gocolly/colly"
)

type OceanOfPDFResponse struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Image  string `json:"image_link"`
	Author string `json:"author"`
}

func OceanOfPDFHandler(searchQuery string) ([]OceanOfPDFResponse, error) {
	var response []OceanOfPDFResponse
	c := colly.NewCollector()
	c.OnHTML("article", func(e *colly.HTMLElement) {
		image_link := e.ChildAttr(".entry-image[src]", "data-src")
		author := e.ChildText(".postmetainfo")
		link := e.ChildAttr("a[href]", "href")
		title := e.ChildText(".entry-title-link")
		eachResponse := OceanOfPDFResponse{
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
	c.Visit("https://oceanofpdf.com/?s=" + searchQuery)
	return response, nil
}
