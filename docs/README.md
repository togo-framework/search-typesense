# search-typesense — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package typesense is a Typesense driver for togo full-text search.
Blank-import it and set SEARCH_DRIVER=typesense, TYPESENSE_HOST,
TYPESENSE_API_KEY (and TYPESENSE_QUERY_BY for the searchable fields).

## Install

```bash
togo install togo-framework/search-typesense
```

Set `SEARCH_DRIVER=typesense`.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `TYPESENSE_API_KEY` | _see provider docs_ |
| `TYPESENSE_HOST` | _see provider docs_ |
| `TYPESENSE_QUERY_BY` | _see provider docs_ |

## Usage

```go
s := k.Search
s.Index(ctx, "posts", doc)
hits, _ := s.Search(ctx, "posts", "query")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/search-typesense
- README: ../README.md
