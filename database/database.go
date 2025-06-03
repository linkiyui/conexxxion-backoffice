package database

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/conexxxion/conexxxion-backoffice/config"
)

var pgConn *pgxpool.Pool

func NewConnection() *pgxpool.Pool {
	c := config.GetConfig()
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Password, c.Database.DatabaseName)
	pgPoolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("postgres_db | NewConnection() | parsing config: ", err.Error())
	}
	pgPoolConfig.MaxConns = c.MaxConnections
	conn, err := pgxpool.NewWithConfig(context.Background(), pgPoolConfig)
	if err != nil {
		log.Fatal("postgres_db | NewConnection() | ", err.Error())
	}

	return conn
}

func Get() *pgxpool.Pool {
	if pgConn == nil {
		pgConn = NewConnection()
	}
	return pgConn
}
