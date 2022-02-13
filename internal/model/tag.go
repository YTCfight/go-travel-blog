package model

type Tag struct {
	*Model
	Name   string `json:"name"`
	Status uint8  `json:"status"`
}

func (a Tag) TableName() string {
	return "blog_name"
}
