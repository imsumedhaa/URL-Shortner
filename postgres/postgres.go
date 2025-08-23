package postgres

import (
	"fmt"

	"github.com/imsumedhaa/In-memory-database/database"
	"github.com/imsumedhaa/In-memory-database/postgres"
)

type Postgres struct {
	dbclient database.Database
}

func NewPostgres(host, port, username, password, dbname string) (*Postgres, error) {

	dbClient, err := postgres.NewPostgres(host, port, username, password, dbname)

	if err != nil {
		return nil, fmt.Errorf("failed to connect %w", err)
	}

	return &Postgres{dbclient: dbClient}, nil
}

func (p *Postgres) GetShortUrl(originalUrl string)( error){

	if originalUrl == ""{
		return fmt.Errorf("original url cannot be empty")
	}

	
	return nil
}
