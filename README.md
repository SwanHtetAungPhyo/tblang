# TBLang



Modern Infrastructure as Code language with plugin architecture.


## Full Documentation 

http://tblang.dev.s3-website-us-east-1.amazonaws.com/

## TBLang Architecture

TBLang follows a modular compiler architecture designed for extensibility and performance:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   .tbl Files    │───▶│   TBLang Core   │───▶│   Providers     │
│  (Source Code)  │    │   (Compiler)    │    │   (Plugins)     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │                        │
                              ▼                        ▼
                       ┌─────────────────┐    ┌─────────────────┐
                       │  .tbstate Files │    │ Cloud Resources │
                       │ (State Tracking)│    │  (AWS, GCP,     │
                       │                 │    │   Azure, etc.)  │
                       └─────────────────┘    └─────────────────┘
```

### Core Components

- **Lexer & Parser**: Processes `.tbl` syntax into Abstract Syntax Tree (AST)
- **Semantic Analyzer**: Validates resource dependencies and type checking
- **Code Generator**: Transforms AST into provider-specific API calls
- **State Manager**: Handles `.tbstate` files for infrastructure tracking
- **Plugin System**: gRPC-based provider architecture for cloud vendors

## Quick Start

### Installation

```bash
# Download and install TBLang
curl -fsSL https://install.tblang.dev | sh

# Verify installation
tblang version
```

### Your First Infrastructure

Create a `main.tbl` file:

```tbl
cloud_vendor "aws" {
    region = "us-east-1"
    profile = "default"
}

declare main_vpc = vpc("my-vpc", {
    cidr_block: "10.0.0.0/16"
    enable_dns_hostnames: true
    tags: {
        Name: "my-vpc"
        Environment: "production"
    }
});

declare web_server = ec2("web-server", {
    ami: "ami-0c55b159cbfafe1f0"
    instance_type: "t2.micro"
    subnet_id: main_vpc
    tags: {
        Name: "web-server"
    }
});
```

Deploy your infrastructure:

```bash
# Initialize TBLang project
tblang init

# Plan infrastructure changes
tblang plan

# Apply changes
tblang apply
```

## Learn More

- **[Language Reference](docs/language-reference.md)** - Complete syntax and semantics
- **[Compiler Architecture](docs/compiler-architecture.md)** - Internal design and implementation
- **[Plugin Development](docs/plugin-development.md)** - Creating custom providers
- **[State Management](docs/state-management.md)** - Understanding `.tbstate` files

---

## Features

### 🚀 Plugin Architecture

Extensible provider system using gRPC for seamless integration with cloud vendors and services.

**Key Benefits:**
- Language-agnostic plugin development
- Hot-swappable providers
- Standardized API contracts
- Built-in versioning and compatibility

### 🎨 Simple Syntax

Easy-to-read `.tbl` files with declarative infrastructure definitions.

**Language Features:**
- Type-safe resource declarations
- Built-in dependency resolution
- Intuitive block-based syntax
- Rich data types (strings, numbers, maps, lists)

### ⚡ Fast & Efficient

Built in Go for performance with optimized compilation and execution.

**Performance Highlights:**
- Sub-second compilation for large projects
- Parallel resource provisioning
- Efficient state diffing algorithms
- Memory-optimized AST processing

### 📦 State Management

Track infrastructure with `.tbstate` files for reliable change management.

**State Features:**
- Atomic state updates
- Conflict detection and resolution
- Remote state backends
- State locking mechanisms

### 🔗 Dependency Resolution

Automatic resource dependency handling with topological sorting.

**Dependency Features:**
- Implicit dependency detection
- Explicit dependency declarations
- Circular dependency prevention
- Parallel execution optimization

---

## Example

### Basic Infrastructure

```tbl
cloud_vendor "aws" {
    region = "us-east-1"
    profile = "default"
}

declare main_vpc = vpc("my-vpc", {
    cidr_block: "10.0.0.0/16"
    enable_dns_hostnames: true
    tags: {
        Name: "my-vpc"
        Environment: "production"
    }
});

declare web_server = ec2("web-server", {
    ami: "ami-0c55b159cbfafe1f0"
    instance_type: "t2.micro"
    subnet_id: main_vpc
    tags: {
        Name: "web-server"
    }
});
```

### Advanced Features

```tbl
# Variables and expressions
variable "environment" {
    type = string
    default = "development"
}

# Conditional resources
declare database = if (var.environment == "production") {
    rds("prod-db", {
        engine: "postgres"
        instance_class: "db.t3.medium"
        allocated_storage: 100
    })
} else {
    rds("dev-db", {
        engine: "postgres"
        instance_class: "db.t3.micro"
        allocated_storage: 20
    })
};

# Loops and collections
declare web_servers = for i in range(3) {
    ec2("web-server-${i}", {
        ami: "ami-0c55b159cbfafe1f0"
        instance_type: "t2.micro"
        subnet_id: main_vpc
        tags: {
            Name: "web-server-${i}"
            Index: i
        }
    })
};
```

---

## Compiler Architecture

### Compilation Pipeline

1. **Lexical Analysis**: Tokenizes `.tbl` source files
2. **Syntax Analysis**: Builds Abstract Syntax Tree (AST)
3. **Semantic Analysis**: Type checking and dependency validation
4. **Optimization**: AST transformations and dead code elimination
5. **Code Generation**: Provider-specific API call generation
6. **State Management**: Diff calculation and state updates

### Plugin System

TBLang uses a gRPC-based plugin architecture for cloud providers:

```go
// Provider interface
type Provider interface {
    Configure(config map[string]interface{}) error
    Plan(resources []Resource) (*Plan, error)
    Apply(plan *Plan) (*State, error)
    Destroy(resources []Resource) error
}
```

### Error Handling

Comprehensive error reporting with:
- Source location tracking
- Contextual error messages
- Suggested fixes
- Integration with IDEs and editors

---

## Get Started

**[Install TBLang →](docs/installation.md)**

**[Quick Start Tutorial →](docs/quickstart.md)**

**[View Examples →](examples/)**

**[API Reference →](docs/api-reference.md)**

**[Contributing →](CONTRIBUTING.md)**

---

## Community

- **GitHub**: [github.com/tblang/tblang](https://github.com/tblang/tblang)
- **Discord**: [discord.gg/tblang](https://discord.gg/tblang)
- **Documentation**: [docs.tblang.dev](https://docs.tblang.dev)
- **Blog**: [blog.tblang.dev](https://blog.tblang.dev)

## License

TBLang is open source software licensed under the [Apache License 2.0](LICENSE).