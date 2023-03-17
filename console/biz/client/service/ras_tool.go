package service

import (
	gorsa "github.com/Lyafei/go-rsa"
)

func applyPriEPubD(data, pirvatekey string) (string, error) {
	prienctypt, err := gorsa.PriKeyEncrypt(data, pirvatekey)
	if err != nil {
		return "", err
	}

	return prienctypt, nil
}
