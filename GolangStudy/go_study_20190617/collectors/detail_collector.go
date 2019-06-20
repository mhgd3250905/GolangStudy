package collectors

import (
	"GolangStudy/GolangStudy/go_study_20190617/modles/bookSet"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
)

const KEY_BOOK_DETAIL_IN_REDIS  = "book_detail"


func GetDetailCollector(conn redis.Conn) (collector *colly.Collector) {

	//解析书本详情的收集器
	detailCollector := colly.NewCollector()

	detailCollector.DetectCharset = true

	detailCollector.OnRequest(func(r *colly.Request) {

	})

	detailDownloadSelectorStr := "#mbm-book-page"
	detailCollector.OnHTML(detailDownloadSelectorStr, func(e *colly.HTMLElement) {
		img := e.DOM.Find("span > img")
		time_span := e.DOM.Find("div.mbm-book-details-outer > div > span.mbm-book-details-published-data > span > span.mbm-book-published-text").First()
		title_a := e.DOM.Find("div.mbm-book-details-outer > div > span:nth-child(4) > a").First()
		author_a := e.DOM.Find("div.mbm-book-details-outer > div > span.mbm-book-details-editors-data > a").First()
		doubanScore_span := e.DOM.Find("#mbm-book-page > div.mbm-book-details-outer > div > span:nth-child(13)").First()
		doubanLink_a := doubanScore_span.Find("a").First()
		introduction_p_arr := e.DOM.Find("div.mbm-book-excerpt > span.mbm-book-excerpt-text p")
		download_a_arr := e.DOM.Find("#mbm-book-links1 > div > ul a")

		title := title_a.Text()
		author := author_a.Text()
		time := time_span.Text()
		image, _ := img.Attr("src")
		doubanScoreStr := doubanScore_span.Text()

		var doubanScore float64
		var doubanScoreCount int64
		if splits := strings.Split(doubanScoreStr, "/"); len(splits) >= 2 {
			doubanScore,_ = strconv.ParseFloat(splits[0][:len(splits[0])-4],64)
			doubanScoreCount,_ = strconv.ParseInt(splits[1][:len(splits[1])-14],10,64)
		}
		doubanLink, _ := doubanLink_a.Attr("href")

		var buffer bytes.Buffer
		introduction_p_arr.Each(func(i int, selection *goquery.Selection) {
			buffer.WriteString(selection.Text())
		})
		introduction := buffer.String()

		downloadLinkEpub := ""
		downloadLinkAzw3 := ""
		downloadLinkMobi := ""

		download_a_arr.Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 0:
				downloadLinkEpub, _ = selection.Attr("href")
			case 1:
				downloadLinkAzw3, _ = selection.Attr("href")
			case 2:
				downloadLinkMobi, _ = selection.Attr("href")
			}
		})

		//fmt.Printf("title: %v,author: %v,time: %v,image: %v,score: %v,count: %v,doubanLink: %v,intro: %v,epub: %v,azw3: %v,mobi: %v",
		//	title, author, time, image, doubanScore, doubanScoreCount, doubanLink, introduction, downloadLinkEpub,
		//	downloadLinkAzw3, downloadLinkMobi)

		bookDetail:=bookSet.BookDetail{
			Title:title,
			Author:author,
			Time:time,
			Image:image,
			DoubanScore:doubanScore,
			DoubleScoreCount:doubanScoreCount,
			DoubanLink:doubanLink,
			Introduction:introduction,
			DownloadLinkEpub:downloadLinkEpub,
			DownloadLinkAzw3:downloadLinkAzw3,
			DownloadLinkMobi:downloadLinkMobi,
		}

		jsonBytes, err := json.Marshal(&bookDetail)
		if err != nil {
			fmt.Printf("%v json.Marshal failed,err= %v\n", title, err)
			return
		}

		err = push2RedisList(conn, KEY_BOOK_DETAIL_IN_REDIS, string(jsonBytes))
		if err != nil {
			fmt.Printf("%v push2RedisList failed,err= %v\n", title, err)
			return
		}

		fmt.Printf("%v save detail success!\n", title)

	})

	detailCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("detailCollector OnScraped")
	})

	return detailCollector
}

func push2RedisList(c redis.Conn, key string, content string) (err error) {
	_, err = c.Do("RPUSH", key, content)
	return
}
