package service

type LoginReq struct {
	LoginType string `json:"login_type"`
	Uid       int    `json:"uid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
