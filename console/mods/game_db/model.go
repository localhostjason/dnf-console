package game_db

import (
	"gorm.io/gorm"
	"sync"
)

// ConnectPool 定义一个类，继承字典(异步,带锁的),一会存入db对象 { key : db }
type ConnectPool struct {
	sync.Map
}

// Get 为这个 类<对象池> 添加方法，分别为 Get,Add,Del
func (p *ConnectPool) Get(name string) *gorm.DB {
	if s, ok := p.Load(name); ok {
		return s.(*gorm.DB)
	} else {
		return nil
	}
}

func (p *ConnectPool) Add(name string, tx *gorm.DB) {
	p.Store(name, tx)
}

func (p *ConnectPool) Del(name string) {
	p.Delete(name)
}
