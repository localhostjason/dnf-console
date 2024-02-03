package service

import (
	"console/biz/gm/model"
	logModel "console/biz/log/model"
	"console/mods/game_db"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/localhostjason/webserver/server/util/uv"
	"gorm.io/gorm"
	"time"
)

func GetAccounts(q *AccountFilter, pi *uv.PagingIn, order *uv.Order) ([]AccountResult, *uv.PagingOut, error) {
	dbx := game_db.DBPools.Get(model.DTaiwan)
	tx := q.FilterQuery(dbx)

	var lst = make([]model.Accounts, 0)
	po, err := uv.PagingFind(tx, &lst, pi, order)

	var data = make([]AccountResult, 0)
	for _, info := range lst {
		roles := getGameRolesByUid(info.Uid)

		if q.HasRoles && roles == 0 {
			continue
		}

		money, capacity := getGameMoneyByUid(info.Uid)
		ceraPoint, cera := getCashByUid(info.Uid)

		data = append(data, AccountResult{
			Accounts:  info,
			Roles:     roles,
			Money:     money,
			Capacity:  capacity,
			CeraPoint: ceraPoint,
			Cera:      cera,
		})
	}

	return data, po, err
}

func getGameRolesByUid(uid int) int64 {
	dbx := game_db.DBPools.Get(model.TaiwanCain)

	var total int64
	dbx.Table("charac_info").Where("m_id = ? AND delete_flag = ?", uid, 0).Count(&total)
	return total
}

func getGameMoneyByUid(uid int) (int, int) {
	dbx := game_db.DBPools.Get(model.TaiwanCain)
	type _R struct {
		Money    int `json:"money"`
		Capacity int `json:"capacity"`
	}

	var result _R
	err := dbx.Debug().Table("account_cargo").Where("m_id = ?", uid).Take(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, 0
	}

	return result.Money, result.Capacity
}

func getCashByUid(uid int) (int, int) {
	dbx := game_db.DBPools.Get(model.TaiwanBilling)

	type CashPoint struct {
		CeraPoint int `gorm:"column:cera_point"`
	}
	var data CashPoint
	dbx.Table("cash_cera_point").Where("account = ?", uid).Take(&data)

	type CashCera struct {
		Cera    int
		ModTran int
	}
	var d CashCera
	dbx.Table("cash_cera").Where("account = ?", uid).Take(&d)

	return data.CeraPoint, d.Cera
}

/*
1. cash_cera_point 充值D点
1. cash_cera 充值D币
*/

func RechargeAccount(uid int, data *RechargeReq, c *gin.Context) error {
	dbx := game_db.DBPools.Get(model.TaiwanBilling)
	var err error
	action := logModel.RechargeCera
	if data.CeraOption == "cera" {
		err = dbx.Table("cash_cera").Where("account = ?", uid).Updates(map[string]interface{}{
			"cera": data.Cera,
		}).Error
	} else {
		err = dbx.Table("cash_cera_point").Where("account = ?", uid).Updates(map[string]interface{}{
			"cera_point": data.CeraPoint,
		}).Error

		action = logModel.RechargeCeraPoint
	}

	if err != nil {
		return err
	}

	logModel.CreateRechargeLog(uid, action, data.Number, c)
	return nil
}

// ResetCreateCharac 重置 创建 角色
func ResetCreateCharac(uid int) error {
	dbx := game_db.DBPools.Get(model.DTaiwan)

	var data = struct{}{}
	dbx.Table("limit_create_character").Where("m_id = ?", uid).Delete(&data)

	// 创建一个 uid = 0

	now := time.Now()
	err := dbx.Table("limit_create_character").Where("m_id = ?", uid).Create(map[string]interface{}{
		"m_id":             uid,
		"count":            0,
		"last_access_time": now.Format("2006-01-02 15:04:05"),
	}).Error

	return err
}

// DeleteAccount 删除账号
func DeleteAccount(uid int) error {
	dbx := game_db.DBPools.Get(model.DTaiwan)

	var account model.Accounts
	err := dbx.Where("UID = ?", uid).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("账号不存在")
	}

	roles := getGameRolesByUid(uid)
	if roles != 0 {
		return errors.New("该账号存在角色，无法删除")
	}
	dbx.Delete(&account)

	var result map[string]interface{}
	dbx.Table("member_info").Where("m_id = ?", uid).Delete(&result)
	dbx.Table("member_white_account").Where("m_id = ?", uid).Delete(&result)

	dbBilling := game_db.DBPools.Get(model.TaiwanBilling)
	dbBilling.Table("cash_cera").Where("account = ?", uid).Delete(&result)
	dbBilling.Table("cash_cera_point").Where("account = ?", uid).Delete(&result)

	dbx2d := game_db.DBPools.Get(model.TaiwanCain2nd)
	dbx2d.Table("member_avatar_coin").Where("m_id = ?", uid).Delete(&result)

	dbLogin := game_db.DBPools.Get(model.TaiwanLogin)
	dbLogin.Table("member_login").Where("m_id = ?", uid).Delete(&result)
	return nil
}

func ChangeAccountPassword(uid int, args *PasswordReq) error {
	dbx := game_db.DBPools.Get(model.DTaiwan)

	var account model.Accounts
	err := dbx.Where("UID = ?", uid).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("账号不存在")
	}

	account.Password = ToMd5(args.Password)
	dbx.Save(&account)
	return nil
}

func CreateAccount(args *CreateAccountReq) error {
	dbx := game_db.DBPools.Get(model.DTaiwan)
	var account model.Accounts
	err := dbx.Where("accountname = ?", args.AccountName).First(&account).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("账号已存在")
	}

	newAccount := &model.Accounts{
		AccountName: args.AccountName,
		Password:    ToMd5(args.Password),
		QQ:          args.QQ,
	}
	dbx.Debug().Create(newAccount)

	uid := newAccount.Uid
	dbx.Table("member_info").Create(map[string]interface{}{
		"m_id":    uid,
		"user_id": uid,
	})
	dbx.Table("member_white_account").Create(map[string]interface{}{
		"m_id": uid,
	})

	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05")
	dbBilling := game_db.DBPools.Get(model.TaiwanBilling)
	dbBilling.Table("cash_cera").Create(map[string]interface{}{
		"account":  uid,
		"cera":     1000,
		"mod_tran": 0,
		"mod_date": nowStr,
		"reg_date": nowStr,
	})
	dbBilling.Table("cash_cera_point").Create(map[string]interface{}{
		"account":    uid,
		"cera_point": 50,
		"mod_date":   nowStr,
		"reg_date":   nowStr,
	})

	dbx2d := game_db.DBPools.Get(model.TaiwanCain2nd)
	dbx2d.Table("member_avatar_coin").Create(map[string]interface{}{
		"m_id": uid,
	})

	dbLogin := game_db.DBPools.Get(model.TaiwanLogin)
	dbLogin.Table("member_login").Create(map[string]interface{}{
		"m_id": uid,
	})

	return nil
}

func UpdateAccountInfo(uid int, args *UpdateAccountReq) error {
	dbx := game_db.DBPools.Get(model.DTaiwan)
	var account model.Accounts
	err := dbx.Where("UID = ?", uid).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("账号不存在")
	}

	account.QQ = args.QQ
	return dbx.Save(&account).Error
}
