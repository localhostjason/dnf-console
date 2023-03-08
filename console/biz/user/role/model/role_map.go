package model

type RolePost struct {
	Role   string `json:"role" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required,oneof=PUT POST GET DELETE"`
}
