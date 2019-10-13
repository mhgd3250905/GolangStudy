package ifanr

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/ifanr"
	"GolangStudy/GolangStudy/go_study_20190617/modles/normal_news"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

const KEY_IFANR_REDIS = "ifanr"
const KEY_IFANR_INFO_REDIS = "ifanr_info"
const MAIN_URL = "https://www.ifanr.com/"

const JSON_URL = "https://sso.ifanr.com//api/v5/wp/web-feed/?limit=%d&offset=%d"

var page = 0
var count = 20

var index = 0

func IfanrSpider(conn redis.Conn,onSpiderFinish func()) {
	startUrl := fmt.Sprintf(JSON_URL, count, page*count)

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()
	//解析下一页收集器

	//pageCollector.SetProxyFunc(rp)
	//nextCollector.SetProxyFunc(rp)

	//解析页面新闻条目
	//newsItemSelectorStr := "#collectionList > .article-item.article-item--list"

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	pageCollector.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
		responseIfanr := ifanr.ResponseIfanr{}
		err := json.Unmarshal(response.Body, &responseIfanr)
		if err != nil {
			fmt.Println("ifaner unmarshal failed,err=", err)
			return
		}
		if len(responseIfanr.Objects) <= 0 {
			fmt.Println("response is empty!")
			return
		}

		news := normal_news.News{}
		for i, _ := range responseIfanr.Objects {
			item := responseIfanr.Objects[i]
			news.Author = normal_news.Author{
				AuthorName: item.CreatedBy.Name,
				AuthorId:   strconv.Itoa(item.CreatedBy.Id),
				AuthorImg:  item.CreatedBy.Avatar,
			}

			timeTemplate1 := "2006-01-02 15:04:05" //常规类型
			////timeTemplate2 := "2006/01/02 15:04:05" //其他类型
			//timeTemplate3 := "2006年" //其他类型
			////timeTemplate4 := "15:04:05"            //其他类型
			stamp, _ := time.ParseInLocation(timeTemplate1, item.CreatedAtFormat, time.Local)

			//仅仅获取最近三天的数据
			threeDaysBefore:=time.Now().AddDate(0,0,-3)
			if stamp.Before(threeDaysBefore) {
				fmt.Println("三天内文章爬取完毕！")
				onSpiderFinish()
			}

			timeStr := strconv.FormatInt(stamp.Unix(), 10)

			//normal_news.Categorys = nil
			news.CreateTime = timeStr
			news.Desc = item.PostExcerpt
			news.ImgLink = item.PostCoverImage
			news.Title = item.PostTitle
			news.NewsLink = item.PostUrl
			news.NewsId = item.PostId

			jsonBytes, err := json.Marshal(&news)
			if err != nil {
				fmt.Printf("%v json.Marshal failed,err= %v\n", news.Title, err)
				return
			}

			err = redis_utils.Push2RedisSortedSet(conn,news.NewsId, KEY_IFANR_REDIS,KEY_IFANR_INFO_REDIS, news.CreateTime, string(jsonBytes))
			if err != nil {
				fmt.Printf("%v push2RedisList failed,err= %v\n", news.Title, err)
				return
			} else {
				fmt.Printf("%s 爬取完毕\n", news.Title)
				IfanrDetailSpider(conn, news)
			}
		}
	})


	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})

	pageCollector.Visit(startUrl)
}
