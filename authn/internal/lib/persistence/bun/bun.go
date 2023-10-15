package bun

import (
	"database/sql"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Bun struct {
}

func NewBun() *Bun {
	return &Bun{}
}

func (b *Bun) Close() {

}

func (b *Bun) Connect() {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
}
