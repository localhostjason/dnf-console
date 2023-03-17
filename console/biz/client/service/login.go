package service

import (
	gmModel "console/biz/gm/model"
	"console/mods/game_db"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
)

func ToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Hex2bin(raw string) string {
	result, _ := hex.DecodeString(raw)
	return string(result)
}

// 10转16进制，不够补0
func uidTo16(uid int) string {
	return fmt.Sprintf("%08x", uid)
}

func Login(data *LoginReq) (string, error) {
	dbx := game_db.DBPools.Get(gmModel.DTaiwan)

	var account gmModel.Accounts

	var err error
	if data.LoginType == "uid" {
		err = dbx.Where("UID = ?", data.Uid).First(&account).Error
	} else {
		err = dbx.Where("accountname = ? AND password = ?", data.Username, ToMd5(data.Password)).First(&account).Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("账号未找到！")
	}

	uid16 := uidTo16(account.Uid)
	fmt.Println("data", uid16)

	rsa := NewRsa()
	key := fmt.Sprintf("%s010101010101010101010101010101010101010101010101010101010101010155914510010403030101", uid16)
	d := Hex2bin(key)

	privateKey, _ := os.ReadFile(rsa.PrivateFile)
	token, err := applyPriEPubD(d, string(privateKey))
	fmt.Println("err", err)
	if err != nil {
		return "", err
	}
	fmt.Println("key:", d)
	//token := base64.StdEncoding.EncodeToString(hashToken)
	fmt.Println("token:", token)
	return token, nil
}
