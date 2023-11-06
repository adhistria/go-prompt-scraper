package service

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/adhistria/go-prompt-scraper/domain"
	"github.com/adhistria/go-prompt-scraper/metadata"
)

type service struct {
	scraper domain.MetadataScraper
	storage domain.MetadataDataStore
}

func (s *service) Fetch(ctx context.Context, urls []string) error {
	err := s.validateUrl(ctx, urls...)
	if err != nil {
		return err
	}
	for _, url := range urls {

		md, err := s.storage.FindBySite(ctx, metadata.ConvertUrlToPath(url))
		if err != nil {
			if err == metadata.ErrNotFound {
				md = &domain.Metadata{}
				startTimne := time.Now()
				md.CreatedAt = startTimne
				md.UpdatedAt = startTimne
			} else {
				return err
			}
		}

		metadata, err := s.scraper.GetMetadata(ctx, url)
		if err != nil {
			return err
		}

		md.NumOfImages = metadata.NumOfImages
		md.NumOfLinks = metadata.NumOfLinks
		md.UpdatedAt = time.Now()
		md.Site = url

		err = s.storage.Save(ctx, md)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) FetchDetail(ctx context.Context, url string) error {
	fmt.Println("fetcg detauk")
	s.validateUrl(ctx, url)
	metadata, err := s.storage.FindBySite(ctx, metadata.ConvertUrlToPath(url))

	if err != nil {
		return err
	}
	metadata.Print()
	return nil
}

func (s *service) validateUrl(ctx context.Context, urls ...string) error {
	for _, u := range urls {
		_, err := url.ParseRequestURI(u)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewService(scr domain.MetadataScraper, str domain.MetadataDataStore) domain.MetadataService {
	s := &service{
		scraper: scr,
		storage: str,
	}
	return s
}
