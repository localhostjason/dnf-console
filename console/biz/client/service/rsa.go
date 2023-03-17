package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
)

type Rsa struct {
	PrivateFile string
	PublicFile  string
}

func NewRsa() *Rsa {
	exeDir, _ := os.Getwd()
	publickeyFile := filepath.Join(exeDir, "source", "publickey.pem")
	privateFile := filepath.Join(exeDir, "source", "private.pem")
	return &Rsa{
		PublicFile:  publickeyFile,
		PrivateFile: privateFile,
	}
}

// GenRsaKey RSA公钥私钥产生
func (r *Rsa) genRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create(r.PrivateFile)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create(r.PublicFile)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
