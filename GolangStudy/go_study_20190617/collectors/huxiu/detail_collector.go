package huxiu

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/parse"
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/detailContainerType"
	"GolangStudy/GolangStudy/go_study_20190617/modles/normal_news"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/net/html/atom"
)

const DETAIL_URL = "https://wwww.huxiu.com/article/308860.html"

func HuxiuDetailSpider(conn redis.Conn, news normal_news.News) {
	startUrl := news.NewsLink
	//startUrl := DETAIL_URL

	////使用代理
	//rp, err := proxy.RoundRobinProxySwitcher("http://222.76.75.7:5781",)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()

	//pageCollector.SetProxyFunc(rp)

	//解析页面新闻条目
	newsItemSelectorStr := "#article_cstaontent308860"
	newsItemSelectorStr_2 := "#article-detail-content"

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	pageCollector.OnResponse(func(response *colly.Response) {
		//str:=string(response.Body)
		//fmt.Println(str)
	})

	onHtmlFunc := func(e *colly.HTMLElement) {
		detail := normal_news.NewsDetail{}
		contents := make([]normal_news.Content, 0)

		//解析主要文本内容
		contents = ParseContentChildRen(e.DOM, contents)

		detail.HuxiuNews=news
		detail.Contents = contents

		jsonBytes, err := json.Marshal(&detail)
		if err != nil {
			fmt.Printf("%v json.Marshal failed,err= %v\n", news.Title, err)
			return
		}

		err = redis_utils.SaveHashMap(conn, KEY_HUXIU_DETIAL_IN_REDIS, detail.HuxiuNews.NewsId, string(jsonBytes))
		if err != nil {
			fmt.Printf("%v push2RedisList failed,err= %v\n", news, err)
			return
		} else {
			fmt.Printf("%s -----------------------Detail爬取完毕-----------------------\n", news.Title)
		}
		//fmt.Println(detail)
	}

	pageCollector.OnHTML(newsItemSelectorStr, onHtmlFunc)
	pageCollector.OnHTML(newsItemSelectorStr_2, onHtmlFunc)

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

func ParseContentChildRen(divContent *goquery.Selection, contents []normal_news.Content) []normal_news.Content {
	divContent.Find("p").Each(func(i int, child *goquery.Selection) {
		content := normal_news.Content{}
		childFirstNode := parse.GetFirstNode(child)
		parentFirstNode := parse.GetFirstNode(child.Parent())

		if parentFirstNode.DataAtom == atom.Blockquote {
			//对于遍历的P,其中部分被<blockquote>包裹，所以需要判断并特殊保存
			content.ContentContainerType = detailContainerType.Blockquote
			content = ParseAtomP(child, content)
		} else if childFirstNode.DataAtom == atom.P {
			if child.HasClass("text-img-note") {
				content.ContentContainerType = detailContainerType.ImgNote
			}
			content = ParseAtomP(child, content)
		}
		contents = append(contents, content)
	})
	return contents
}

func ParseAtomP(p *goquery.Selection, content normal_news.Content) normal_news.Content {
	if p.Children().Length() == 0 {
		//保存单纯的文本内容
		//具有特殊class的文字也需要特殊保存
		if p.HasClass("text-big-title") {
			//大字标题
			content = parse.SaveParagraphTitle(p.Text(), content)
		} else {
			//保存标准文本
			content = parse.SaveNormalText(p.Text(), content)
		}
	} else {
		//如果有子节点，则遍历分析
		//已知：spam/img/strong/a/img
		p.Children().Each(func(i int, child *goquery.Selection) {

			prevText := ""
			nextText := ""

			//获取包含的节点
			firstNode := parse.GetFirstNode(child)

			//解析前置文本
			if child.Index() == 0 {
				if firstNode != nil && firstNode.PrevSibling != nil && firstNode.PrevSibling.Type == 1 {
					prevText = child.Nodes[0].PrevSibling.Data
				}
			}

			//解析后置文本
			if firstNode != nil && firstNode.NextSibling != nil && firstNode.NextSibling.Type == 1 {
				nextText = child.Nodes[0].NextSibling.Data
			}

			//如果prev存在就保存
			if prevText != "" {
				//保存PrevText
				content = parse.SaveNormalText(prevText, content)
			}

			//保存具有样式的文字
			if firstNode.DataAtom == atom.Br {
				content = parse.SaveBrNode(child, content)
			} else if firstNode.DataAtom == atom.Img {
				content = parse.SaveImgNode(child, content)
				content.ContentContainerType = detailContainerType.Img
			} else {
				content = parse.SaveSpecialText(child, content)
			}

			//如果prev
			if nextText != "" {
				//保存nextText
				if firstNode.DataAtom == atom.Img {
					content = parse.SaveImgNote(nextText, content)
				} else {
					content = parse.SaveNormalText(nextText, content)
				}
			}
		})
	}
	return content
}
