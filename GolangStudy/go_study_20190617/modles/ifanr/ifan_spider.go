package ifanr

type ResponseIfanr struct {
	Objects []Object `json:"objects"`
}

type Object struct {
	CategoryUri string
	CreatedAt int `json:"created_at"`
	CreatedAtFormat string `json:"created_at_format"`
	CreatedBy Author `json:"created_by"`
	Id int `json:"id"`
	PostCategory string `json:"post_category"`
	PostCoverImage string `json:"post_cover_image"`
	PostExcerpt string `json:"post_excerpt"`
	PostId string `json:"post_id"`
	PostTitle string `json:"post_title"`
	PostType string `json:"post_type"`
	PostUrl string `json:"post_url"`
}

type Author struct {
	AuthorUrl string `json:"author_url"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	Id int `json:"id"`
	Job string `json:"job"`
	Name string `json:"name"`
	Signature string `json:"signature"`
}
