package test_db

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/db"
	"github.com/localhostjason/webserver/server/util/uv"
	"gorm.io/gorm"
)

type TUser struct {
	gorm.Model

	Name string `json:"name" gorm:"not null,unique"`

	Desc string `json:"desc"`

	Languages []TLanguages `gorm:"many2many:user_languages;"`
}

type TLanguages struct {
	gorm.Model

	Name string `json:"name" gorm:"not null"`
}

func init() {
	db.RegTables(&TUser{}, &TLanguages{})
}

func InitTestDbRouter(r *gin.RouterGroup) {
	t := r.Group("db")
	{
		t.GET("", getFunc)
		t.PUT(":id", putFunc)
		t.POST("", postFunc)
		t.DELETE(":id", delFunc)
	}

}

func getFunc(c *gin.Context) {
	var user = make([]TUser, 0)

	db.DB.Preload("Languages").Find(&user)
	c.JSON(200, user)
}

func putFunc(c *gin.Context) {
	uid := uv.PID(c)

	var user TUser
	err := db.DB.First(&user, uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(200, err)
		return
	}

	var body struct {
		Desc      string   `json:"desc"`
		Languages []string `json:"languages" binding:"required"`
	}

	uv.PB(c, &body)

	var lang []TLanguages
	db.DB.Where("name IN ?", body.Languages).Find(&lang)

	user.Desc = body.Desc
	db.DB.Debug().Save(&user)

	db.DB.Model(&user).Association("Languages").Replace(lang)
	c.JSON(200, user)
}

func postFunc(c *gin.Context) {
	var body struct {
		Name      string   `json:"name"  binding:"required"`
		Languages []string `json:"languages" binding:"required"`
	}

	uv.PB(c, &body)

	var lang []TLanguages
	for _, v := range body.Languages {
		lang = append(lang, TLanguages{Name: v})
	}
	u := TUser{
		Name:      body.Name,
		Languages: lang,
	}
	db.DB.Create(&u)
	c.JSON(200, u)
}

func delFunc(c *gin.Context) {
	uid := uv.PID(c)

	var user TUser
	err := db.DB.First(&user, uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(200, err)
		return
	}
	db.DB.Model(&user).Association("Languages").Clear()
	c.JSON(200, user)
}
