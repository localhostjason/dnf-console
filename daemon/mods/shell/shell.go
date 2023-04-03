package shell

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Exec(cmd string) (string, error) {
	var c, err = getConfig()
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to get shell exec timeout config:%v", err))
	}
	return ExecTimeout(cmd, time.Duration(c.Timeout)*time.Second)
}

func ExecTimeout(cmd string, timeout time.Duration) (string, error) {
	fmt.Println(cmd)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	command := exec.CommandContext(ctx, "bash", "-c", cmd)

	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr
	err := command.Run()

	strOut := strings.TrimSpace(out.String())
	if err != nil {
		return strOut, errors.New(stderr.String())
	}
	return strOut, nil
}
