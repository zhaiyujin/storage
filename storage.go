/**
 * package: gfoss
 * author: zhaiyujin
 * description:
 */
package storage

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"sync"
)

type Storage struct {
}

var once sync.Once
var s Storager
var driver string
var serverRoot string

func NewStorage() Storager {

	if driver != g.Cfg().GetString("storage.driver", "local") {
		s=nil
	}

	once.Do(func() {
		var err error
		//s = &Storager{}
		driver = g.Cfg().GetString("storage.driver", "local")
		serverRoot = gfile.Join(g.Cfg().GetString("server.serverRoot"), "/")

		switch driver {
		case "alioss":
			if s, err = NewAliOss(); err != nil {
				panic(err.Error())
			}
			break
		case "aws":
		case "local":
			if s, err = NewLocalFile(); err != nil {
				panic(err.Error())
			}
		default:
			if s, err = NewLocalFile(); err != nil {
				panic(err.Error())
			}
		}

	})
	return s
}
