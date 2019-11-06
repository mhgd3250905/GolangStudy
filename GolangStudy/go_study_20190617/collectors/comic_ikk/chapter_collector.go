package comic_ikk

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/redis_utils"
	"GolangStudy/GolangStudy/go_study_20190617/modles/comic"
	"encoding/base64"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"os"
	"regexp"
	"strconv"
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
const BASEURL = "http://m.kukudm.com/"
const BASE_IMAGE_URL = "https://s1.kukudm.com/"

const IMAGE_DIR_PATH = "E:/comic_spider/"

func ChapterSpider(bookId string, chapter comic.Chapter, conn redis.Conn, onSpiderFinish func()) {

	//startUrl := "http://m.kukudm.com/comiclist/1512/43058/1.htm"
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

		//document.write("<a href='/comiclist/1748/32430/2.htm'><IMG
		re, _ := regexp.Compile(`document.write.+<IMG`)
		all := re.FindAll([]byte(htmlStr), 1)
		nextHtmlStr := ""
		for i, _ := range all {
			nextHtmlStr = string(all[i])
		}

		///comiclist/941/18386/2.htm
		re, _ = regexp.Compile(`/comiclist/.+\.htm`)
		all = re.FindAll([]byte(nextHtmlStr), 1)
		nextUrl := ""
		for i, _ := range all {
			nextUrl = string(all[i])
		}

		///manhua/3629/432953_2.html
		if strings.HasPrefix(nextUrl, "/comiclist/") {
			hasNextPage = true
		}

		//><IMG SRC='"+m2007+"newkuku/2013/201303/20130314/new/亚人01/0003F.jpg'></a><span style='display:none'><img src='"+m201304d+"newkuku/2013/201303/20130314/new/亚人01/0106I.jpg'></span>
		re, _ = regexp.Compile(`><IMG SRC=.+></a>`)
		all = re.FindAll([]byte(htmlStr), 1)
		imageHtmlStr := ""
		for i, _ := range all {
			imageHtmlStr = string(all[i])
		}

		imageUrl := ""

		re, _ = regexp.Compile(`newkuku/\d+/\d+/.+\.jpg`)
		all = re.FindAll([]byte(imageHtmlStr), 1)
		for i, _ := range all {
			imageUrl = string(all[i])
		}

		if imageUrl == "" {
			//kuku7comic7/201009/20100922/进击的巨人/01/cccc_00203K.jpg
			re, _ = regexp.Compile(`kuku7comic7/\d+/\d+/.+\.jpg`)
			all = re.FindAll([]byte(imageHtmlStr), 1)
			for i, _ := range all {
				imageUrl = string(all[i])
			}
		}

		if imageUrl == "" {
			//kuku7comic7/201009/20100922/进击的巨人/01/cccc_00203K.jpg
			re, _ = regexp.Compile(`kuku8comic8/\d+/\d+/.+\.jpg`)
			all = re.FindAll([]byte(imageHtmlStr), 1)
			for i, _ := range all {
				imageUrl = string(all[i])
			}
		}

		strbytes := []byte(fmt.Sprintf("%s%s", BASE_IMAGE_URL, imageUrl))
		encoded := base64.StdEncoding.EncodeToString(strbytes)

		//re, _ = regexp.Compile(`[0-9]+`)
		//all = re.FindAll([]byte(imageUrl), 2)
		//index := ""
		//for i, _ := range all {
		//	index = string(all[i])
		//}
		//timeStr := strconv.FormatInt(time.Now().Unix(), 10)

		fmt.Println(imageUrl)
		indexStr := imageUrl[len(imageUrl)-10 : len(imageUrl)-4]

		index := 0
		sqrt := 10
		for i := len(indexStr) - 1; i >= 0; i-- {
			a := int(indexStr[i])
			sqrt *= 10
			index += a * sqrt
		}

		indexStr = strconv.Itoa(index)

		err := redis_utils.SaveZset(conn, chapter.ChapterId, indexStr, encoded)
		if err != nil {
			fmt.Printf("%v SaveZset failed,err= %v\n", chapter.Name, err)
			return
		} else {
			fmt.Printf("%s 保存章节基本信息完毕\n", chapter.Name)
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
