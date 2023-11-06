package rest

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/adhistria/go-prompt-scraper/domain"
)

type ScraperRepository struct{}

func (sr *ScraperRepository) GetMetadata(ctx context.Context, url string) (*domain.Metadata, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*5))
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}
	c := http.Client{}
	res, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var buff bytes.Buffer

	body := io.TeeReader(res.Body, &buff)

	html, err := sr.fetchHTML(body)
	if err != nil {
		return nil, err
	}

	err = sr.saveHTML(url, html)
	if err != nil {
		return nil, err
	}

	md, err := sr.fetchMetadata(&buff)
	if err != nil {
		return nil, err
	}

	return md, nil
}

func (s *ScraperRepository) fetchHTML(body io.Reader) ([]byte, error) {
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return []byte(string(content)), nil
}

func (s *ScraperRepository) saveHTML(url string, html []byte) error {
	fileName := url
	if strings.Contains(url, "https") {
		fileName = strings.Replace(url, "https://", "", -1)
	} else {
		fileName = strings.Replace(url, "http://", "", -1)
	}
	fileName = strings.Replace(fileName, "/", "_", -1)
	err := ioutil.WriteFile(fileName+".html", html, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *ScraperRepository) fetchMetadata(body io.Reader) (*domain.Metadata, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	total_links := len(doc.Find("a").Nodes)
	total_images := len(doc.Find("img").Nodes)
	md := &domain.Metadata{
		NumOfLinks:  total_links,
		NumOfImages: total_images,
	}
	return md, nil
}

func NewScraperRepository() domain.MetadataScraper {
	return &ScraperRepository{}
}
