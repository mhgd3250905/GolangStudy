package huxiu

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/huxiu"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"regexp"
)

const KEY_HUXIU_IN_REDIS = "huxiu"
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

		author := huxiu.Author{}

		// /member/1854035.html 查找AuthorId
		re, _ := regexp.Compile("[0-9]+")
		all := re.FindAll([]byte(authorLink), 1)
		for i, _ := range all {
			author.AuthorId = string(all[i])
		}

		authorName := a_author.Find("span").First().Text()
		authorImg, exist := img_author.Attr("src")
		if !exist {
			return
		}
		author.AuthorName = authorName
		author.AuthorImg = authorImg

		desc := div_desc.Text()

		imgLink, exist := img_imglink.Attr("data-original")
		if !exist {
			return
		}

		//https://img.huxiucdn.com/article/cover/201907/01/163823923163.jpg?...
		re, _ = regexp.Compile(`[0-9]+/[0-9]+/[0-9]+`)
		all = re.FindAll([]byte(imgLink), 1)
		timeStr := ""
		for i, _ := range all {
			timeStr = string(all[i])
		}
		fmt.Println("time:", timeStr)

		categorys := make([]huxiu.Category, 0)
		div_category.Each(func(i int, selection *goquery.Selection) {
			categoryName := selection.Text()
			categoryLink, exist := selection.Attr("href")
			if !exist {
				return
			}
			re, _ = regexp.Compile(`[0-9]+`)
			all = re.FindAll([]byte(categoryLink), 1)
			categoryId := ""
			for i, _ := range all {
				categoryId = string(all[i])
			}
			category := huxiu.Category{}
			category.CategoryId = categoryId
			category.CategoryName = categoryName
			categorys = append(categorys, category)
		})

		news := huxiu.HuxiuNews{
			Title:      title,
			NewsLink:   e.Request.AbsoluteURL(newsLink),
			Author:     author,
			CreateTime: timeStr,
			Desc:       desc,
			ImgLink:    imgLink,
			Categorys:  categorys,
		}
		fmt.Println(news)

		jsonBytes, err := json.Marshal(&news)
		if err != nil {
			fmt.Printf("%v json.Marshal failed,err= %v\n", title, err)
			return
		}

		err = redis_utils.Push2RedisList(conn, KEY_HUXIU_IN_REDIS, string(jsonBytes))
		if err != nil {
			fmt.Printf("%v push2RedisList failed,err= %v\n", title, err)
			return
		}
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
		//pageCollector.Visit(startUrl)
	})
	// Start scraping on https://hackerspaces.org
	pageCollector.Visit(startUrl)
}
