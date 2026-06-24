<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/search-typesense</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/search-typesense"><img src="https://pkg.go.dev/badge/github.com/togo-framework/search-typesense.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/search-typesense
```

<!-- /togo-header -->

# search-typesense

> A **Typesense** driver for [togo](https://to-go.dev) full-text search.

Registers the `typesense` driver on the [`search`](https://github.com/togo-framework/search) plugin.

## Install

```sh
togo install togo-framework/search-typesense
```

Then in `.env`:

```sh
SEARCH_DRIVER=typesense
TYPESENSE_HOST=https://xxx.a1.typesense.net
TYPESENSE_API_KEY=your-key
TYPESENSE_QUERY_BY=title,body,name   # the collection's queryable fields
```

> Typesense requires the searchable fields via `query_by` — set `TYPESENSE_QUERY_BY`
> to your collection's text fields (defaults to `title,body,name`).

## License

MIT

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
