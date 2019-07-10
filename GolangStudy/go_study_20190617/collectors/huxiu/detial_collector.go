package huxiu

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/parse"
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/huxiu"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/gomodule/redigo/redis"
)

const DETAIL_URL = "https://www.huxiu.com/article/307777.html"

func HuxiuDetailSpider(conn redis.Conn, news huxiu.HuxiuNews) {
	//startUrl := news.NewsLink
	startUrl := DETAIL_URL

	//使用代理
	rp, err := proxy.RoundRobinProxySwitcher("http://123.52.19.47:3128",
		"https://218.91.112.188:9999",
		"https://1.192.243.9:9999",
		"https://58.253.157.205:9999",
		"https://180.119.141.211:9999",
		"https://120.83.107.7:9999",
		"https://1.192.242.122:9999")

	if err != nil {
		fmt.Println(err)
	}

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()

	pageCollector.SetProxyFunc(rp)

	//解析页面新闻条目
	newsItemSelectorStr := "#article_cstaontent307777"

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
		detail := huxiu.HuxiuDetail{}
		contents := make([]huxiu.Content, 0)
		e.DOM.Find("p").Each(func(i int, p *goquery.Selection) {
			if p.Children().Length() == 0 {
				//保存单纯的文本内容
				contents = parse.SaveNormalText(p.Text(), contents)
			} else {
				p.Children().Each(func(i int, child *goquery.Selection) {
					prevText := ""
					childText := child.Text()
					nextText := ""

					firstNode := parse.GetFirstNode(child)
					if child.Index() == 0 {
						if firstNode != nil && firstNode.PrevSibling != nil && firstNode.PrevSibling.Type == 1 {
							prevText = child.Nodes[0].PrevSibling.Data
						}
					}

					if firstNode != nil && firstNode.NextSibling != nil && firstNode.NextSibling.Type == 1 {
						nextText = child.Nodes[0].NextSibling.Data
					}

					//如果prev
					if prevText != "" {
						//保存PrevText
						contents = parse.SaveNormalText(prevText, contents)
					}

					//如果childText为空则判断您是否为换行
					if childText == "" {
						//保存换行符
						contents = parse.SaveBrNode(child, contents)
					} else {
						//保存具有样式的文字
						contents = parse.SaveSpicalText(child, contents)
					}

					//如果prev
					if nextText != "" {
						//保存nextText
						contents = parse.SaveNormalText(nextText, contents)
					}
				})
			}
		})
		detail.Contents = contents

		jsonBytes, err := json.Marshal(&detail)
		if err != nil {
			fmt.Printf("%v json.Marshal failed,err= %v\n", news.Title, err)
			return
		}

		err = redis_utils.Push2RedisSortedSet(conn, KEY_HUXIU_DETIAL_IN_REDIS, news.CreateTime, string(jsonBytes))
		if err != nil {
			fmt.Printf("%v push2RedisList failed,err= %v\n", news, err)
			return
		} else {
			fmt.Printf("%s -----------------------Detail爬取完毕-----------------------\n", news.Title)
		}
		//fmt.Println(detail)
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
