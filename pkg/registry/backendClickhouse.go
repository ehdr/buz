// Copyright (c) 2022 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/buz/blob/main/LICENSE

package registry

import (
	"github.com/rs/zerolog/log"
	"github.com/silverton-io/buz/pkg/config"
	"github.com/silverton-io/buz/pkg/db"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type ClickhouseSchemaCacheBackend struct {
	gormDb        *gorm.DB
	registryTable string
}

func (b *ClickhouseSchemaCacheBackend) Initialize(conf config.Backend) error {
	connParams := db.ConnectionParams{
		Host: conf.ClickhouseHost,
		Port: conf.ClickhousePort,
		Db:   conf.ClickhouseDbName,
		User: conf.ClickhouseUser,
		Pass: conf.ClickhousePass,
	}
	connString := db.GenerateClickhouseDsn(connParams)
	gormDb, err := gorm.Open(clickhouse.Open(connString), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("🔴 could not open clickhouse connection")
		return err
	}
	b.gormDb, b.registryTable = gormDb, conf.RegistryTable
	ensureErr := db.EnsureTable(b.gormDb, b.registryTable, ClickhouseRegistryTable{})
	return ensureErr
}

func (b *ClickhouseSchemaCacheBackend) GetRemote(schema string) (contents []byte, err error) {
	var s RegistryTable
	b.gormDb.Table(b.registryTable).Where("name = ?", schema).First(&s)
	err = b.gormDb.Error
	if err != nil {
		return nil, err
	}
	return s.Contents, nil
}

func (b *ClickhouseSchemaCacheBackend) Close() {
	log.Info().Msg("🟢 closing clickhouse schema cache backend")
}
