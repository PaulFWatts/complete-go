# Copilot Instructions for complete-go

## Quick Start

### Run Applications
From the repository root:
```bash
go run ./go_basics    # Run the Go basics demo app
go run ./api_project  # Run the API project starter app
```

### Build Applications
```bash
go build ./go_basics    # Build go_basics executable
go build ./api_project  # Build api_project executable
```

### Go Version
Go 1.26.2 (specified in `go.work` and individual `go.mod` files)

## Architecture

This repository uses a **Go workspace** with two independent modules:

- **`go_basics/`** — Educational demo code covering Go fundamentals (variables, arrays, slices, maps, structs, loops, control flow). Each concept area is organized into its own package: `gobasics`, `loops`, `structs`, `library`.
  
- **`api_project/`** — A minimal starter module with only `main.go` and `go.mod`. This is the designated module for new feature development.

**Key Design**: The two modules are completely independent with no shared code. The `go.work` file at the repository root manages both as a workspace.

## Code Conventions

### Module Paths
- Full module paths use the pattern: `github.com/PaulFWatts/complete_go/{module_name}`
- Internal imports within a module use the full path (e.g., `github.com/PaulFWatts/complete_go/go_basics/gobasics`)

### Naming
- **Package names**: lowercase, match directory names (`gobasics`, `loops`, `structs`, `library`)
- **Entry point functions**: Use the `Run*` pattern (e.g., `RunBasics()`, `RunLoops()`, `RunStructs()`) — these are top-level functions called from `main()`
- **Exported names**: PascalCase for types and functions
- **Unexported names**: camelCase for private variables and helper functions

### Documentation
- All exported functions and types should have doc comments in the format: `// FunctionName description` or `// TypeName description`
- Doc comments appear on the line immediately before the declaration

### Struct Usage
- Use named field syntax in struct literals: `Person{Name: "Alice", Age: 30, Active: true}`
- Avoid positional arguments when creating struct instances

### Imports
- Group imports: standard library, then third-party (if any)
- Always use the full module path for imports within the workspace

## Project Organization

- **`go.work`** — Workspace file that references both modules
- **`go_basics/go.mod`** — Module definition for the basics app
- **`api_project/go.mod`** — Module definition for the API project starter
- **`README.md`** — Overview of the workspace layout and how to run/build

## Development Patterns

### When Adding Code
- If it belongs to an existing concept area in `go_basics`, add it to the appropriate package (`gobasics`, `loops`, `structs`, or `library`)
- If it's new development work, add it to the `api_project` module
- Create a new package directory if the code represents a new logical unit
- Always include a `Run*` entry point function for demo code in `go_basics`

### Testing
No test files currently exist. If adding tests, follow Go conventions:
- Test files named `*_test.go` in the same package as the code they test
- Test functions named `TestFunctionName(t *testing.T)`

### Error Handling
Current code is minimal and focuses on education/demonstrations. When adding production code to `api_project`, implement proper error handling with error returns and checking.
