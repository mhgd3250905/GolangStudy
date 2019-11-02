package comic_ikk

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
	"strconv"
	"strings"
	"time"
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
const BASEURL = "http://comic.ikkdm.com/"
const BASE_IMAGE_URL = "http://n9.1whour.com/"

const IMAGE_DIR_PATH="E:/comic_spider/"

func ChapterSpider(bookId string, chapter comic.Chapter, conn redis.Conn, onSpiderFinish func()) {

	startUrl := chapter.ChapterUrl

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
		htmlStr := ConvertToString(string(response.Body), "gbk", "utf-8")

		hasNextPage := false

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
		if strings.HasPrefix(nextUrl, "comiclist/") {
			hasNextPage = true
		}

		re, _ = regexp.Compile(`kuku7comic7/.+<span`)
		all = re.FindAll([]byte(htmlStr), 1)
		imageHtmlStr := ""
		for i, _ := range all {
			imageHtmlStr = string(all[i])
		}

		imageUrl := ""
		if imageHtmlStr == "" {
			re, _ = regexp.Compile(`newkuku/.+<span`)
			all = re.FindAll([]byte(htmlStr), 1)
			imageHtmlStr := ""
			for i, _ := range all {
				imageHtmlStr = string(all[i])
			}

			re, _ = regexp.Compile(`newkuku/\d+/\d+/.+\.jpg`)
			all = re.FindAll([]byte(imageHtmlStr), 1)
			for i, _ := range all {
				imageUrl = string(all[i])
			}
		} else {
			re, _ = regexp.Compile(`[a-z]*kuku*/\d+/\d+/.+\.jpg`)
			all = re.FindAll([]byte(imageHtmlStr), 1)
			for i, _ := range all {
				imageUrl = string(all[i])
			}
		}


		//http://comic.ikkdm.com/comiclist/1748/32430/1.htm
		re, _ = regexp.Compile(`[0-9]+`)
		all = re.FindAll([]byte(startUrl), 3)
		imageName := ""
		for i, _ := range all {
			imageName = string(all[i])
		}

		//构建保存路径 ../名称/章节/id.png
		dirPath:=fmt.Sprintf("%s%s",IMAGE_DIR_PATH,bookId)

		exist,err:=isFileExist(dirPath)
		if err != nil {
			fmt.Println("isFileExist run failed,err = ",err)
			return
		}

		if !exist {
			err=os.Mkdir(dirPath, os.ModePerm)
		}

		dirPath=fmt.Sprintf("%s%s/%s",IMAGE_DIR_PATH,bookId,chapter.ChapterId)

		exist,err=isFileExist(dirPath)
		if err != nil {
			fmt.Println("isFileExist run failed,err = ",err)
			return
		}

		if !exist {
			err=os.Mkdir(dirPath, os.ModePerm)
		}

		res, err := http.Get(fmt.Sprintf("%s%s", BASE_IMAGE_URL, imageUrl))
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
			//if hasNextPage {
			//	chapter.ChapterUrl=e.Request.AbsoluteURL(nextUrl)
			//	ChapterSpider(bookId,chapter,conn,onSpiderFinish)
			//}else {
			//	fmt.Println("当前已是章节最后一页！")
			//}
		}

		//把id保存到一个序列里面
		//把内容保存到一个hashMap里

		//strbytes := []byte(fmt.Sprintf("%s%s", BASE_IMAGE_URL, imageUrl))
		//encoded := base64.StdEncoding.EncodeToString(strbytes)
		//
		timeStr := strconv.FormatInt(time.Now().Unix(), 10)
		err = redis_utils.SaveZset(conn, chapter.ChapterId, timeStr, imagePath)
		if err != nil {
			fmt.Printf("%v SaveZset failed,err= %v\n", chapter.Name, err)
			return
		} else {
			fmt.Printf("%s 保存章节基本信息完毕 url: %s\n", chapter.Name, chapter.ChapterUrl)
		}

		if hasNextPage {
			chapter.ChapterUrl = fmt.Sprintf("%s%s", BASEURL, nextUrl)
			ChapterSpider(bookId, chapter, conn, onSpiderFinish)
		} else {
			fmt.Println("当前已是章节最后一页！")
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
