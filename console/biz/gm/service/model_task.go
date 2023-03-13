package service

type TaskId struct {
	Ids []int `json:"ids"`
}

type Email struct {
	Code   int `json:"code" binding:"required"`
	Number int `json:"number" binding:"required"`
}
