package service

type StatCountResult struct {
	UserTotal      int `json:"user_total"`
	CharacTotal    int `json:"charac_total"`
	CeraTotal      int `json:"cera_total"`
	CeraPointTotal int `json:"cera_point_total"`
}

type TopInfo struct {
	Uid         int    `json:"uid"`
	AccountName string `json:"account_name"`
	Total       int    `json:"total"`
}

type StatTop5Result struct {
	Cera           []TopInfo `json:"cera"`
	CeraPoint      []TopInfo `json:"cera_point"`
	CeraTotal      int       `json:"cera_total"`
	CeraPointTotal int       `json:"cera_point_total"`
}
