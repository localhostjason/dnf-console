package casbinx

type CasbinRule struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	PType     string `json:"p_type" gorm:"column:ptype" description:"策略类型"`
	Role      string `json:"role" gorm:"column:v0" description:"角色"`
	Path      string `json:"path" gorm:"column:v1" description:"api路径 obj"`
	Method    string `json:"method" gorm:"column:v2" description:"访问方法 act"`
	V3        string `json:"-" gorm:"column:v3"`
	V4        string `json:"-" gorm:"column:v4"`
	V5        string `json:"-" gorm:"column:v5" `
	V6        string `json:"-" gorm:"column:v6" `
	V7        string `json:"-" gorm:"column:v7" `
	ApiName   string `json:"api_name" gorm:"-"`
	GroupName string `json:"group_name" gorm:"-"`
	Desc      string `json:"desc" description:"策略描述"`
}
