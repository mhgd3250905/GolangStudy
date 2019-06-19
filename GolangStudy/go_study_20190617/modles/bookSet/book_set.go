package bookSet

type Book struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	AuthorLink string `json:"author_link"`
	Image string `json:"image"`
	BookLink   string `json:"book_link"`
}
