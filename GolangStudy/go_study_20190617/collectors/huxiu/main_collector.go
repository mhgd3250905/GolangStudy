package huxiu

import (
	"GolangStudy/GolangStudy/go_study_20190617/modles/huxiu"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"regexp"
)

const KEY_BOOK_IN_REDIS = "huxiu"
const MAIN_URL = "https://www.huxiu.com/"

func HuxiuSpider(conn redis.Conn) {
	startUrl := MAIN_URL
	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()
	//解析下一页收集器
	nextCollector := colly.NewCollector()

	//解析页面新闻条目
	newsItemSelectorStr := "#index > div.wrap-left.pull-left > div.mod-info-flow > div.mod-b.mod-art.clearfix"

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	/**
	type HuxiuNews struct {
		Title      string   `json:"title"`
		NewsLink   string   `json:"news_link"`
		Author     Author   `json:"author"`
		CreateTime string   `json:"create_time"`
		Desc       string   `json:"desc"`
		ImgLink    string   `json:"image_link"`
		category   Category `json:"category"`
	}
	type Author struct {
		AuthorName string `json:"author_name"`
		AuthorId int `json:"author_id"`
	}
	type Category struct {
		categoryName string `json:"category_name"`
		categoryId   int    `json:"category_id"`
	}
	*/
	pageCollector.OnHTML(newsItemSelectorStr, func(e *colly.HTMLElement) {
		a_title := e.DOM.Find("div.mob-ctt.index-article-list-yh > h2 > a").First()
		img_author := e.DOM.Find("div.mob-ctt.index-article-list-yh > div.mob-author > div > a > img").First()
		a_author := e.DOM.Find("div.mob-ctt.index-article-list-yh > div.mob-author > a:nth-child(2)").First()
		div_desc := e.DOM.Find("div.mob-ctt.index-article-list-yh > div.mob-sub").First()
		img_imglink := e.DOM.Find("div.mod-thumb.pull-left > a > img").First()
		//可能包含多个分类信息
		div_category := e.DOM.Find("div.column-link-box >a.column-link")

		//解析结果
		title := a_title.Text()
		newsLink, exist := a_title.Attr("href")
		if !exist {
			return
		}
		authorLink, exist := a_author.Attr("href")
		if !exist {
			return
		}

		author:=huxiu.Author{}

		// /member/1854035.html 查找AuthorId
		re,_:=regexp.Compile("[0-9]+")
		all:=re.FindAll([]byte(authorLink),1)
		for i,_ := range all {
			author.AuthorId=string(all[i])
		}
		

		authorName := a_author.Find("span").First().Text()
		authorImg, exist := img_author.Attr("src")
		if !exist {
			return
		}
		//todo 解析时间

		desc := div_desc.Text()
		imgLink, exist := img_imglink.Attr("data-original")
		if !exist {
			return
		}

		categorys:=make([]huxiu.Category,2)
		div_category.Each(func(i int, selection *goquery.Selection) {
			categoryName:=selection.Text()
			//categoryLink,exist:=selection.Attr("href")
			//if !exist {
			//	return
			//}
			category:=huxiu.Category{}
			category.CategoryName=categoryName
			categorys=append(categorys, category)
		})
		fmt.Printf("title: %v\n link: %v\n authorLink:%v\n authorName:%v\n authorImg:%v\n" +
			" desc: %v\n imgLink:%v\n categorys:%v\n",
			title, newsLink,authorLink,authorName,authorImg,desc,imgLink,categorys)

	})
	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
		//nextCollector.Visit(startUrl)
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
		//pageCollector.Visit(startUrl)
	})
	// Start scraping on https://hackerspaces.org
	pageCollector.Visit(startUrl)
}
