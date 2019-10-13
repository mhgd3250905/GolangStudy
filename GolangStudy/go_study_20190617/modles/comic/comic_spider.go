package comic



type ComicBook struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ImageLink string `json:"image_link"`
	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	ChapterId string `json:"chapter_id"`
	Name string `json:"name"`
	ChapterUrl string `json:"chapter_url"`
}

type ChapterDetails struct {
	ChapterId string `json:"chapter_id"`
	imageUrls []string `json:"image_urls"`
}





