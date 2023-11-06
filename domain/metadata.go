package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Metadata struct {
	ID          int       `db:"id" json:"-"`
	NumOfLinks  int       `db:"num_of_links" json:"num_of_links"`
	Site        string    `db:"site" json:"site"`
	NumOfImages int       `db:"num_of_images" json:"num_of_images"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (m *Metadata) Print() {
	data, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(data))
}

type MetadataService interface {
	Fetch(ctx context.Context, urls []string) error
	FetchDetail(ctx context.Context, url string) error
}

type MetadataDataStore interface {
	Save(context.Context, *Metadata) error
	FindBySite(context.Context, string) (*Metadata, error)
	Update(context.Context, *Metadata) error
}

type MetadataScraper interface {
	GetMetadata(context.Context, string) (*Metadata, error)
}
