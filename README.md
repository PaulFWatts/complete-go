# complete-go

This repository now uses a Go workspace with two separate apps and no shared code between them.

## Layout

- `go.work` - workspace root
- `go_basics/` - the current code, moved into its own app module
- `api_project/` - a new starter app module

## Run

From the repository root:

```bash
go run ./go_basics
go run ./api_project
```

## Build

From the repository root:

```bash
go build ./go_basics
go build ./api_project
```

## Notes

- `go_basics` contains the current demo code, refactored to live entirely inside that app.
- `api_project` starts with a `main.go` and `go.mod` only, ready for new work.
