package pentaq

import (
	"GolangStudy/GolangStudy/go_study_20190617/modles/pentaq"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
)

const KEY_PENTAQ_REDIS = "pentaq"
const KEY_IFANR_INFO_REDIS = "pentaq_info"


const JSON_URL="http://www.pentaq.com/wp-json/wp/v2/posts?_embed=1&page=%d&per_page=%d"

var page = 1
var count = 20

var index = 0

func PentaqSeider(conn redis.Conn,onSpiderFinish func()) {
	startUrl := fmt.Sprintf(JSON_URL, page,count)

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()
	//解析下一页收集器

	//pageCollector.SetProxyFunc(rp)
	//nextCollector.SetProxyFunc(rp)

	//解析页面新闻条目
	//newsItemSelectorStr := "#collectionList > .article-item.article-item--list"


	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	pageCollector.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
		var esponsePentaq []pentaq.PentaqItem
		err := json.Unmarshal(response.Body, &esponsePentaq)
		fmt.Println(string(response.Body))
		if err != nil {
			fmt.Println("ifaner unmarshal failed,err=", err)
			return
		}
		if len(esponsePentaq) <= 0 {
			fmt.Println("response is empty!")
			return
		}

		//news := normal_news.News{}
		//for i, _ := range responseIfanr.Objects {
		//	item := responseIfanr.Objects[i]
		//	news.Author = normal_news.Author{
		//		AuthorName: item.CreatedBy.Name,
		//		AuthorId:   strconv.Itoa(item.CreatedBy.Id),
		//		AuthorImg:  item.CreatedBy.Avatar,
		//	}
		//
		//	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
		//	////timeTemplate2 := "2006/01/02 15:04:05" //其他类型
		//	//timeTemplate3 := "2006年" //其他类型
		//	////timeTemplate4 := "15:04:05"            //其他类型
		//	stamp, _ := time.ParseInLocation(timeTemplate1, item.CreatedAtFormat, time.Local)
		//
		//	//仅仅获取最近三天的数据
		//	threeDaysBefore:=time.Now().AddDate(0,0,-3)
		//	if stamp.Before(threeDaysBefore) {
		//		fmt.Println("三天内文章爬取完毕！")
		//		onSpiderFinish()
		//	}
		//
		//	timeStr := strconv.FormatInt(stamp.Unix(), 10)
		//
		//	//normal_news.Categorys = nil
		//	news.CreateTime = timeStr
		//	news.Desc = item.PostExcerpt
		//	news.ImgLink = item.PostCoverImage
		//	news.Title = item.PostTitle
		//	news.NewsLink = item.PostUrl
		//	news.NewsId = item.PostId
		//
		//	jsonBytes, err := json.Marshal(&news)
		//	if err != nil {
		//		fmt.Printf("%v json.Marshal failed,err= %v\n", news.Title, err)
		//		return
		//	}
		//
		//	err = redis_utils.Push2RedisSortedSet(conn,news.NewsId, KEY_IFANR_REDIS,KEY_IFANR_INFO_REDIS, news.CreateTime, string(jsonBytes))
		//	if err != nil {
		//		fmt.Printf("%v push2RedisList failed,err= %v\n", news.Title, err)
		//		return
		//	} else {
		//		fmt.Printf("%s 爬取完毕\n", news.Title)
		//		IfanrDetailSpider(conn, news)
		//	}
		//}
	})

	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
		page++
		startUrl := fmt.Sprintf(JSON_URL, count, page*count)
		pageCollector.Visit(startUrl)
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		r.Headers.Set("Connection","keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests","1")
		r.Headers.Set("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
		r.Headers.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
		r.Headers.Set("Accept-Encoding","gzip, deflate")
		r.Headers.Set("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8")
		r.Headers.Set("Cookie","PHPSESSID=8d18ps742m3p3bo4324ak2q5e9")
	})


	//pageCollector.Cookies("PHPSESSID=8d18ps742m3p3bo4324ak2q5e9")

	pageCollector.Visit(startUrl)
}
