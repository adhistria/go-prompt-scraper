package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adhistria/scraper-prompt/domain"
	"github.com/adhistria/scraper-prompt/metadata"
	"github.com/jmoiron/sqlx"
)

type MetadataRepository struct {
	sqlClient *sqlx.DB
}

func (m *MetadataRepository) Save(ctx context.Context, md *domain.Metadata) error {
	if md.ID == 0 {
		query := `INSERT INTO metadatas (site, num_of_links, num_of_images, created_at, updated_at)
		VALUES(:site, :num_of_links, :num_of_images, :created_at, :updated_at);`
		_, err := m.sqlClient.NamedExecContext(ctx, query, md)
		if err != nil {
			return fmt.Errorf("unable to save execute query due: %w", err)
		}
		return nil
	} else {
		query := `UPDATE metadatas
				SET num_of_links = :num_of_links , num_of_images = :num_of_images, updated_at = :updated_at
				WHERE id = :id;
				`

		_, err := m.sqlClient.NamedExecContext(ctx, query, md)
		if err != nil {
			return fmt.Errorf("unable to save execute query due: %w", err)
		}
		return nil
	}
}

func (m *MetadataRepository) Update(ctx context.Context, md *domain.Metadata) error {
	return nil
}

func (m *MetadataRepository) FindBySite(ctx context.Context, site string) (*domain.Metadata, error) {
	query := `SELECT * FROM metadatas WHERE site = $1 `
	md := domain.Metadata{}
	if err := m.sqlClient.GetContext(ctx, &md, query, site); err != nil {
		if err == sql.ErrNoRows {
			return nil, metadata.ErrNotFound
		}
		return nil, fmt.Errorf("unable to execute query due: %w", err)
	}
	return &md, nil
}

func NewMetadataRepository(db *sqlx.DB) domain.MetadataDataStore {
	return &MetadataRepository{sqlClient: db}
}
