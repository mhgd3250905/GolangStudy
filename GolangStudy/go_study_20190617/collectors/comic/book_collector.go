package comic

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/comic"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"regexp"
	"strconv"
	"time"
)

/**
漫画的存储数据结构
- 漫画id name
	- 章节*n
		- 章节中具体的图片*n
 */

//太古漫画 鬼灭之刃地址
const MAIN_URL = "https://www.dagumanhua.com/manhua/3629/"

const KEY_COMIC_BOOK_ID_IN_REDIS  = "COMIC_BOOK_ID"
const KEY_COMIC_BOOK_INFO_IN_REDIS  = "COMIC_BOOK_INFO"
const KEY_COMIC_CHAPTER_LIST_IN_REDIS  = "COMIC_CHAPTER_LIST"
const KEY_COMIC_CHAPTER_DETAIL_IN_REDIS  = "COMIC_CHAPTER_DETAIL"

func ComicSpider(conn redis.Conn, onSpiderFinish func()) {

	book:=comic.ComicBook{}

	startUrl := MAIN_URL

	//解析网页漫画信息收集器
	pageCollector := colly.NewCollector()


	//解析漫画信息容器
	bodySelectorStr := "body"
	itemSelectorStr := "#play_0 > ul > li"

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	pageCollector.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})

	pageCollector.OnHTML(bodySelectorStr, func(e *colly.HTMLElement) {


		imageLink,exist:=e.DOM.Find("#intro_l > div.cy_info_cover > img").Attr("src")
		if !exist {
			fmt.Println("图片链接不存在！")
		}

		name,exist:=e.DOM.Find("#intro_l > div.cy_info_cover > img").Attr("alt")
		if !exist {
			fmt.Println("漫画名称不存在！")
		}

		desc:=e.DOM.Find("#comic-description").Text()


		chapters:=make([]comic.Chapter,0)
		e.DOM.Find(itemSelectorStr).Each(func(i int, li *goquery.Selection) {
			chapter:=comic.Chapter{}

			title,exist:=li.Find("a").First().Attr("title")
			if !exist {
				fmt.Println("不存在此标题！")
			}
			url,exist:=li.Find("a").First().Attr("href")
			if !exist {
				fmt.Println("不存在此链接！")
			}

			// /manhua/3629/608970.html
			re, _ := regexp.Compile(`[0-9]+`)
			all := re.FindAll([]byte(url), 2)
			chapterId := ""
			for i, _ := range all {
				chapterId = string(all[i])
			}

			//fmt.Printf("标题 %s,链接%s\n",title,e.Request.AbsoluteURL(url))
			chapter.Name=title
			chapter.ChapterId=chapterId
			chapter.ChapterUrl=e.Request.AbsoluteURL(url)
			chapters=append(chapters, chapter)
		})

		book.Name=name
		book.Desc=desc
		book.ImageLink=imageLink
		book.Id="3629"
		book.Chapters=chapters

		//fmt.Println(book)

		jsonBytes, err := json.Marshal(&book)
		if err != nil {
			fmt.Printf("%v json.Marshal failed,err= %v\n", book.Name, err)
			return
		}

		timeStr := strconv.FormatInt(time.Now().Unix(), 10)

		//把id保存到一个序列里面
		//把内容保存到一个hashMap里
		err = redis_utils.Push2RedisSortedSet(conn,book.Id , KEY_COMIC_BOOK_ID_IN_REDIS, KEY_COMIC_BOOK_INFO_IN_REDIS, timeStr, string(jsonBytes))
		if err != nil {
			fmt.Printf("%v push2RedisList failed,err= %v\n", book.Name, err)
			return
		} else {
			fmt.Printf("%s 保存漫画基本信息完毕\n", book.Name)
			startSpiderChapter(book,conn,onSpiderFinish)
		}

	})

	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})


	pageCollector.UserAgent = "Mozilla/5.0 (Linux; U; Android 8.1.0; zh-cn; BLA-AL00 Build/HUAWEIBLA-AL00) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.132 MQQBrowser/8.9 Mobile Safari/537.36"

	pageCollector.Visit(startUrl)
}

func startSpiderChapter(book comic.ComicBook,conn redis.Conn, onSpiderFinish func()) {
	for i := 0; i< len(book.Chapters);i++{
		ChapterSpider(book.Chapters[i],conn,onSpiderFinish)
	}
}


