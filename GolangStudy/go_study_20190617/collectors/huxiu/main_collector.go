package huxiu

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/huxiu"
	"GolangStudy/GolangStudy/go_study_20190617/modles/normal_news"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const KEY_HUXIU_IN_REDIS = "huxiu"
const KEY_HUXIU_DETIAL_IN_REDIS = "huxiu_detail"
const MAIN_URL = "https://wwww.huxiu.com/"

var page = 1
var lastDateline = ""

func HuxiuSpider(conn redis.Conn,onSpiderFinish func()) {
	startUrl := MAIN_URL

	////使用代理
	//rp, err := proxy.RoundRobinProxySwitcher("http://123.52.19.47:3128",
	//	"https://218.91.112.188:9999",
	//	"https://1.192.243.9:9999",
	//	"https://58.253.157.205",
	//	"https://180.119.141.211",
	//	"https://120.83.107.7",
	//	"https://1.192.242.122:9999")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()
	//解析下一页收集器
	nextCollector := colly.NewCollector()

	//pageCollector.SetProxyFunc(rp)
	//nextCollector.SetProxyFunc(rp)

	//解析页面新闻条目
	newsItemSelectorStr := "#index > div.wrap-left.pull-left"

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
		e.DOM.Find("div.mod-info-flow > div.mod-b.mod-art.clearfix").Each(func(i int, selection *goquery.Selection) {
			a_title := selection.Find("div.mob-ctt.index-article-list-yh > h2 > a").First()
			img_author := selection.Find("div.mob-ctt.index-article-list-yh > div.mob-author > div > a > img").First()
			a_author := selection.Find("div.mob-ctt.index-article-list-yh > div.mob-author > a:nth-child(2)").First()
			div_desc := selection.Find("div.mob-ctt.index-article-list-yh > div.mob-sub").First()
			img_imglink := selection.Find("div.mod-thumb.pull-left > a img").First()
			//可能包含多个分类信息
			div_category := selection.Find("div.column-link-box >a.column-link")

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

			author := normal_news.Author{}

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

			timeTemplate1 := "200601/02/150405" //常规类型
			//timeTemplate2 := "2006/01/02 15:04:05" //其他类型
			//timeTemplate3 := "2006-01-02"          //其他类型
			//timeTemplate4 := "15:04:05"            //其他类型

			stamp, _ := time.ParseInLocation(timeTemplate1, timeStr[:16], time.Local)

			//仅仅获取最近三天的数据
			threeDaysBefore:=time.Now().AddDate(0,0,-3)
			if stamp.Before(threeDaysBefore) {
				return
			}

			//fmt.Println(stamp)

			timeStr = strconv.FormatInt(stamp.Unix(), 10)

			categorys := make([]normal_news.Category, 0)
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
				category := normal_news.Category{}
				category.CategoryId = categoryId
				category.CategoryName = categoryName
				categorys = append(categorys, category)
			})

			// /article/309213.html
			re, _ = regexp.Compile(`[0-9]+`)
			all = re.FindAll([]byte(newsLink), 1)
			newsId := ""
			for i, _ := range all {
				newsId = string(all[i])
			}

			news := normal_news.News{
				NewsId:     newsId,
				Title:      title,
				NewsLink:   e.Request.AbsoluteURL(newsLink),
				Author:     author,
				CreateTime: timeStr,
				Desc:       desc,
				ImgLink:    imgLink,
				Categorys:  categorys,
			}

			jsonBytes, err := json.Marshal(&news)
			if err != nil {
				fmt.Printf("%v json.Marshal failed,err= %v\n", title, err)
				return
			}

			err = redis_utils.Push2RedisSortedSet(conn, KEY_HUXIU_IN_REDIS, news.CreateTime, string(jsonBytes))
			if err != nil {
				fmt.Printf("%v push2RedisList failed,err= %v\n", title, err)
				return
			} else {
				fmt.Printf("%s 爬取完毕\n", news.Title)
				HuxiuDetailSpider(conn, news)
			}
		})

		div_Next := e.DOM.Find("div.get-mod-more.js-get-mod-more-list.transition").First()

		pageStr, _ := div_Next.Attr("data-cur_page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return
		}
		lastDateline, _ = div_Next.Attr("data-last_dateline")

		visitNextPage(page, lastDateline, nextCollector)
	})

	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})

	nextCollector.OnRequest(func(r *colly.Request) {

	})

	nextCollector.OnResponse(func(response *colly.Response) {
		body := response.Body
		huxiuNext := huxiu.HuxiuNext{}
		err := json.Unmarshal(body, &huxiuNext)
		if err != nil {
			return
		}
		lastDateline = huxiuNext.LastDateline
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(huxiuNext.Data))
		if err != nil {
			return
		}
		doc.Find("div.mod-b.mod-art").Each(func(i int, selector *goquery.Selection) {
			a_title := selector.Find("div.mob-ctt > h2 > a").First()
			img_author := selector.Find("div.mob-ctt > div.mob-author > div > a > img").First()
			a_author := selector.Find("div.mob-ctt > div.mob-author > a:nth-child(2)").First()
			div_desc := selector.Find("div.mob-ctt > div.mob-sub").First()
			img_imglink := selector.Find("div.mod-thumb > a > img").First()
			//可能包含多个分类信息
			div_category := selector.Find("div.column-link-box >a.column-link")

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

			author := normal_news.Author{}

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

			timeTemplate1 := "200601/02/150405" //常规类型
			//timeTemplate2 := "2006/01/02 15:04:05" //其他类型
			//timeTemplate3 := "2006-01-02"          //其他类型
			//timeTemplate4 := "15:04:05"            //其他类型

			stamp, _ := time.ParseInLocation(timeTemplate1, timeStr[:16], time.Local)

			//fmt.Println(stamp)

			//仅仅获取最近三天的数据
			threeDaysBefore:=time.Now().AddDate(0,0,-3)
			if stamp.Before(threeDaysBefore) {
				fmt.Println("三天内文章爬取完毕！")
				onSpiderFinish()
			}

			timeStr = strconv.FormatInt(stamp.Unix(), 10)

			categorys := make([]normal_news.Category, 0)
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
				category := normal_news.Category{}
				category.CategoryId = categoryId
				category.CategoryName = categoryName
				categorys = append(categorys, category)
			})

			// /article/309213.html
			re, _ = regexp.Compile(`[0-9]+`)
			all = re.FindAll([]byte(newsLink), 1)
			newsId := ""
			for i, _ := range all {
				newsId = string(all[i])
			}

			news := normal_news.News{
				NewsId:     newsId,
				Title:      title,
				NewsLink:   response.Request.AbsoluteURL(newsLink),
				Author:     author,
				CreateTime: timeStr,
				Desc:       desc,
				ImgLink:    imgLink,
				Categorys:  categorys,
			}
			//fmt.Println(normal_news.Title)

			jsonBytes, err := json.Marshal(&news)
			if err != nil {
				fmt.Printf("%v json.Marshal failed,err= %v\n", title, err)
				return
			}

			err = redis_utils.Push2RedisSortedSet(conn, KEY_HUXIU_IN_REDIS, news.CreateTime, string(jsonBytes))
			if err != nil {
				fmt.Printf("%v push2RedisList failed,err= %v\n", title, err)
				return
			} else {
				fmt.Printf("%s 爬取完毕\n", news.Title)
				HuxiuDetailSpider(conn, news)
			}
		})

		page = page + 1
		visitNextPage(page, lastDateline, nextCollector)

	})

	nextCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("nextCollector.OnError: ", e)
	})

	pageCollector.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	pageCollector.Visit(startUrl)
}

func visitNextPage(curPage int, lastDateline string, nextCollector *colly.Collector) {
	fmt.Printf("Page： %d , lastDateline:%s----------------------------------------------------Visit\n", curPage, lastDateline)
	formData := make(map[string]string)
	formData["page"] = strconv.Itoa(curPage)
	formData["last_dateline"] = lastDateline
	fmt.Println()
	nextCollector.Post("https://www.huxiu.com/v2_action/article_list", formData)
}
