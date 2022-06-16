package tool

import (
	"fmt"
	"time"
	"xorm.io/xorm"
)

type DB struct {
	Dsn             string        `koanf:"dsn"`
	MigrationPath   string        `koanf:"migration_path"`
	MaxIdleConns    int           `koanf:"max_idle_conns"`
	MaxOpenConns    int           `koanf:"max_open_conns"`
	ConnMaxLifeTime time.Duration `koanf:"conn_max_life_time"`
	ConnMaxIdleTime time.Duration `koanf:"conn_max_idle_time"`
}

func NewDB(dbCfg DB) (*xorm.Engine, error) {
	db, err := xorm.NewEngine("postgres", dbCfg.Dsn)
	if err != nil {
		return nil, fmt.Errorf("data set up error: %s, dsn: %s", err, dbCfg.Dsn)
	}
	db.ShowSQL(true)
	db.SetMaxIdleConns(dbCfg.MaxIdleConns)
	db.SetMaxOpenConns(dbCfg.MaxOpenConns)
	db.SetConnMaxLifetime(dbCfg.ConnMaxLifeTime)
	//db.SetConnMaxIdleTime(dbCfg.ConnMaxIdleTime)
	db.TZLocation = time.UTC
	db.DatabaseTZ = time.UTC

	return db, nil
}
