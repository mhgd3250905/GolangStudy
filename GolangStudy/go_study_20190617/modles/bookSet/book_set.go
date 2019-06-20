package bookSet

type Book struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	AuthorLink string `json:"author_link"`
	Image string `json:"image"`
	BookLink   string `json:"book_link"`
}

type BookDetail struct {
	Title            string  `json:"title"`
	Author           string  `json:"author"`
	Time             string  `json:"time"`
	Image            string  `json:"image"`
	DoubanScore      float64 `json:"douban_score"`
	DoubleScoreCount int64     `json:"double_score_c_ount"`
	DoubanLink       string  `json:"douban_link"`
	Introduction     string  `json:"introduction"`
	DownloadLinkEpub string  `json:"download_link_epub"`
	DownloadLinkAzw3 string  `json:"download_link_azw3"`
	DownloadLinkMobi string  `json:"download_link_mobi"`
}
