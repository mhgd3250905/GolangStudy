package main

import (

	"GolangStudy/GolangStudy/go_study_20190617/modles/huxiu"
	"GolangStudy/GolangStudy/goquery_study/parse"
	"GolangStudy/GolangStudy/goquery_study/res_str"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

/*
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
	scopeMarkerNode
*/

func main() {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(res_str.HUXIU_DETAIL_HTML_TEXT))
	if err != nil {
		fmt.Println("goquery failed,err= ", err)
	}
	//fmt.Println(dom.Find("p").Text())

	detail := huxiu.HuxiuDetail{}
	contents := make([]huxiu.Content, 0)

	divContent := dom.Find("#article_cstaontent307896").First()

	divContent.Find("p").Each(func(i int, p *goquery.Selection) {
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

	fmt.Println(detail)
}
