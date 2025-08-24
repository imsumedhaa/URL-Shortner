package postgres

import (
	"fmt"
	"database/sql"
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

func (p *Postgres) CreateShortUrl(originalUrl, shortURL string) error {

	if originalUrl == "" || shortURL == "" {
		return fmt.Errorf("original url or ShortUrl cannot be empty")
	}
	

	query := `INSERT INTO urls (original_url, short_url)
          VALUES ($1, $2)
          ON CONFLICT (original_url) DO UPDATE SET short_url = EXCLUDED.short_url`

	
	return nil
}
