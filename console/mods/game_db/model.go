package game_db

import (
	"gorm.io/gorm"
	"sync"
)

// DBPool 定义一个类，继承字典(异步,带锁的),一会存入db对象 { key : db }
type DBPool struct {
	sync.Map
}

// Get 为这个 类<对象池> 添加方法，分别为 Get,Add,Del
func (p *DBPool) Get(name string) *gorm.DB {
	if s, ok := p.Load(name); ok {
		return s.(*gorm.DB)
	} else {
		return nil
	}
}

func (p *DBPool) Add(name string, tx *gorm.DB) {
	p.Store(name, tx)
}

func (p *DBPool) Del(name string) {
	p.Delete(name)
}
