package parse

import (
	"GolangStudy/GolangStudy/go_study_20190617/modles/normal_news"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

/**
保存基本文字
没有任何样式的文字
*/
func SaveNormalText(text string, contents []normal_news.Content) []normal_news.Content {
	content := normal_news.Content{}
	content.ContentType = "text"
	//文字类型
	content.AppendContent(text)
	content.TextStyle = "N"
	contents = append(contents, content)
	return contents
}

/**
保存换行
*/
func SaveBrNode(child *goquery.Selection, contents []normal_news.Content) []normal_news.Content {
	if len(child.Nodes) > 0 {
		if child.Nodes[0] != nil {
			if child.Nodes[0].Data == "br" {
				content := normal_news.Content{}
				content.ContentType = "text" //文字类型
				content.AppendContent("\n")
				content.TextStyle = "Br"
				contents = append(contents, content)
			}
		}
	}
	return contents
}

/**
保存特殊的具有样式的文本
譬如：span/b/..
*/
func SaveSpicalText(child *goquery.Selection, contents []normal_news.Content) []normal_news.Content {
	firstNode := GetFirstNode(child)
	content := normal_news.Content{}
	content.ContentType = "text"
	//文字类型
	content.AppendContent(child.Text())

	//处理特殊样式信息
	if firstNode != nil {
		if firstNode.Data == "span" {
			content.TextStyle = "span"
		} else if firstNode.Data == "a" {
			content.TextStyle = "link"
			link,exist:=child.Attr("href")
			if exist {
				content.Extra=link
			}
		}else if firstNode.Data == "strong" {
			content.TextStyle = "bold"
		} else {
			content.TextStyle = "N"
		}
	}

	contents = append(contents, content)
	return contents
}

/**
获取html.Child的第一个Node
*/
func GetFirstNode(child *goquery.Selection) *html.Node {
	if len(child.Nodes) <= 0 {
		return nil
	}
	return child.Nodes[0]
}
