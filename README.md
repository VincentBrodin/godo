# GODO

Godo is a modern, Make‑inspired task runner, designed to make your build and automation workflows more intuitive and platform‑agnostic.

---

## Table of Contents
1. [Features](#features)
2. [Getting Started](#getting-started)
3. [Configuration Syntax](#configuration-syntax)
4. [Installation](#installation)
5. [Usage & Tips](#usage--tips)
6. [Contributing](#contributing)
7. [License](#license)

---

## Features

- **`where`**: Execute commands from any subdirectory, without manual `cd` steps.
- **`variants`**: Define OS‑ or environment‑specific command variants (based on `GOOS`).
- **`description`**: Provide human‑readable explanations for each task.
- **Fallback Help**: Running `godo` with no arguments lists all available tasks and their descriptions.

## Getting Started

Create a `godo.yaml` file in your project root to define your tasks. Here's a minimal example:

```yaml
commands:
  test:
    run:
      - go clean -testcache
      - go test ./... -v
    description: Runs all Go tests

  os-info:
    variants:
      - run: echo "Windows"
        platform: windows
      - run: echo "Linux"
        platform: linux
      - run: echo "Unknown OS"
        platform: default
    description: Prints the current operating system
```

### Task Configuration Options

Each task under `commands:` can include:

- **`run`** *(array of strings)*: One or more shell commands executed in sequence.  
- **`where`** *(string)*: Relative path to the directory where commands run (defaults to project root).  
- **`type`** *(string)*: Execution mode (`raw`, `shell`, or `path`). Usually, the default is sufficient.  
- **`description`** *(string)*: A short description for `godo`’s help output.  
- **`variants`** *(array)*: Platform or environment–specific overrides (ignores `run` at the root):
  - **`run`**: Command(s) to execute.  
  - **`platform`**: A `GOOS` value (e.g., `darwin`, `linux`, `windows`) or `default` fallback.  
  - **`type`**: Inherited from the task’s `type` if omitted.

> For a full list of supported `GOOS` values, see the [official Go documentation](https://go.dev/doc/install/source#environment) or this [gist of platforms](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63).

## Installation

Install the latest version via `go install`:

```bash
go install github.com/VincentBrodin/godo@latest
```

Ensure `$GOPATH/bin` (or `$GOBIN`) is in your `PATH`.

## Usage & Tips

- Run `godo` to list all tasks and descriptions.
- Execute a task: `godo <task-name>` (e.g., `godo test`).
- Combine complex workflows by breaking them into smaller, reusable `godo` tasks.

## Contributing

Contributions and feedback are welcome! To get started:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature/my-cool-task`.
3. Make your changes and add tests if applicable.
4. Submit a pull request with a clear description of your improvements.

## License

This project is licensed under the [MIT License](LICENSE).
