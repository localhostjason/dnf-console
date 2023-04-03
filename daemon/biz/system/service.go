package system

import (
	"bufio"
	"daemon/mods/shell"
	log "github.com/sirupsen/logrus"
	"io"
	"os/exec"
	"sync"
	"time"
)

type DnfService struct {
	StartShell string
	StopShell  string

	Quit chan bool
}

func NewDnfService() *DnfService {
	cfg := getDnfServiceConfig()
	return &DnfService{
		StartShell: cfg.StartShell,
		StopShell:  cfg.StopShell,
		Quit:       make(chan bool),
	}
}

func (d *DnfService) Run() error {
	_, _ = shell.Exec(d.StopShell)
	time.Sleep(500 * time.Millisecond)
	_, _ = shell.Exec(d.StopShell)
	time.Sleep(500 * time.Millisecond)

	err := d.shellStart(d.StartShell)
	if err != nil {
		return err
	}

	return nil
}

func (d *DnfService) Stop() {
	d.Quit <- true
	data, _ := shell.Exec(d.StopShell)
	time.Sleep(500 * time.Millisecond)
	log.Info(data)
	data, _ = shell.Exec(d.StopShell)
	log.Info(data)
}

func (d *DnfService) shellStart(cmd string) error {
	c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			select {
			case <-d.Quit:
				return

			default:
				readString, err := reader.ReadString('\n')
				if err != nil || err == io.EOF {
					return
				}
				log.Info(readString)
			}

		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}
