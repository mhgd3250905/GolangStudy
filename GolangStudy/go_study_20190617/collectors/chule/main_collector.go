package chule

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/normal_news"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"regexp"
	"strconv"
	"time"
)

const KEY_CHULE_IN_REDIS = "chule"
const KEY_CHULE_INFO_IN_REDIS = "chule_info"
//const MAIN_URL = "http://www.chuapp.com/category/daily"
//const MAIN_URL = "http://www.chuapp.com/category/pcz"
//const MAIN_URL = "http://www.chuapp.com/tag/index/id/20369.html"
const MAIN_URL = "http://www.chuapp.com/category/zsyx"

var  CHULE_URL_LIST=[]string{
	"http://www.chuapp.com/category/daily",
	"http://www.chuapp.com/category/pcz",
	"http://www.chuapp.com/tag/index/id/20369.html",
	"http://www.chuapp.com/category/zsyx",
	}

var page = 1
var lastDateline = ""

var index=0

func ChuleSpider(conn redis.Conn) {
	startUrl := CHULE_URL_LIST[index]

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()
	//解析下一页收集器

	//pageCollector.SetProxyFunc(rp)
	//nextCollector.SetProxyFunc(rp)

	//解析页面新闻条目
	newsItemSelectorStr := "body > div.content.category.fn-clear > div.category-left.fn-left > div > a"

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
		//解析新闻条目
		e.DOM.Each(func(i int, a *goquery.Selection) {
			news := normal_news.News{}

			link, exist := a.Attr("href")
			if !exist {
				return
			}

			// /article/309213.html
			re, _ := regexp.Compile(`[0-9]+`)
			all := re.FindAll([]byte(link), 1)
			newsId := ""
			for i, _ := range all {
				newsId = string(all[i])
			}

			title, exist := a.Attr("title")
			if !exist {
				return
			}
			img := a.Find("img").First()
			imgLink, exist := img.Attr("src")
			if !exist {
				return
			}
			dd_desc := a.Find("dl > dd:nth-child(3)").First()
			desc := dd_desc.Text()
			em_author := a.Find("dl > dd > span.fn-left > em").First()
			authorName := em_author.Text()
			span_time := a.Find("dl>dd>span.fn-left")
			timeStr := span_time.Text()
			re, _ = regexp.Compile(`[0-9]+月/[0-9]+日`)
			all = re.FindAll([]byte(timeStr), 1)
			for i, _ := range all {
				timeStr = string(all[i])
			}

			timeTemplate1 := "2006年01月02日" //常规类型
			////timeTemplate2 := "2006/01/02 15:04:05" //其他类型
			timeTemplate3 := "2006年" //其他类型
			////timeTemplate4 := "15:04:05"            //其他类型
			stamp, _ := time.ParseInLocation(timeTemplate1,
				fmt.Sprintf("%s%s", time.Now().Format(timeTemplate3),
					timeStr[len(timeStr)-10:]), time.Local)
			fmt.Println(stamp)

			timeStr = strconv.FormatInt(stamp.Unix(), 10)

			news.Author = normal_news.Author{
				AuthorName: authorName,
				AuthorId:   "",
				AuthorImg:  "",
			}
			//normal_news.Categorys = nil
			news.CreateTime = timeStr
			news.Desc = desc
			news.ImgLink = imgLink
			news.Title = title
			news.NewsLink = e.Request.AbsoluteURL(link)
			news.NewsId = newsId

			jsonBytes, err := json.Marshal(&news)
			if err != nil {
				fmt.Printf("%v json.Marshal failed,err= %v\n", title, err)
				return
			}

			err = redis_utils.Push2RedisSortedSet(conn,news.NewsId, KEY_CHULE_IN_REDIS,KEY_CHULE_INFO_IN_REDIS, news.CreateTime, string(jsonBytes))
			if err != nil {
				fmt.Printf("%v push2RedisList failed,err= %v\n", title, err)
				return
			} else {
				fmt.Printf("%s 爬取完毕\n", news.Title)
				ChuleDetailSpider(conn, news)
			}

		})

	})

	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
		index++
		if index > 3 {
			return
		}
		startUrl=CHULE_URL_LIST[index]
		pageCollector.Visit(startUrl)
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})

	pageCollector.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	pageCollector.Visit(startUrl)
}


