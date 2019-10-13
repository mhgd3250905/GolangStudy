package pentaq

type PentaqItem struct {
	Id int `json:"id"`
	Date string `json:"date"`
	DateGmt string `json:"date_gmt"`
	Modified string `json:"modified"`
	Link string `json:"link"`
	Title Title `json:"title"`
	Content Content `json:"content"`
	Embedded interface{} `json:"_embedded"`
}

type Title struct {
	Rendered string`json:"rendered"`
}

type Content struct {
	Rendered string`json:"rendered"`
}

type Excerpt struct {
	Rendered string`json:"rendered"`
}

type Embedded struct {
	Authors []Author `json:"author"`
	wpFeaturedMedias []interface{} `json:"wp:featuredmedia"`
}

type Author struct {
	Name string`json:"name"`
}



type WpFeaturedmedia struct{
	MediaDetails MediaDetails `json:"media_details"`
	SourceUrl string `json:"source_url"`
}

type MediaDetails struct {
	Sizes Sizes `json:"sizes"`
}


type Sizes struct {
	Thumbnail Item `json:"thumbnail"`
	Medium Item `json:"medium"`
	SmallPoster Item `json:"small-poster"`
	IndexThumb Item `json:"index-thumb"`
	Full Item `json:"full"`
}

type Item struct {
	SourceUrl string `json:"source_url"`
}


