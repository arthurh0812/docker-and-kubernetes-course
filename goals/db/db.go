package db

import (
	"fmt"
	"reflect"

	"github.com/go-pg/pg"
)

type ConnectionInfo struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

func (i ConnectionInfo) getEmpty() (empty []string) {
	info := reflect.ValueOf(i)
	for i := 0; i < info.NumField(); i++ {
		field := info.Field(i)
		if field.String() == "" {
			empty = append(empty, info.Type().Field(i).Name)
		}
	}
	return
}

func (i *ConnectionInfo) Validate() error {
	missing := i.getEmpty()
	if len(missing) == 0 {
		return nil
	}
	fieldList := constructList(missing)
	return fmt.Errorf("validating connection information: empty fields: %s", fieldList)
}

func (i *ConnectionInfo) CreateOpts() (*pg.Options, error) {
	err := i.Validate()
	if err != nil {
		return nil, err
	}
	addr := fmt.Sprintf("%s:%s", i.Host, i.Port)
	opts := &pg.Options{
		User:     i.User,
		Password: i.Password,
		Addr:     addr,
		Database: i.Database,
	}
	return opts, nil
}

func Connect(info *ConnectionInfo) (*pg.DB, error) {
	opts, err := info.CreateOpts()
	if err != nil {
		return nil, err
	}
	db := pg.Connect(opts)
	if db == nil {
		return nil, fmt.Errorf("failed to connect to database server")
	}
	return db, nil
}

func constructList(parts []string) string {
	var list string
	for i, p := range parts {
		if i == len(parts)-1 {
			list += fmt.Sprintf("%s.", p)
			continue
		}
		list += fmt.Sprintf("%s, ", p)
	}
	return list
}

