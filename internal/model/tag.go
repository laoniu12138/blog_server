package model

// Tag ...
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TableName ...
func (a Tag) TableName() string {
	return "blog_tag"
}
