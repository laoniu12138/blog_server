package model

// Article ...
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct{
	List []*Article
	Pager *app.Pager
}

// TableName ...
func (a Article) TableName() string {
	return "blog_article"
}
