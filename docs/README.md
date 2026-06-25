# search-typesense — documentation

Typesense driver for togo full-text search

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

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `TYPESENSE_API_KEY` |
| `TYPESENSE_HOST` |
| `TYPESENSE_QUERY_BY` |

## Usage

```go
s := k.Search
s.Index(ctx, "posts", doc)
hits, _ := s.Search(ctx, "posts", "query")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/search-typesense
- Full README: ../README.md
