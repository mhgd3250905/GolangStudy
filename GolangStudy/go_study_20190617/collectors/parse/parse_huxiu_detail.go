package parse

import (
	"GolangStudy/GolangStudy/go_study_20190617/modles/detailText"
	"GolangStudy/GolangStudy/go_study_20190617/modles/detailType"
	"GolangStudy/GolangStudy/go_study_20190617/modles/normal_news"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)


/**
保存基本文字
没有任何样式的文字
*/
func SaveNormalText(text string, content normal_news.Content) normal_news.Content {
	content = saveContent(content, func(contendDetail *normal_news.ContentDetail) {
		contendDetail.ContentType = detailType.TEXT
		//文字类型
		contendDetail.AppendContent(text)
		contendDetail.TextStyle = detailText.Normal
	})
	return content
}
/**
保存图片注释
*/
func SaveImgNote(text string, content normal_news.Content) normal_news.Content {
	content = saveContent(content, func(contendDetail *normal_news.ContentDetail) {
		contendDetail.ContentType = detailType.TEXT
		//文字类型
		contendDetail.AppendContent(text)
		contendDetail.TextStyle = detailText.ImgNote
	})
	return content
}

/**
保存段落标题
*/
func SaveParagraphTitle(text string, content normal_news.Content) normal_news.Content {
	//段落标题
	content = saveContent(content, func(contendDetail *normal_news.ContentDetail) {
		contendDetail.ContentType = detailType.TEXT
		contendDetail.AppendContent(text)
		contendDetail.TextStyle = detailText.ParagraphTitle
	})
	return content
}

/**
保存换行
*/
func SaveBrNode(child *goquery.Selection, content normal_news.Content) normal_news.Content {
	if len(child.Nodes) > 0 {
		if child.Nodes[0] != nil {
			if child.Nodes[0].DataAtom == atom.Br {
				content = saveContent(content, func(contendDetail *normal_news.ContentDetail) {
					contendDetail.ContentType = detailType.TEXT //文字类型
					contendDetail.AppendContent("</br>")
					contendDetail.TextStyle = detailText.Br
				})
			}
		}
	}
	return content
}

/**
保存图片
*/
func SaveImgNode(child *goquery.Selection, content normal_news.Content) normal_news.Content {
	if len(child.Nodes) > 0 {
		if child.Nodes[0] != nil {
			if child.Nodes[0].DataAtom == atom.Img {
				content = saveContent(content, func(contendDetail *normal_news.ContentDetail) {
					contendDetail.ContentType = detailType.IMG //图片类型
					contendDetail.TextStyle = detailText.Img
					img, exist := child.Attr("src")
					if exist {
						contendDetail.Extra = img
					}
				})

			}
		}
	}
	return content
}

/**
保存特殊的具有样式的文本
譬如：span/b/..
*/
func SaveSpecialText(child *goquery.Selection, content normal_news.Content) normal_news.Content {

	firstNode := GetFirstNode(child)

	content = saveContent(content, func(contendDetail *normal_news.ContentDetail) {
		//处理特殊样式信息
		if firstNode != nil {
			if firstNode.DataAtom == atom.Span {
				//span类型 灰色小字体
				contendDetail.ContentType = detailType.TEXT
				contendDetail.AppendContent(child.Text())
				contendDetail.TextStyle = detailText.Span
			} else if firstNode.DataAtom == atom.A {
				//超链接
				contendDetail.ContentType = detailType.TEXT
				contendDetail.AppendContent(child.Text())
				contendDetail.TextStyle = detailText.Link
				link, exist := child.Attr("href")
				if exist {
					contendDetail.Extra = link
				}
			} else if firstNode.DataAtom == atom.Strong {
				//粗体
				contendDetail.ContentType = detailType.TEXT
				contendDetail.AppendContent(child.Text())
				contendDetail.TextStyle = detailText.Bold
			} else {
				//正常文本
				contendDetail.ContentType = detailType.TEXT
				contendDetail.AppendContent(child.Text())
				contendDetail.TextStyle = detailText.Normal
			}
		}
	})
	return content
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

func saveContent(content normal_news.Content, setContentInfo func(contendDetail *normal_news.ContentDetail)) normal_news.Content {
	contendDetail := normal_news.ContentDetail{}

	setContentInfo(&contendDetail)

	content.ContentDetails = append(content.ContentDetails, contendDetail)
	return content
}