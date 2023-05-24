package pg

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func NewDatabase(url string) *pg.DB {
	orm.SetTableNameInflector(func(s string) string {
		return s
	})

	opt, err := pg.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}

	return pg.Connect(opt)
}
