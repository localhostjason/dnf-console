package model

type Task struct {
	CharacNo     int `json:"charac_no"`
	Play1        int `json:"play_1" gorm:"column:play_1;"`
	Play1Trigger int `json:"play_1_trigger" gorm:"column:play_1_trigger;"`
}
