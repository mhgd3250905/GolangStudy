package huxiu

type HuxiuNews struct {
	Title      string   `json:"title"`
	NewsLink   string   `json:"news_link"`
	Author     Author   `json:"author"`
	CreateTime string   `json:"create_time"`
	Desc       string   `json:"desc"`
	ImgLink    string   `json:"image_link"`
	Categorys   []Category `json:"category"`
}

type Author struct {
	AuthorName string `json:"author_name"`
	AuthorImg string `json:"author_img"`
	AuthorId string `json:"author_id"`
}

type Category struct {
	CategoryName string `json:"category_name"`
	CategoryId   string    `json:"category_id"`
}
