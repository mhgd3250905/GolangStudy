package ifanr

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/parse"
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/detailContainerType"
	"GolangStudy/GolangStudy/go_study_20190617/modles/detailText"
	"GolangStudy/GolangStudy/go_study_20190617/modles/detailType"
	"GolangStudy/GolangStudy/go_study_20190617/modles/normal_news"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/net/html/atom"
)

const KEY_IFANR_DETAIL_IN_REDIS = "ifanr_detail"

func IfanrDetailSpider(conn redis.Conn, news normal_news.News) {
	startUrl := news.NewsLink

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()

	//pageCollector.SetProxyFunc(rp)

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

	//解析页面新闻条目
	newsItemSelectorStr := "div#entry-content"

	onHtmlFunc := func(e *colly.HTMLElement) {
		detail := normal_news.NewsDetail{}
		contents := make([]normal_news.Content, 0)

		//解析主要文本内容
		contents = parseIfanrContentChildRen(e.DOM, contents)

		detail.HuxiuNews = news
		detail.Contents = contents

		jsonBytes, err := json.Marshal(&detail)
		if err != nil {
			fmt.Printf("%v json.Marshal failed,err= %v\n", news.Title, err)
			return
		}

		err = redis_utils.SaveHashMap(conn, KEY_IFANR_DETAIL_IN_REDIS, detail.HuxiuNews.NewsId, string(jsonBytes))
		if err != nil {
			fmt.Printf("%v push2RedisList failed,err= %v\n", news, err)
			return
		} else {
			fmt.Printf("%s -----------------------Detail爬取完毕-----------------------\n", news.Title)
		}
		//fmt.Println(detail)
	}

	pageCollector.OnHTML(newsItemSelectorStr, onHtmlFunc)

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

func parseIfanrContentChildRen(divContent *goquery.Selection, contents []normal_news.Content) []normal_news.Content {
	divContent.Children().Each(func(i int, child *goquery.Selection) {
		content := normal_news.Content{}
		childFirstNode := parse.GetFirstNode(child)
		if childFirstNode.DataAtom == atom.H3 {
			//解析小标题
			content = ParseIfanrTitle(child)
			contents = append(contents, content)
		} else if childFirstNode.DataAtom == atom.Div && child.HasClass("editor-image-source") {
			//图片注释
			content = ParseIfanrImgNote(child)
			contents = append(contents, content)

		} else if childFirstNode.DataAtom == atom.P {
			//解析文本
			contents=parseIfanPChild(child, content,contents)
		}
	})

	return contents
}

func ParseIfanrImgNote(div *goquery.Selection) normal_news.Content {
	//保存标题内容
	content := normal_news.Content{}
	content.ContentDetails = make([]normal_news.ContentDetail, 0)
	contentDetail := normal_news.ContentDetail{}
	content.ContentContainerType = detailContainerType.ImgNote
	contentDetail.ContentType = detailType.TEXT
	contentDetail.ContentDetail = div.Find("p").First().Text()
	contentDetail.TextStyle = detailText.ImgNote
	content.ContentDetails = append(content.ContentDetails, contentDetail)
	return content
}

//解析chule小标题
func ParseIfanrTitle(h3 *goquery.Selection) normal_news.Content {
	//保存标题内容
	content := normal_news.Content{}
	content.ContentDetails = make([]normal_news.ContentDetail, 0)
	contentDetail := normal_news.ContentDetail{}
	content.ContentContainerType = detailContainerType.Normal
	contentDetail.ContentType = detailType.TEXT
	contentDetail.ContentDetail = h3.Text()
	contentDetail.TextStyle = detailText.ParagraphTitle
	content.ContentDetails = append(content.ContentDetails, contentDetail)
	return content
}

func parseIfanPChild(p *goquery.Selection, content normal_news.Content, contents []normal_news.Content) []normal_news.Content {
	if p.Children().Length() == 0 {
		//保存单纯的文本内容
		//具有特殊class的文字也需要特殊保存
		content = parseIfanrAtomP(p, content)
		contents = append(contents, content)
	} else {
		//如果有子节点，则遍历分析
		//已知：spam/img/strong/a/img
		p.Children().Each(func(i int, child *goquery.Selection) {
			content = parseIfanrAtomP(child, content)
			contents = append(contents, content)
		})
	}
	return contents
}

func parseIfanrAtomP(child *goquery.Selection, content normal_news.Content) normal_news.Content {
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
		content = parse.SaveNormalText(child.Text(), content)
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
	return content
}
