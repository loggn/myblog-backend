package model

// 分类
type Class struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Picture string `json:"picture"`
	Title   string `json:"title"`
	Text    string `json:"text"`
}

// 文章详情
type Article struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	ClassID int    `json:"class_id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Tags    string `json:"tags"` // 存储为逗号分隔字符串，例如 "docker,devops"
	File    string `json:"file"`
}
