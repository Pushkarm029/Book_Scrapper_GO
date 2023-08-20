package handlers

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

type ArchiveResponse struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Image  string `json:"image_link"`
	Author string `json:"author"`
}

//not working

func ArchiveHandler(searchQuery string) ([]ArchiveResponse, error) {
	var response []ArchiveResponse
	var wg sync.WaitGroup

	var responseCount int
	c := colly.NewCollector()
	c.OnHTML("article", func(e *colly.HTMLElement) {
		responseCount++
		if responseCount >= 20 {
			return // If we've reached the limit, stop appending
		}

		image_link := e.ChildAttr(".contain[src]", "src")
		author := e.ChildAttr("span[title].truncated", "title")
		link := e.ChildAttr("a[href]", "href")
		title := e.ChildAttr("h4.truncated", "title")
		// .item-info
		eachResponse := ArchiveResponse{
			Title:  title,
			Link:   link,
			Image:  image_link,
			Author: author,
		}
		fmt.Print(eachResponse)
		response = append(response, eachResponse)
		wg.Done()
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(" Visiting -> ", r.URL.String())
		wg.Add(1)
	})
	c.OnScraped(func(r *colly.Response) {
		wg.Wait() // Wait until all scraping is done
	})
	c.Visit("https://archive.org/search?query=" + searchQuery)
	return response, nil
}
