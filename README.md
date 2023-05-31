# WIP http police api

## Usage

## Requirements

- Node.js
- [wrangler](https://developers.cloudflare.com/workers/wrangler/)
  - just run `npm install -g wrangler`
- tinygo (0.28 or haigher)
  - [Installation](https://tinygo.org/getting-started/)

## Setup

### Cludflare setup

```
make create-db
```

## Development

### Commands

```
make dev     # run dev server
make build   # build Go Wasm binary
make deploy # deploy worker
```

if you change `app` directory, you should run `make generate` before `make dev`.

### Testing dev server

- Just send HTTP request using some tools like curl.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"domain":"tsurutatakumi.info"}' http://localhost:8787

# However, it currently does not work well with workers
failed to check: operation not implemented
```
