package comic_ikk

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
	"strings"
	"time"
)

/**
漫画的存储数据结构
- 漫画id name
	- 章节*n
		- 章节中具体的图片*n
*/

//ikk 进击的巨人
//const MAIN_URL = "http://comic.ikkdm.com/comiclist/941/"
//鬼灭之刃
const MAIN_URL = "http://comic.ikkdm.com/comiclist/2126/"

const KEY_COMIC_BOOK_ID_IN_REDIS = "COMIC_BOOK_ID"
const KEY_COMIC_BOOK_INFO_IN_REDIS = "COMIC_BOOK_INFO"
const KEY_COMIC_CHAPTER_LIST_IN_REDIS = "COMIC_CHAPTER_LIST"
const KEY_COMIC_CHAPTER_DETAIL_IN_REDIS = "COMIC_CHAPTER_DETAIL"

func ComicSpider(conn redis.Conn, onSpiderFinish func()) {

	book := comic.ComicBook{}

	startUrl := MAIN_URL

	//解析网页漫画信息收集器
	pageCollector := colly.NewCollector()

	//解析漫画信息容器
	bodySelectorStr := "body"
	itemSelectorStr := "#comiclistn > dd"

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		//r.Headers.Add("content-type","text/html")
		//r.Headers.Add("Cache-Control","no-cache")
		//r.Headers.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
		//r.Headers.Add("Accept-Encoding","gzip, deflate, br")
		//r.Headers.Add("Connection","keep-alive")
		//r.Headers.Add("Pragma","no-cache")
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	pageCollector.OnResponse(func(response *colly.Response) {
		responseStr := ConvertToString(string(response.Body), "gbk", "utf-8")
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(responseStr))
		if err != nil {
			fmt.Println("str to dom failed,err= ", err)
			return
		}
		body := dom.Find(bodySelectorStr).First()

		img := body.Find("table:nth-child(5) > tbody > tr > td:nth-child(2) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(2) > td:nth-child(1) > img")

		imageLink, exist := img.Attr("src")
		if !exist {
			fmt.Println("图片链接不存在！")
		}

		name := body.Find("table:nth-child(5) > tbody > tr > td:nth-child(2) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td").First().Text()

		desc := body.Find("#ComicInfo").Text()

		chapters := make([]comic.Chapter, 0)
		body.Find(itemSelectorStr).Each(func(i int, dd *goquery.Selection) {
			chapter := comic.Chapter{}

			title := dd.Find("a").First().Text()

			url, exist := dd.Find("a").First().Attr("href")
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
			chapter.Name = title
			chapter.ChapterId = chapterId
			chapter.ChapterUrl = response.Request.AbsoluteURL(url)
			chapters = append(chapters, chapter)
		})

		// /manhua/3629/608970.html
		re, _ := regexp.Compile(`[0-9]+`)
		all := re.FindAll([]byte(MAIN_URL), 1)
		bookId := ""
		for i, _ := range all {
			bookId = string(all[i])
		}

		book.Name = ConvertToString(name, "gbk", "utf-8")
		book.Desc = ConvertToString(desc, "gbk", "utf-8")
		book.ImageLink = imageLink
		book.Id = bookId
		book.Chapters = chapters

		//fmt.Println(book)

		jsonBytes, err := json.Marshal(&book)
		if err != nil {
			fmt.Printf("%v json.Marshal failed,err= %v\n", book.Name, err)
			return
		}

		timeStr := strconv.FormatInt(time.Now().Unix(), 10)

		//把id保存到一个序列里面
		//把内容保存到一个hashMap里
		err = redis_utils.Push2RedisSortedSet(conn, book.Id, KEY_COMIC_BOOK_ID_IN_REDIS, KEY_COMIC_BOOK_INFO_IN_REDIS, timeStr, string(jsonBytes))
		if err != nil {
			fmt.Printf("%v push2RedisList failed,err= %v\n", book.Name, err)
			return
		} else {
			fmt.Printf("%s 保存漫画基本信息完毕\n", book.Name)
			startSpiderChapter(book, conn, onSpiderFinish)
		}
	})

	pageCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("pageCollector OnScraped")
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})

	pageCollector.UserAgent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Mobile Safari/537.36"

	pageCollector.Visit(startUrl)
}

func startSpiderChapter(book comic.ComicBook, conn redis.Conn, onSpiderFinish func()) {
	for i := 0; i < len(book.Chapters); i++ {
		//if book.Chapters[i].ChapterId != "475934" {
		//	continue
		//}
		ChapterSpider(book.Id, book.Chapters[i], conn, onSpiderFinish)
	}
}

