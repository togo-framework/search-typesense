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
