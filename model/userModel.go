package model

type User struct {
	Id       int64 `json:"id,omitempty"`
	Username string
	Password string
	Nickname string
	Token    string
}
