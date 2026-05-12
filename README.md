# complete-go

This repository uses a root `go.work` workspace to bind two separate apps together, with no shared code between them.

## Layout

- `go.work` - workspace root that includes both modules
- `go_basics/` - the current code, moved into its own app module
- `api_project/` - a new starter app module

## Run

From the repository root, the workspace lets you run either module directly:

```bash
go run ./go_basics
go run ./api_project
```

## Build

From the repository root, the workspace lets you build either module directly:

```bash
go build ./go_basics
go build ./api_project
```

## Notes

- `go_basics` contains the current demo code, refactored to live entirely inside that app.
- `api_project` starts with a `main.go` and `go.mod` only, ready for new work.
