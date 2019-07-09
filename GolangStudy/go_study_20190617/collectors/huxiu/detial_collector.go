package huxiu

import (
	"GolangStudy/GolangStudy/go_study_20190617/modles/huxiu"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
)

const DETAIL_URL = "https://www.huxiu.com/article/307777.html"

func HuxiuDetailSpider(conn redis.Conn) {
	startUrl := DETAIL_URL
	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()

	//解析页面新闻条目
	newsItemSelectorStr := "#article_cstaontent307777"

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	pageCollector.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})

	pageCollector.OnHTML(newsItemSelectorStr, func(e *colly.HTMLElement) {
		detail := huxiu.HuxiuDetail{}
		contents := make([]huxiu.Content, 0)
		e.DOM.Find("p").Each(func(i int, p *goquery.Selection) {
			content := huxiu.Content{}
			//遍历是否有text-remarks
			p.Find("span.text-remarks").Each(func(i int, spanTextRemarks *goquery.Selection) {
				content.ContentType = "text" //文字类型
				content.AppendContent(spanTextRemarks.Text())
				content.TextStyle = "remarks"
				return
			})

			p.Find("br").Each(func(i int, br *goquery.Selection) {
				content.ContentType = "text" //文字类型
				content.AppendContent("\n")
				content.TextStyle = "br"
				return
			})

			content.ContentType = "text" //文字类型
			content.AppendContent(p.Text())
			content.TextStyle = "remarks"

			contents = append(contents, content)
			fmt.Println(content.ContentDetail)
		})
		detail.Contents = contents

		fmt.Println(detail)
	})

	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})

	pageCollector.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	pageCollector.Visit(startUrl)
}
