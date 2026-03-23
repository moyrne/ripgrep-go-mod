## ADDED Requirements

### Requirement: Command-line tool accepts path argument
The tool SHALL accept a path argument specifying the directory where `go list -json -m all` should be executed.

#### Scenario: Tool is called with a valid path
- **WHEN** the user runs the tool with a valid directory path
- **THEN** the tool changes to that directory

#### Scenario: Tool is called without a path argument
- **WHEN** the user runs the tool without a path argument
- **THEN** the tool uses the current working directory

### Requirement: Execute go list command
The tool SHALL execute `go list -json -m all` in the specified directory.

#### Scenario: Go command is available
- **WHEN** the `go` command is available in the PATH
- **THEN** the tool executes `go list -json -m all` successfully

#### Scenario: Go command is not available
- **WHEN** the `go` command is not available in the PATH
- **THEN** the tool returns an error indicating that `go` is not found

### Requirement: Parse JSON output into memory
The tool SHALL parse the JSON output from `go list -json -m all` into an in-memory data structure.

#### Scenario: Valid JSON output is received
- **WHEN** `go list -json -m all` produces valid JSON output
- **THEN** the tool parses the JSON into a structured format in memory

#### Scenario: Invalid JSON output is received
- **WHEN** `go list -json -m all` produces invalid JSON output
- **THEN** the tool returns an error indicating JSON parsing failure