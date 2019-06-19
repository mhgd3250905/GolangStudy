package collectors

import (
	"fmt"
	"github.com/gocolly/colly"
)

func GetDetailCollector() (collector *colly.Collector) {

	//解析书本详情的收集器
	detailCollector := colly.NewCollector()

	detailCollector.OnRequest(func(r *colly.Request) {

	})

	detailDownloadSelectorStr := "#mbm-book-links1 > div > ul > li.mbm-book-download-links-listitem > a";
	detailCollector.OnHTML(detailDownloadSelectorStr, func(e *colly.HTMLElement) {
		downloadType := e.Text
		downloadUrl := e.Attr("href")
		fmt.Printf("%v : %v\n", downloadType, downloadUrl)
	})

	detailCollector.OnScraped(func(response *colly.Response) {
		fmt.Println("detailCollector OnScraped")
	})

	return detailCollector

}
