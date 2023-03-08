package cmds

import (
	"console/biz/view"
)

func Run() {
	s := NewMainServer()

	s.LoadView(view.SetView)

	s.Run()
}
