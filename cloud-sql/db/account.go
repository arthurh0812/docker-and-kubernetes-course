package db

import (
	"sync"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// "accounts" table
type Account struct {
	//tableName struct {} `sql:"accounts_collection"`
	ID int64 `sql:"id,pk"`
	Name string `sql:"name,unique"`
	Owner int64 `sql:"owner"`
	Token string `sql:"token"`
	Entities map[string]string `sql:"type:jsonb"`

	mu *sync.Mutex `sql:"-"`
}

func CreateAccountsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&Account{}, opts)
	return err
}