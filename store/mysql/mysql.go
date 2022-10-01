package mysql

import (
	"Skywing/settings"
	"Skywing/store"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"sync"
)

type datastore struct {
	db *sqlx.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}
func (ds *datastore) Roles() store.RoleStore {
	return newRoles(ds)
}

func (ds *datastore) Close() error {
	db := ds.db
	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr(cfg *settings.MySQLConfig) (store.Factory, error) {
	if cfg == nil && mysqlFactory == nil {
		zap.L().Fatal("数据库初始化失败！")
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var dbIns *sqlx.DB
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
		dbIns, err = sqlx.Connect("mysql", dsn)
	})

	mysqlFactory = &datastore{dbIns}

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}
	return mysqlFactory, nil
}
