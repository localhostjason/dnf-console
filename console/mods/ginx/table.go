package ginx

import "github.com/localhostjason/webserver/db"

type Authz struct {
	ID        int    `json:"id"`
	GroupName string `json:"group_name"`
	ApiName   string `json:"api_name"`
	Url       string `json:"url" description:"对象 obj"`
	Method    string `json:"method" description:"act"`
}

func init() {
	db.RegTables(&Authz{})
}
