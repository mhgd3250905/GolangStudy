package comic_ikk

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/comic"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"os"
	"regexp"
	"strings"
)

/**
漫画的存储数据结构
- 漫画id name
	- 章节*n
		- 章节中具体的图片*n
 */

/**
对各个章节图片进行下载
 */
const BASEURL  = "http://comic.ikkdm.com/"

func ChapterSpider(bookId string,chapter comic.Chapter,conn redis.Conn, onSpiderFinish func()) {

	startUrl:=chapter.ChapterUrl

	//解析网页漫画图片收集器
	chapterCollector := colly.NewCollector()


	chapterCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	chapterCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	chapterCollector.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
		htmlStr:=string(response.Body)

		hasNextPage:=false

		///comiclist/941/18386/2.htm
		re, _ := regexp.Compile(`<script src='/ad/sc_soso\.js'.+/images/d\.gif`)
		all := re.FindAll([]byte(htmlStr), 1)
		nextHtmlStr := ""
		for i, _ := range all {
			nextHtmlStr = string(all[i])
		}

		///comiclist/941/18386/2.htm
		re, _ = regexp.Compile(`comiclist/.+\.htm`)
		all = re.FindAll([]byte(nextHtmlStr), 1)
		nextUrl := ""
		for i, _ := range all {
			nextUrl = string(all[i])
		}

		///manhua/3629/432953_2.html
		if  strings.HasPrefix(nextUrl,"comiclist/"){
			hasNextPage=true
		}

		re, _ = regexp.Compile(`kuku7comic7/.+<span`)
		all = re.FindAll([]byte(htmlStr), 1)
		imageHtmlStr := ""
		for i, _ := range all {
			imageHtmlStr = string(all[i])
		}

		re, _ = regexp.Compile(`kuku7comic7/\d+/\d+/.+\.jpg`)
		all = re.FindAll([]byte(imageHtmlStr), 1)
		imageUrl := ""
		for i, _ := range all {
			imageUrl = string(all[i])
		}

		if hasNextPage {
			chapter.ChapterUrl=fmt.Sprintf("%s%s",BASEURL,nextUrl)
			ChapterSpider(bookId,chapter,conn,onSpiderFinish)
		}else {
			fmt.Println("当前已是章节最后一页！")
		}

		//把id保存到一个序列里面
		//把内容保存到一个hashMap里
		err := redis_utils.SaveList(conn,chapter.ChapterId,imageUrl)
		if err != nil {
			fmt.Printf("%v SaveList failed,err= %v\n", chapter.Name, err)
			return
		} else {
			fmt.Printf("%s 保存漫画基本信息完毕\n", chapter.Name)
		}
	})


	chapterCollector.OnScraped(func(response *colly.Response) {
		//fmt.Println("chapterCollector OnScraped")
	})

	chapterCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("chapterCollector.OnError: ", e)
	})


	chapterCollector.UserAgent = "Mozilla/5.0 (Linux; U; Android 8.1.0; zh-cn; BLA-AL00 Build/HUAWEIBLA-AL00) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.132 MQQBrowser/8.9 Mobile Safari/537.36"


	chapterCollector.Visit(startUrl)
}


//判断文件文件夹是否存在
func isFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

