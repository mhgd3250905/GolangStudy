package huxiu

import "bytes"

type HuxiuNews struct {
	Title      string     `json:"title"`
	NewsLink   string     `json:"news_link"`
	Author     Author     `json:"author"`
	CreateTime string     `json:"create_time"`
	Desc       string     `json:"desc"`
	ImgLink    string     `json:"image_link"`
	Categorys  []Category `json:"category"`
}

type Author struct {
	AuthorName string `json:"author_name"`
	AuthorImg  string `json:"author_img"`
	AuthorId   string `json:"author_id"`
}

type Category struct {
	CategoryName string `json:"category_name"`
	CategoryId   string `json:"category_id"`
}

type Content struct {
	ContentType   string `json:"content_type"`   //内容类型 0=>文字 1：图片
	ContentDetail string `json:"content_detail"` //具体内容,文字的话就是内容，图片的话就是链接
	TextStyle     string `json:"text_style"`     //文字的类型，譬如标签，大标题，小标题，粗体等等
	Extra         string `json:"extra"`
}

func (this *Content) AppendContent(text string) {
	var buffer bytes.Buffer
	buffer.WriteString(this.ContentDetail)
	buffer.WriteString(text)
	this.ContentDetail = buffer.String()
}

type HuxiuDetail struct {
	HuxiuNews HuxiuNews `json:"huxiu_news"`
	Contents  []Content `json:"contents""`
}
