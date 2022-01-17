package model

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
	Company  int    `json:"company"`
}
