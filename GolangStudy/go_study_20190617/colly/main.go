package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {

	startUrl:="https://bookset.me/"

	//解析页面书本收集器
	pageCollector := colly.NewCollector()

	//解析下一页的收集器
	nextCollector:=colly.NewCollector()

	//解析书本详情的收集器
	detailCollector:=colly.NewCollector()


	//解析页面书本
	bookItemSelectorStr:="#cardslist > div > div.card-item"

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	pageCollector.OnHTML(bookItemSelectorStr, func(e *colly.HTMLElement) {
		a_1 := e.DOM.Find("div.thumb-img.focus>a").First()
		//h3:=e.DOM.Find("h3").First()
		a_2 := e.DOM.Find("p > a").First()

		url, _ := a_1.Attr("href")
		title, _ := a_1.Attr("title")
		author := a_2.Text()

		fmt.Printf("title: %v\t,author: %v\t,url: %v\n",
			title, author, e.Request.AbsoluteURL(url))

		detailCollector.Visit(e.Request.AbsoluteURL(url))
	})

	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
		nextCollector.Visit(startUrl)
	})

	detailCollector.OnRequest(func(r *colly.Request) {

	})

	detailDownloadSelectorStr:="#mbm-book-links1 > div > ul > li.mbm-book-download-links-listitem > a";
	detailCollector.OnHTML(detailDownloadSelectorStr, func(e *colly.HTMLElement) {
		downloadType:=e.Text
		downloadUrl:=e.Attr("href")
		fmt.Printf("%v : %v\n",downloadType,downloadUrl)
	})

	detailCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("detailCollector OnScraped")
	})


	nextCollector.OnRequest(func(r *colly.Request) {

	})

	//解析下一页
	bookNextPageSelectorStr:="body > section > div.content-wrap > div > div.pagination > ul > li.next-page > a";
	nextCollector.OnHTML(bookNextPageSelectorStr, func(e *colly.HTMLElement) {
		nextUrl:=e.Attr("href")
		startUrl=e.Request.AbsoluteURL(nextUrl)
	})

	nextCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("nextCollector OnScraped")
		pageCollector.Visit(startUrl)
	})


	// Start scraping on https://hackerspaces.org
	pageCollector.Visit(startUrl)
}
