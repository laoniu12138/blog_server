package model

// Tag ...
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TagSwagger ...
type TagSwagger struct{
	List []*Tag
	Pager *app.Pager
}

// TableName ...
func (a Tag) TableName() string {
	return "blog_tag"
}
