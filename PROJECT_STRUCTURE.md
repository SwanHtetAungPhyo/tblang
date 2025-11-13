# TBLang Project Structure

## Overview

TBLang is organized into clear, logical directories for easy navigation and maintenance.

## Directory Structure

```
tblang/
├── core/                       # TBLang Core Engine
│   ├── cmd/tblang/            # CLI application entry point
│   ├── internal/              # Internal packages
│   │   ├── compiler/          # Language compiler & parser
│   │   ├── engine/            # Core execution engine
│   │   ├── state/             # State management
│   │   ├── ast/               # Abstract Syntax Tree
│   │   └── graph/             # Dependency graph
│   ├── pkg/plugin/            # Plugin protocol (public API)
│   ├── grammar/               # ANTLR grammar files
│   ├── parser/                # Generated parser code
│   └── go.mod
│
├── plugin/                    # Provider Plugins
│   └── aws/                   # AWS Provider Plugin
│       ├── internal/provider/ # Provider implementation
│       ├── main.go           # Plugin entry point
│       └── go.mod
│
├── docs/                      # Documentation
│   ├── README.md             # Documentation index
│   ├── FEATURES.md           # Feature list
│   ├── LOOP_GUIDE.md         # Loop syntax guide
│   ├── LOOP_IMPLEMENTATION_SUMMARY.md  # Technical details
│   ├── COMPLETION_GUIDE.md   # Shell completion setup
│   ├── HOMEBREW_SETUP.md     # Homebrew installation
│   ├── MAINTAINER_GUIDE.md   # Maintainer workflows
│   └── RELEASE_GUIDE.md      # Release process
│
├── scripts/                   # Utility Scripts
│   ├── README.md             # Script documentation
│   ├── install-tblang.sh     # Installation script
│   ├── enable-completion.sh  # Completion setup
│   └── sync-homebrew.sh      # Release automation
│
├── tblang-demo/              # Example Projects
│   ├── infrastructure.tbl    # Basic example
│   ├── aws-loop-test.tbl     # Loop with AWS
│   ├── loop-resources-test.tbl  # Multiple resources
│   └── ...
│
├── tblang-demo-two/          # Additional Examples
│
├── .tblang/                  # Local State (gitignored)
│   ├── terraform.tbstate     # State file
│   └── plugins/              # Local plugins
│
├── README.md                 # Main project README
├── LICENSE                   # MIT License
├── Makefile                  # Build automation
├── tblang.rb                 # Homebrew formula
└── .gitignore               # Git ignore rules
```

## Key Directories

### `/core` - TBLang Core

The heart of TBLang containing:
- **CLI Application**: User-facing command-line interface
- **Compiler**: Parses `.tbl` files into AST
- **Engine**: Executes infrastructure changes
- **Plugin System**: gRPC-based plugin protocol
- **State Management**: Tracks infrastructure state

**Language**: Go
**Entry Point**: `core/cmd/tblang/main.go`

### `/plugin` - Provider Plugins

Cloud provider implementations:
- **AWS**: Amazon Web Services provider
- Future: Azure, GCP, Kubernetes, etc.

Each plugin is a standalone gRPC server that implements the provider protocol.

**Language**: Go
**Protocol**: gRPC (defined in `core/pkg/plugin/`)

### `/docs` - Documentation

All project documentation:
- User guides
- Technical documentation
- Installation guides
- Developer documentation

**Format**: Markdown
**Purpose**: Comprehensive project documentation

### `/scripts` - Utility Scripts

Automation and setup scripts:
- Installation
- Shell completion
- Release management

**Language**: Bash
**Purpose**: Developer and user utilities

### `/tblang-demo` - Examples

Real-world example projects demonstrating:
- Basic infrastructure
- Loop usage
- AWS integration
- Best practices

**Format**: `.tbl` files
**Purpose**: Learning and testing

## File Purposes

### Root Files

- **README.md**: Project overview, quick start, and main documentation
- **LICENSE**: MIT License
- **Makefile**: Build, install, test, and release automation
- **tblang.rb**: Homebrew formula for package distribution
- **.gitignore**: Git ignore patterns

### Configuration Files

- **go.mod**: Go module dependencies (in core/ and plugin/aws/)
- **.tbstate**: Infrastructure state (generated, gitignored)

### Generated Files

- **core/parser/**: Generated from ANTLR grammar
- **core/tblang**: Compiled binary (gitignored)
- **plugin/aws/tblang-provider-aws**: Compiled plugin (gitignored)

## Build Artifacts

Build artifacts are gitignored:
- `core/tblang` - Core binary
- `plugin/aws/tblang-provider-aws` - AWS plugin binary
- `dist/` - Release binaries
- `.tblang/` - Local state and plugins

## Installation Locations

When installed system-wide:
- **Binary**: `/usr/local/bin/tblang`
- **Plugins**: `/usr/local/lib/tblang/plugins/`
- **Homebrew**: `/opt/homebrew/Cellar/tblang/`

## Development Workflow

1. **Edit Code**: Modify files in `core/` or `plugin/`
2. **Build**: `make build`
3. **Test**: `make test`
4. **Install**: `make install`
5. **Test Locally**: Use examples in `tblang-demo/`
6. **Document**: Update files in `docs/`
7. **Release**: Use `scripts/sync-homebrew.sh`

## Navigation Tips

- **Start Here**: [README.md](README.md)
- **Learn Features**: [docs/FEATURES.md](docs/FEATURES.md)
- **See Examples**: [tblang-demo/](tblang-demo/)
- **Build & Install**: [Makefile](Makefile) or [scripts/install-tblang.sh](scripts/install-tblang.sh)
- **Contribute**: [docs/MAINTAINER_GUIDE.md](docs/MAINTAINER_GUIDE.md)

## Related Documentation

- [README.md](README.md) - Main project documentation
- [docs/README.md](docs/README.md) - Documentation index
- [scripts/README.md](scripts/README.md) - Script usage guide
- [docs/MAINTAINER_GUIDE.md](docs/MAINTAINER_GUIDE.md) - Maintainer workflows

## Keeping It Organized

### Adding New Files

- **Documentation**: Add to `docs/` and update `docs/README.md`
- **Scripts**: Add to `scripts/` and update `scripts/README.md`
- **Examples**: Add to `tblang-demo/` with descriptive names
- **Code**: Follow existing package structure in `core/` or `plugin/`

### Naming Conventions

- **Markdown**: `UPPERCASE_WITH_UNDERSCORES.md`
- **Scripts**: `lowercase-with-hyphens.sh`
- **Go Files**: `lowercase_with_underscores.go`
- **Directories**: `lowercase` or `kebab-case`

### Best Practices

1. Keep root directory clean (only essential files)
2. Group related files in appropriate directories
3. Update README files when adding new content
4. Use descriptive names for files and directories
5. Document new features in `docs/`
6. Add examples for new features in `tblang-demo/`

---

Last Updated: November 14, 2024
Version: 1.1.1
