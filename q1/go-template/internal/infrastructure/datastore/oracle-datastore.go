package datastore

import (
	"fmt"
	"go-template/pkg/config"
	"log"

	"github.com/jmoiron/sqlx"

	"database/sql"

	_ "github.com/sijms/go-ora/v2"
)

type OracleDataStore struct {
	DB              *sqlx.DB
	CloseConnection func()
}

func NewOracleDatastore(config config.OracleConfig) OracleDataStore {
	var oracleConnector OracleDataStore
	oracleConnector.connect(config)
	return oracleConnector
}

func (p *OracleDataStore) connect(config config.OracleConfig) {
	connString := fmt.Sprintf("oracle://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open("oracle", connString)
	if err != nil {
		log.Panic(fmt.Errorf("NewOracleDatastore: %w", err))
	}

	p.DB = sqlx.NewDb(db, "oracle")

	// test connection
	err = p.DB.Ping()
	if err != nil {
		log.Panic(fmt.Errorf("NewOracleDatastore: %w", err))
	}

	p.DB.SetConnMaxLifetime(0)
	p.DB.SetMaxIdleConns(0)

	p.CloseConnection = func() {
		db.Close()
		log.Println("Database connection closed")
	}

	log.Println(fmt.Sprintf("%s, DB: %s", "Successfully connected to postgresql", config.DBName))
}
