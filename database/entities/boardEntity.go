package entities

import "time"

type Board struct {
	BoardNo    int        `gorm:"primaryKey;autoIncrement;not null;column:boardNo"`
	Title      string     `gorm:"not null;column:title"`
	Content    string     `gorm:"not null;column:content"`
	CategoryNo int        `gorm:"not null;column:categoryNo"`
	Like       int        `gorm:"not null;column:like;default:0"`
	Dislike    int        `gorm:"not null;column:dislike;default:0"`
	Report     int        `gorm:"not null;column:report;default:0"`
	UserNo     int        `gorm:"not null;column:userNo"`
	IsUsed     int        `gorm:"not null;column:isUsed;default:1"`
	CreateDtm  *time.Time `gorm:"not null;column:createDtm;default:current_timestamp"`
	UpdateDtm  time.Time  `gorm:"column:updateDtm"`
}

type Reply struct {
	ReplyNo   int        `gorm:"primaryKey;autoIncrement;not null;column:replyNo"`
	Content   string     `gorm:"not null;column:content"`
	Report    int        `gorm:"not null;column:report;default:0"`
	BoardNo   int        `gorm:"not null;column:boardNo"`
	UserNo    int        `gorm:"not null;column:userNo"`
	Like      int        `gorm:"not null;column:like;default:0"`
	Dislike   int        `gorm:"not null;column:dislike;default:0"`
	IsUsed    int        `gorm:"not null;column:isUsed;default:1"`
	CreateDtm *time.Time `gorm:"not null;column:createDtm;default:current_timestamp"`
	UpdateDtm time.Time  `gorm:"column:updateDtm"`
}

type Recommand struct {
	RecNo     int        `gorm:"primaryKey;autoIncrement;not null;column:recNo"`
	UserNo    *int       `gorm:"not null;column:userNo"`
	BoardNo   *int       `gorm:"not null;column:boardNo"`
	ReplyNo   *int       `gorm:"not null;column:replyNo"`
	LikeType  int        `gorm:"not null;column:likeType"`
	CreateDtm *time.Time `gorm:"not null;column:createDtm;default:current_timestamp"`
}

type Category struct {
	CategoryNo   int    `gorm:"primaryKey;autoIncrement;not null;column:categoryNo"`
	CategoryName string `gorm:"column:categoryName"`
}

func (r *Recommand) SetUserNo(userNo int) {
	r.UserNo = &userNo
}

func (r *Recommand) SetBoardNo(boardNo int) {
	r.BoardNo = &boardNo
}

func (r *Recommand) SetReplyNo(replyNo int) {
	r.ReplyNo = &replyNo
}
