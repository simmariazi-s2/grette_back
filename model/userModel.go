package model

type User struct {
	UserNo   int    `json:userNo`
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
	Company  int    `json:"company"`
}
