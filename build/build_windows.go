package main

import (
	"bytes"
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmds := []string{
		"cd ..\\web && npm i",
		"cd ..\\web && npm run build",
		"cd ..\\dist\\web && rmdir /s /q static",
		"cd ..\\dist\\web && mkdir static",
		"cd ..\\web && xcopy /s /e /i dist ..\\dist\\web\\static /Y",

		"cd ..\\console && go build main.go",
		"cd ..\\console && copy main.exe ..\\dist\\main.exe /Y && del main.exe",
	}

	for _, cmd := range cmds {
		out, err := WindowsExecByShell(cmd)
		if err != nil {
			log.Error(out)
		}
		log.Info("OUT:", out)
	}
}

func WindowsExecByShell(cmd string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	command := exec.CommandContext(ctx, "cmd", "/C", cmd)
	stdOut := bytes.Buffer{}
	stderr := bytes.Buffer{}
	command.Stdout = &stdOut
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		s, _ := GbkToUtf8(stderr.Bytes())
		return string(s), errors.New(stderr.String())
	}
	return stdOut.String(), nil
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
