package data

// 分类
type Class struct {
	ID      int    `json:"id"`
	Picture string `json:"picture"`
	Title   string `json:"title"`
	Text1   string `json:"text1"`
	Text2   string `json:"text2"`
}

var Classes = []Class{
	{ID: 1, Picture: "/可爱电路板.png", Title: "硬件", Text1: "老手艺不能忘，祖宗之法不能丢啊！！", Text2: "硬件小玩具也是玩具"},
	{ID: 2, Picture: "/可爱小猫.png", Title: "嵌入式", Text1: "你为啥学后端了？", Text2: "我绝对不说，我原本是学嵌入式的，被带偏了"},
	{ID: 3, Picture: "/可爱猫猫.png", Title: "后端", Text1: "哇~~这何尝不是又一种泼天的流量呢？", Text2: "我接"},
	{ID: 4, Picture: "/可爱小狗.png", Title: "前端", Text1: "一名后端开发怎么能不会一点儿前端呢？", Text2: "嘿嘿，我的最爱，写小工具可方便了"},
}

// 文章列表
type ListItem struct {
	ID      int    `json:"id"`
	ClassID int    `json:"class_id"`
	Title   string `json:"title"`
}

var List = []ListItem{
	{1, 1, "硬件基础笔记"},
	{2, 1, "基础电路"},
	{3, 1, "数电基础"},
	{4, 1, "示波器使用常识"},
	{5, 2, "嵌入式小项目"},
	{6, 2, "ROTS"},
	{7, 2, "串口协议"},
	{8, 3, "go语言基础"},
	{9, 3, "搭建高并发服务器流程记录"},
	{10, 4, "随手记"},
	{11, 4, "软件小项目"},
	{12, 4, "前端常识"},
}

// 文章详情
type Article struct {
	ID    int      `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
	File  string   `json:"file"`
}

var Articles = []Article{
	{2, "Docker 学习笔记", "2025-08-22", []string{"docker", "devops"}, "docker-notes.md"},
	{3, "测试文章", "2025-08-22", []string{"测试", "文章"}, "example.md"},
	{4, "Hello World", "2025-08-20", []string{"intro", "hello"}, "hello-world.md"},
	{5, "Vue3 基础教程", "2025-08-21", []string{"vue3", "frontend"}, "Vue3教程.md"},
	{1, "搭建 docker 开发环境", "2025-08-21", []string{"docker", "开发环境"}, "搭建docker开发环境.md"},
}
