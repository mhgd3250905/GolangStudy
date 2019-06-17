package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	c.AllowedDomains = []string{"bookset.me"}

	c.OnHTML("a[href][title]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		title := e.Attr("title")
		fmt.Printf("Link found: %q -> %s\n", title, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})


	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})



	// Start scraping on https://hackerspaces.org
	c.Visit("https://bookset.me/")
}
