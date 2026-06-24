// Package typesense is a Typesense driver for togo full-text search.
// Blank-import it and set SEARCH_DRIVER=typesense, TYPESENSE_HOST,
// TYPESENSE_API_KEY (and TYPESENSE_QUERY_BY for the searchable fields).
package typesense

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/togo-framework/search"
	"github.com/togo-framework/togo"
)

func init() {
	search.RegisterDriver("typesense", func(k *togo.Kernel) (search.Searcher, error) {
		host := strings.TrimRight(os.Getenv("TYPESENSE_HOST"), "/")
		key := os.Getenv("TYPESENSE_API_KEY")
		if host == "" || key == "" {
			return nil, errors.New("search-typesense: TYPESENSE_HOST and TYPESENSE_API_KEY required")
		}
		qb := os.Getenv("TYPESENSE_QUERY_BY")
		if qb == "" {
			qb = "title,body,name"
		}
		return &searcher{host: host, key: key, queryBy: qb, client: &http.Client{Timeout: 15 * time.Second}}, nil
	})
}

type searcher struct {
	host, key, queryBy string
	client             *http.Client
}

func (s *searcher) do(ctx context.Context, method, path string, body any) (*http.Response, error) {
	var r io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		r = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, s.host+path, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-TYPESENSE-API-KEY", s.key)
	req.Header.Set("Content-Type", "application/json")
	return s.client.Do(req)
}

func (s *searcher) Index(ctx context.Context, index, id string, doc map[string]any) error {
	d := map[string]any{"id": id}
	for k, v := range doc {
		d[k] = v
	}
	resp, err := s.do(ctx, http.MethodPost, "/collections/"+url.PathEscape(index)+"/documents?action=upsert", d)
	if err != nil {
		return err
	}
	return drain(resp)
}

func (s *searcher) Search(ctx context.Context, index, query string, limit int) ([]search.Hit, error) {
	if limit <= 0 {
		limit = 20
	}
	q := url.Values{"q": {query}, "query_by": {s.queryBy}, "per_page": {strconv.Itoa(limit)}}
	resp, err := s.do(ctx, http.MethodGet, "/collections/"+url.PathEscape(index)+"/documents/search?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("search-typesense: %s: %s", resp.Status, b)
	}
	var out struct {
		Hits []struct {
			Document  map[string]any `json:"document"`
			TextMatch float64        `json:"text_match"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	hits := make([]search.Hit, 0, len(out.Hits))
	for _, h := range out.Hits {
		id, _ := h.Document["id"].(string)
		hits = append(hits, search.Hit{ID: id, Score: h.TextMatch, Doc: h.Document})
	}
	return hits, nil
}

func (s *searcher) Delete(ctx context.Context, index, id string) error {
	resp, err := s.do(ctx, http.MethodDelete, "/collections/"+url.PathEscape(index)+"/documents/"+url.PathEscape(id), nil)
	if err != nil {
		return err
	}
	return drain(resp)
}

func drain(resp *http.Response) error {
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("search-typesense: %s: %s", resp.Status, b)
	}
	io.Copy(io.Discard, resp.Body)
	return nil
}
