package handlers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type AmazonResponse struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Image    string `json:"image_link"`
	Author   string `json:"author"`
	Language string `json:"language"`
	Price    string `json:"price"`
}

// maybe use a official api

func AmazonHandler(searchQuery string) ([]AmazonResponse, error) {
	var response []AmazonResponse
	class := "s-card-container s-overflow-hidden aok-relative puis-wide-grid-style puis-wide-grid-style-t1 puis-include-content-margin puis puis-v3b48cl1js792724v4d69zlbwph s-latency-cf-section s-card-border"
	modifiedQuery := strings.Replace(searchQuery, " ", "+", -1)
	modifiedClass := strings.Replace(class, " ", ".", -1)
	c := colly.NewCollector()
	c.OnHTML(modifiedClass, func(e *colly.HTMLElement) {
		image_link := e.ChildAttr("img.s-image", "src")
		author := e.ChildText(".sg-col-inner.a-row.a")
		language := e.ChildText(".sg-col-inner.a-row.a.a-size-base.a-color-secondary")
		price := e.ChildText(".a-price-whole")
		link := "https://amazon.in" + e.ChildAttr(".a-link-normal.s-no-outline", "href")
		title := e.ChildText(".a-size-mini.a-spacing-none.a-color-base.s-line-clamp-2.a-size-medium.a-color-base.a-text-normal")
		eachResponse := AmazonResponse{
			Title:    title,
			Link:     link,
			Image:    image_link,
			Author:   author,
			Language: language,
			Price:    price,
		}
		response = append(response, eachResponse)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(" Visiting -> ", r.URL.String())
	})

	c.Visit("https://amazon.in/s?k=" + modifiedQuery)
	fmt.Print("https://amazon.in/s?k=" + modifiedQuery)
	return response, nil
}
