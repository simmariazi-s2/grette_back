package model

type Board struct {
	BoardNo      int    `json:"boardNo"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Like         int    `json:"like"`
	Dislike      int    `json:"disLike"`
	UserNo       int    `json:"userNo"`
	IsUsed       int    `json:"isUsed"`
	CategoryNo   int    `json:"categoryNo"`
	CategoryName string `json:"categoryName"`
}

type Reply struct {
	ReplyNo int    `json:"replyNo"`
	Content string `json:"content"`
	Report  int    `json:"report"`
	BoardNo int    `json:"boardNo"`
	UserNo  int    `json:"userNo"`
	Like    int    `json:"like"`
	Dislike int    `json:"dislLike"`
	IsUsed  int    `json:"isUsed"`
}
