package comic

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/comic"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"io"
	"net/http"
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

func ChapterSpider(chapter comic.Chapter,conn redis.Conn, onSpiderFinish func()) {

	startUrl:=chapter.ChapterUrl

	//解析网页漫画图片收集器
	chapterCollector := colly.NewCollector()

	bodySelectorStr := "body > div.t_1200 > div.mh_cont > div.mh_list > p > a"

	chapterCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	chapterCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	chapterCollector.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})

	chapterCollector.OnHTML(bodySelectorStr, func(e *colly.HTMLElement) {

		hasNextPage:=false

		nextUrl,exist:=e.DOM.Attr("href")
		if !exist {
			fmt.Println("下一个图片链接不存在")
		}
		///manhua/3629/432953_2.html
		if  strings.HasPrefix(nextUrl,"/manhua/"){
			hasNextPage=true
		}

		imageLink,exist:=e.DOM.Find("img").First().Attr("src")
		if !exist {
			fmt.Println("图片链接不存在！")
		}

		re, _ := regexp.Compile(`[0-9]+[_]*[0-9]*`)
		all := re.FindAll([]byte(startUrl), 2)
		imageName := ""
		for i, _ := range all {
			imageName = string(all[i])
		}

		//构建保存路径 ../名称/章节/id.png
		dirPath:=fmt.Sprintf("E:/comic_spider/gmzr")

		exist,err:=isFileExist(dirPath)
		if err != nil {
			fmt.Println("isFileExist run failed,err = ",err)
			return
		}

		if !exist {
			err=os.Mkdir(dirPath, os.ModePerm)
		}

		dirPath=fmt.Sprintf("E:/comic_spider/gmzr/%s",chapter.ChapterId)

		exist,err=isFileExist(dirPath)
		if err != nil {
			fmt.Println("isFileExist run failed,err = ",err)
			return
		}

		if !exist {
			err=os.Mkdir(dirPath, os.ModePerm)
		}

		res, err := http.Get(imageLink)
		if err != nil {
			panic(err)
		}

		imagePath:=fmt.Sprintf("%s/%s.png",dirPath,imageName)

		f, err := os.Create(imagePath)
		if err != nil {
			panic(err)
		}
		_,err=io.Copy(f, res.Body)
		if err == nil {
			fmt.Printf("保存图片 %s 成功\n",imageName)
			if hasNextPage {
				chapter.ChapterUrl=e.Request.AbsoluteURL(nextUrl)
				ChapterSpider(chapter,conn,onSpiderFinish)
			}else {
				fmt.Println("当前已是章节最后一页！")
			}
		}

		//把id保存到一个序列里面
		//把内容保存到一个hashMap里
		err = redis_utils.SaveList(conn,chapter.ChapterId,imagePath)
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

