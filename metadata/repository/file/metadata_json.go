package file

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/adhistria/go-prompt-scraper/domain"
	"github.com/adhistria/go-prompt-scraper/metadata"
)

type MetadataJSONRepository struct{}

func (m *MetadataJSONRepository) Save(ctx context.Context, md *domain.Metadata) error {
	file, err := json.MarshalIndent(md, "", " ")
	if err != nil {
		return fmt.Errorf("can't marshal json %w", err)
	}

	err = ioutil.WriteFile(metadata.ConvertUrlToPath(md.Site)+".json", file, 0644)
	if err != nil {
		return fmt.Errorf("unable to save json file due: %w", err)
	}
	return nil
}

func (m *MetadataJSONRepository) FindBySite(ctx context.Context, site string) (*domain.Metadata, error) {
	file, err := ioutil.ReadFile(site + ".json")
	if err != nil {
		return nil, metadata.ErrNotFound
	}
	md := domain.Metadata{}

	err = json.Unmarshal([]byte(file), &md)
	if err != nil {
		return nil, err
	}
	return &md, nil
}

func (m *MetadataJSONRepository) Update(ctx context.Context, md *domain.Metadata) error {
	return m.Save(ctx, md)
}

func NewMetadataJSONRepository() domain.MetadataDataStore {
	return &MetadataJSONRepository{}
}
