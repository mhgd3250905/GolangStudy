package bookset

import (
	"GolangStudy/GolangStudy/go_study_20190617/modles/bookSet"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
)

const KEY_BOOK_IN_REDIS = "book"


func BooksetSpider(conn redis.Conn) {
	startUrl := "https://bookset.me/"
	//解析页面书本收集器
	pageCollector := colly.NewCollector()
	//解析下一页的收集器
	nextCollector := colly.NewCollector()
	//解析页面书本
	bookItemSelectorStr := "#cardslist > div > div.card-item"
	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	pageCollector.OnHTML(bookItemSelectorStr, func(e *colly.HTMLElement) {
		a_1 := e.DOM.Find("div.thumb-img.focus>a").First()
		img_1 := a_1.Find("img.thumb").First()
		a_2 := e.DOM.Find("p > a").First()

		url, _ := a_1.Attr("href")
		title, _ := a_1.Attr("title")
		author := a_2.Text()
		authorLink, _ := a_2.Attr("href")
		imgPath, _ := img_1.Attr("src")

		book := bookSet.Book{
			Title:      title,
			Author:     author,
			BookLink:   e.Request.AbsoluteURL(url),
			AuthorLink: e.Request.AbsoluteURL(authorLink),
			Image:      imgPath,
		}

		//fmt.Println(book)

		jsonBytes, err := json.Marshal(&book)
		if err != nil {
			fmt.Printf("%v json.Marshal failed,err= %v\n", title, err)
			return
		}

		err = push2RedisList(conn, KEY_BOOK_IN_REDIS, string(jsonBytes))
		if err != nil {
			fmt.Printf("%v push2RedisList failed,err= %v\n", title, err)
			return
		}

		fmt.Printf("%v save success!\n", title)

		//继续访问内部
		GetDetailCollector(conn).Visit(e.Request.AbsoluteURL(url))
	})
	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
		nextCollector.Visit(startUrl)
	})
	nextCollector.OnRequest(func(r *colly.Request) {

	})
	//解析下一页
	bookNextPageSelectorStr := "body > section > div.content-wrap > div > div.pagination > ul > li.next-page > a"
	nextCollector.OnHTML(bookNextPageSelectorStr, func(e *colly.HTMLElement) {
		nextUrl := e.Attr("href")
		startUrl = e.Request.AbsoluteURL(nextUrl)
	})
	nextCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("nextCollector OnScraped")
		pageCollector.Visit(startUrl)
	})
	// Start scraping on https://hackerspaces.org
	pageCollector.Visit(startUrl)
}