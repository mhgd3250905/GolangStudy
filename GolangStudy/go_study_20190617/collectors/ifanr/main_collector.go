package ifanr

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/huxiu"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"regexp"
	"strconv"
	"time"
)

const KEY_CHULE_IN_IFANR = "ifanr"
const MAIN_URL = "https://www.ifanr.com/"

const JOSN_URL="https://sso.ifanr.com//api/v5/wp/web-feed/?published_at__lte=2019-08-04+15%3A24%3A03&limit=20&offset=0"


var page = 1
var lastDateline = ""

var index=0

func ChuleSpider(conn redis.Conn) {
	startUrl := MAIN_URL

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()
	//解析下一页收集器

	//pageCollector.SetProxyFunc(rp)
	//nextCollector.SetProxyFunc(rp)

	//解析页面新闻条目
	newsItemSelectorStr := "#collectionList > .article-item.article-item--list"

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
		e.DOM.Each(func(i int, item *goquery.Selection) {
			news := huxiu.HuxiuNews{}

			a_image:=item.Find("div.article-image.cover-image > a.article-link.cover-block").First()
			//a_category:=item.Find("div.article-image.cover-image > a.article-label").First()
			a_title:=item.Find("div.article-info.js-transform > h3 > a").First()
			div_author:=item.Find("div.article-info.js-transform > div.article-meta > div.author-info").First()
			span_authorName:=div_author.Find("span").First()
			img_author:=div_author.Find("img").First()
			div_desc:=item.Find("div.article-info.js-transform > div.article-summary").First()
			time_createTime:=item.Find("div.article-info.js-transform > div.article-meta > time").First()



			link, exist := a_image.Attr("href")
			if !exist {
				return
			}

			re, _ := regexp.Compile(`[0-9]+`)
			all := re.FindAll([]byte(link), 1)
			newsId := ""
			for i, _ := range all {
				newsId = string(all[i])
			}

			styleStr, exist := a_image.Attr("style")
			if !exist {
				return
			}

			re, _ = regexp.Compile(`https://.+!260`)
			all = re.FindAll([]byte(styleStr), 1)
			imgLink:=""
			for i, _ := range all {
				imgLink = string(all[i])
			}


			title := a_title.Text()

			desc := div_desc.Text()

			authorImage,exist:=img_author.Attr("src")
			if !exist{
				return
			}

			authorName := span_authorName.Text()

			timeStr,exist:=time_createTime.Attr("data-time")
			if !exist{
				return
			}

			timeTemplate1 := "2006-01-02 15:04:05" //常规类型
			////timeTemplate2 := "2006/01/02 15:04:05" //其他类型
			//timeTemplate3 := "2006年" //其他类型
			////timeTemplate4 := "15:04:05"            //其他类型
			stamp, _ := time.ParseInLocation(timeTemplate1,timeStr, time.Local)
			//fmt.Println(stamp)

			timeStr = strconv.FormatInt(stamp.Unix(), 10)

			news.Author = huxiu.Author{
				AuthorName: authorName,
				AuthorId:   "",
				AuthorImg:  authorImage,
			}

			//news.Categorys = nil
			news.CreateTime = timeStr
			news.Desc = desc
			news.ImgLink = imgLink
			news.Title = title
			news.NewsLink = link
			news.NewsId = newsId

			jsonBytes, err := json.Marshal(&news)
			if err != nil {
				fmt.Printf("%v json.Marshal failed,err= %v\n", title, err)
				return
			}

			err = redis_utils.Push2RedisSortedSet(conn, KEY_CHULE_IN_IFANR, news.CreateTime, string(jsonBytes))
			if err != nil {
				fmt.Printf("%v push2RedisList failed,err= %v\n", title, err)
				return
			} else {
				fmt.Printf("%s 爬取完毕\n", news.Title)
				//ChuleDetailSpider(conn, news)
			}

		})

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


