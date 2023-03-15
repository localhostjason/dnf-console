package view

import (
	"console/biz/gm/service"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
)

func getAccounts(c *gin.Context) {
	var pi, order, q = &uv.PagingIn{}, &uv.Order{}, &service.AccountFilter{}
	uv.PQ(c, pi, order, q)
	lst, po, err := service.GetAccounts(q, pi, order)
	uv.PEIf(E_ACCOUNT_GET, err)

	c.JSON(200, uv.PagedOut(lst, po))
}

func rechargeAccount(c *gin.Context) {
	uid := uv.PPID(c, "id")

	data := &service.RechargeReq{}
	uv.PB(c, data)

	err := service.RechargeAccount(uid, data)
	uv.PEIf(E_RECHARGE_POST, err)
	c.Status(201)
}

func resetCreateCharac(c *gin.Context) {
	uid := uv.PPID(c, "id")
	err := service.ResetCreateCharac(uid)
	uv.PEIf(E_RESET_CREATE_CHARAC, err)
	c.Status(201)
}

func deleteAccount(c *gin.Context) {
	uid := uv.PPID(c, "id")
	err := service.DeleteAccount(uid)
	uv.PEIf(E_ACCOUNT_DELETE, err)
	c.Status(201)
}

func changePassword(c *gin.Context) {
	args := &service.PasswordReq{}
	uv.PB(c, args)

	uid := uv.PPID(c, "id")
	err := service.ChangeAccountPassword(uid, args)
	uv.PEIf(E_ACCOUNT_CHANGE_PASSWORD, err)
	c.Status(201)
}

func createAccount(c *gin.Context) {
	args := &service.CreateAccountReq{}
	uv.PB(c, args)

	err := service.CreateAccount(args)
	uv.PEIf(E_ACCOUNT_CREARE, err)
	c.Status(201)
}

func updateAccount(c *gin.Context) {
	args := &service.UpdateAccountReq{}
	uv.PB(c, args)

	uid := uv.PPID(c, "id")
	err := service.UpdateAccountInfo(uid, args)
	uv.PEIf(E_ACCOUNT_UPDATE, err)
	c.Status(201)
}
