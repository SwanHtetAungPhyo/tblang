# TBLang - Infrastructure as Code Language

TBLang is a modern, plugin-based Infrastructure as Code (IaC) language designed for simplicity and extensibility. Built with a clean separation between the core language engine and cloud provider plugins, TBLang offers a powerful yet intuitive way to manage infrastructure.

## Installation

```bash
brew tap swanhtetaungphyo/tblang
brew install tblang
```

Verify installation:
```bash
tblang --version
tblang plugins list
```

**Shell Completion:**
Completions are automatically installed for Bash, Zsh, and Fish. See [COMPLETION_GUIDE.md](docs/COMPLETION_GUIDE.md) for setup instructions.

## Features

- 🚀 **Plugin Architecture**: Extensible provider system using gRPC
- 🎨 **Colorful CLI**: Beautiful terminal output with fatih/color
- 📦 **State Management**: Track infrastructure with `.tbstate` files
- 🔗 **Dependency Resolution**: Automatic resource dependency handling
- 🔐 **Profile Support**: AWS profile configuration in language
- ⚡ **Fast & Efficient**: Built in Go for performance

## Quick Start

### Installation

#### Option 1: Homebrew (Recommended)

```bash
brew tap swanhtetaungphyo/tblang
brew install tblang
```

#### Option 2: Installation Script

```bash
git clone https://github.com/SwanHtetAungPhyo/tblang.git
cd tblang
./scripts/install-tblang.sh
```

#### Option 3: Manual Build

```bash
# Clone the repository
git clone https://github.com/SwanHtetAungPhyo/tblang.git
cd tblang

# Build core
cd core && go build -o tblang ./cmd/tblang
sudo cp tblang /usr/local/bin/

# Build AWS plugin
cd ../plugin/aws && go build -o tblang-provider-aws main.go
sudo mkdir -p /usr/local/lib/tblang/plugins
sudo cp tblang-provider-aws /usr/local/lib/tblang/plugins/
```

### Your First Infrastructure

Create `infrastructure.tbl`:

```tblang
cloud_vendor "aws" {
    region = "us-east-1"
    profile = "default"
}

declare vpc_config = {
    cidr_block: "10.0.0.0/16"
    enable_dns_hostnames: true
    enable_dns_support: true
    tags: {
        Name: "my-vpc"
        Environment: "production"
        ManagedBy: "TBLang"
    }
}

declare main_vpc = vpc("my-vpc", vpc_config);

declare subnet_config = {
    vpc_id: main_vpc
    cidr_block: "10.0.1.0/24"
    availability_zone: "us-east-1a"
    map_public_ip: true
    tags: {
        Name: "my-public-subnet"
        Type: "public"
    }
}

declare public_subnet = subnet("my-public-subnet", subnet_config);
```

Run commands:

```bash
tblang plan infrastructure.tbl
tblang apply infrastructure.tbl
tblang show
tblang destroy infrastructure.tbl
```

## Language Reference

### Cloud Vendor Configuration

Define your cloud provider and credentials:

```tblang
cloud_vendor "aws" {
    region = "us-east-1"
    profile = "my-aws-profile"
    account_id = "123456789012"  // Optional
}
```

**Supported Providers:**
- `aws` - Amazon Web Services

### Variable Declarations

Declare reusable configuration blocks:

```tblang
declare variable_name = {
    key: "value"
    number: 42
    boolean: true
    nested: {
        inner_key: "inner_value"
    }
    list: ["item1", "item2"]
}
```

### Resource Creation

Create infrastructure resources:

```tblang
declare resource_var = resource_type("resource-name", configuration);
```

**Example:**
```tblang
declare my_vpc = vpc("production-vpc", {
    cidr_block: "10.0.0.0/16"
    enable_dns_hostnames: true
});
```

### For Loops

Create multiple resources dynamically:

```tblang
declare subnets = [
    { name: "subnet-1", cidr: "10.0.1.0/24" },
    { name: "subnet-2", cidr: "10.0.2.0/24" }
];

for config in subnets {
    declare subnet = subnet(config.name, {
        cidr_block: config.cidr
        vpc_id: my_vpc
    });
}
```

See [LOOP_GUIDE.md](docs/LOOP_GUIDE.md) for detailed loop documentation.

### Resource References

Reference other resources using variable names:

```tblang
declare my_vpc = vpc("my-vpc", vpc_config);

declare my_subnet = subnet("my-subnet", {
    vpc_id: my_vpc  // Reference to VPC resource
    cidr_block: "10.0.1.0/24"
});
```

TBLang automatically resolves references to actual resource IDs (e.g., `vpc-xxxxx`).

### Supported Resource Types

#### VPC (Virtual Private Cloud)

```tblang
declare my_vpc = vpc("vpc-name", {
    cidr_block: "10.0.0.0/16"
    enable_dns_hostnames: true
    enable_dns_support: true
    tags: {
        Name: "my-vpc"
        Environment: "production"
    }
});
```

**Attributes:**
- `cidr_block` (required): CIDR block for VPC
- `enable_dns_hostnames` (optional): Enable DNS hostnames
- `enable_dns_support` (optional): Enable DNS support
- `tags` (optional): Resource tags

**Computed Attributes:**
- `vpc_id`: AWS VPC ID
- `state`: VPC state

#### Subnet

```tblang
declare my_subnet = subnet("subnet-name", {
    vpc_id: my_vpc
    cidr_block: "10.0.1.0/24"
    availability_zone: "us-east-1a"
    map_public_ip: true
    tags: {
        Name: "my-subnet"
        Type: "public"
    }
});
```

**Attributes:**
- `vpc_id` (required): VPC reference or ID
- `cidr_block` (required): CIDR block for subnet
- `availability_zone` (required): AWS availability zone
- `map_public_ip` (optional): Auto-assign public IP
- `tags` (optional): Resource tags

**Computed Attributes:**
- `subnet_id`: AWS Subnet ID
- `state`: Subnet state

#### Security Group

```tblang
declare my_sg = security_group("sg-name", {
    vpc_id: my_vpc
    name: "web-sg"
    description: "Security group for web servers"
    ingress_rules: [
        {
            protocol: "tcp"
            from_port: 80
            to_port: 80
            cidr_blocks: ["0.0.0.0/0"]
        },
        {
            protocol: "tcp"
            from_port: 443
            to_port: 443
            cidr_blocks: ["0.0.0.0/0"]
        }
    ]
    egress_rules: [
        {
            protocol: "-1"
            from_port: 0
            to_port: 0
            cidr_blocks: ["0.0.0.0/0"]
        }
    ]
    tags: {
        Name: "web-sg"
    }
});
```

**Attributes:**
- `vpc_id` (required): VPC reference or ID
- `name` (required): Security group name
- `description` (optional): Description
- `ingress_rules` (optional): Inbound rules
- `egress_rules` (optional): Outbound rules
- `tags` (optional): Resource tags

**Computed Attributes:**
- `group_id`: AWS Security Group ID

### Comments

```tblang
// Single-line comment

/*
 * Multi-line comment
 */
```

## CLI Commands

### Plan

Preview infrastructure changes:

```bash
tblang plan infrastructure.tbl
```

Shows:
- Resources to create (green +)
- Resources to update (yellow ~)
- Resources to delete (red -)

### Apply

Create or update infrastructure:

```bash
tblang apply infrastructure.tbl
```

Prompts for confirmation before making changes.

### Show

Display current infrastructure state:

```bash
tblang show
```

Shows all resources with their attributes and IDs.

### Destroy

Remove all infrastructure:

```bash
tblang destroy infrastructure.tbl
```

Prompts for confirmation before destroying resources.

### Graph

Visualize dependency graph:

```bash
tblang graph infrastructure.tbl
```

Shows resource dependencies and deployment order.

### Plugins

List available provider plugins:

```bash
tblang plugins list
```

### Global Flags

- `--aws-profile <profile>`: Override AWS profile
- `--no-color`: Disable colored output
- `--help`: Show help information

## Plugin Development Guide

### Overview

TBLang uses a plugin architecture where provider plugins communicate with the core engine via gRPC. Each plugin is a separate executable that implements the provider protocol.

### Plugin Architecture

```
┌─────────────────┐
│   TBLang Core   │
│    (Engine)     │
└────────┬────────┘
         │ gRPC
         │
┌────────▼────────┐
│  Plugin Server  │
│   (Provider)    │
└────────┬────────┘
         │
┌────────▼────────┐
│   Cloud SDK     │
│  (AWS/Azure)    │
└─────────────────┘
```

### Creating a New Plugin

#### Step 1: Project Setup

```bash
mkdir -p plugin/myprovider/internal/provider
cd plugin/myprovider
go mod init github.com/tblang/plugin-myprovider
```

Add dependencies:

```bash
go get github.com/tblang/core/pkg/plugin
go get google.golang.org/grpc
```

#### Step 2: Implement Provider Interface

Create `internal/provider/provider.go`:

```go
package provider

import (
    "context"
    "fmt"
    "github.com/tblang/core/pkg/plugin"
)

type MyProvider struct {
    region string
    client *MyCloudClient
}

func NewMyProvider() *MyProvider {
    return &MyProvider{}
}

func (p *MyProvider) GetSchema(ctx context.Context, req *plugin.GetSchemaRequest) (*plugin.GetSchemaResponse, error) {
    return &plugin.GetSchemaResponse{
        Provider: &plugin.Schema{
            Version: 1,
            Block: &plugin.SchemaBlock{
                Attributes: map[string]*plugin.Attribute{
                    "region": {
                        Type:        "string",
                        Description: "Cloud region",
                        Required:    true,
                    },
                },
            },
        },
        ResourceSchemas: map[string]*plugin.Schema{
            "compute_instance": {
                Version: 1,
                Block: &plugin.SchemaBlock{
                    Attributes: map[string]*plugin.Attribute{
                        "name": {
                            Type:        "string",
                            Description: "Instance name",
                            Required:    true,
                        },
                        "instance_type": {
                            Type:        "string",
                            Description: "Instance type",
                            Required:    true,
                        },
                        "instance_id": {
                            Type:        "string",
                            Description: "Instance ID",
                            Computed:    true,
                        },
                    },
                },
            },
        },
    }, nil
}

func (p *MyProvider) Configure(ctx context.Context, req *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
    config, ok := req.Config.(map[string]interface{})
    if !ok {
        return &plugin.ConfigureResponse{
            Diagnostics: []*plugin.Diagnostic{
                {
                    Severity: "error",
                    Summary:  "Invalid configuration",
                    Detail:   "Configuration must be a map",
                },
            },
        }, nil
    }

    if region, exists := config["region"]; exists {
        if regionStr, ok := region.(string); ok {
            p.region = regionStr
        }
    }

    client, err := NewMyCloudClient(p.region)
    if err != nil {
        return &plugin.ConfigureResponse{
            Diagnostics: []*plugin.Diagnostic{
                {
                    Severity: "error",
                    Summary:  "Failed to initialize client",
                    Detail:   err.Error(),
                },
            },
        }, nil
    }

    p.client = client
    return &plugin.ConfigureResponse{}, nil
}

func (p *MyProvider) ApplyResourceChange(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
    isDestroy := req.PlannedState == nil && req.PriorState != nil

    if isDestroy {
        return p.destroyResource(ctx, req)
    }

    return p.createOrUpdateResource(ctx, req)
}

func (p *MyProvider) createOrUpdateResource(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
    config, ok := req.Config.(map[string]interface{})
    if !ok {
        return &plugin.ApplyResourceChangeResponse{
            Diagnostics: []*plugin.Diagnostic{
                {
                    Severity: "error",
                    Summary:  "Invalid configuration",
                    Detail:   "Configuration must be a map",
                },
            },
        }, nil
    }

    switch req.TypeName {
    case "compute_instance":
        instance, err := p.client.CreateInstance(ctx, config)
        if err != nil {
            return &plugin.ApplyResourceChangeResponse{
                Diagnostics: []*plugin.Diagnostic{
                    {
                        Severity: "error",
                        Summary:  "Failed to create instance",
                        Detail:   err.Error(),
                    },
                },
            }, nil
        }

        newState := make(map[string]interface{})
        for k, v := range config {
            newState[k] = v
        }
        newState["instance_id"] = instance.ID

        return &plugin.ApplyResourceChangeResponse{
            NewState: newState,
        }, nil

    default:
        return &plugin.ApplyResourceChangeResponse{
            Diagnostics: []*plugin.Diagnostic{
                {
                    Severity: "error",
                    Summary:  "Unsupported resource type",
                    Detail:   fmt.Sprintf("Resource type %s is not supported", req.TypeName),
                },
            },
        }, nil
    }
}

func (p *MyProvider) destroyResource(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
    priorState, ok := req.PriorState.(map[string]interface{})
    if !ok {
        return &plugin.ApplyResourceChangeResponse{
            Diagnostics: []*plugin.Diagnostic{
                {
                    Severity: "error",
                    Summary:  "Invalid prior state",
                    Detail:   "Prior state must be a map",
                },
            },
        }, nil
    }

    switch req.TypeName {
    case "compute_instance":
        instanceID, ok := priorState["instance_id"].(string)
        if !ok || instanceID == "" {
            return &plugin.ApplyResourceChangeResponse{
                Diagnostics: []*plugin.Diagnostic{
                    {
                        Severity: "error",
                        Summary:  "Missing instance ID",
                        Detail:   "instance_id is required to destroy instance",
                    },
                },
            }, nil
        }

        if err := p.client.DeleteInstance(ctx, instanceID); err != nil {
            return &plugin.ApplyResourceChangeResponse{
                Diagnostics: []*plugin.Diagnostic{
                    {
                        Severity: "error",
                        Summary:  "Failed to delete instance",
                        Detail:   err.Error(),
                    },
                },
            }, nil
        }

        return &plugin.ApplyResourceChangeResponse{
            NewState: nil,
        }, nil

    default:
        return &plugin.ApplyResourceChangeResponse{
            Diagnostics: []*plugin.Diagnostic{
                {
                    Severity: "error",
                    Summary:  "Unsupported resource type",
                    Detail:   fmt.Sprintf("Resource type %s is not supported", req.TypeName),
                },
            },
        }, nil
    }
}

func (p *MyProvider) PlanResourceChange(ctx context.Context, req *plugin.PlanResourceChangeRequest) (*plugin.PlanResourceChangeResponse, error) {
    return &plugin.PlanResourceChangeResponse{
        PlannedState: req.ProposedNewState,
    }, nil
}

func (p *MyProvider) ReadResource(ctx context.Context, req *plugin.ReadResourceRequest) (*plugin.ReadResourceResponse, error) {
    return &plugin.ReadResourceResponse{
        NewState: req.CurrentState,
    }, nil
}

func (p *MyProvider) ImportResource(ctx context.Context, req *plugin.ImportResourceRequest) (*plugin.ImportResourceResponse, error) {
    return &plugin.ImportResourceResponse{}, nil
}

func (p *MyProvider) ValidateResourceConfig(ctx context.Context, req *plugin.ValidateResourceConfigRequest) (*plugin.ValidateResourceConfigResponse, error) {
    return &plugin.ValidateResourceConfigResponse{}, nil
}
```

#### Step 3: Create Cloud Client

Create `internal/provider/client.go`:

```go
package provider

import (
    "context"
    "fmt"
)

type MyCloudClient struct {
    region string
}

type Instance struct {
    ID   string
    Name string
}

func NewMyCloudClient(region string) (*MyCloudClient, error) {
    return &MyCloudClient{
        region: region,
    }, nil
}

func (c *MyCloudClient) CreateInstance(ctx context.Context, config map[string]interface{}) (*Instance, error) {
    name, _ := config["name"].(string)
    
    return &Instance{
        ID:   "inst-12345",
        Name: name,
    }, nil
}

func (c *MyCloudClient) DeleteInstance(ctx context.Context, instanceID string) error {
    return nil
}
```

#### Step 4: Create Plugin Main

Create `main.go`:

```go
package main

import (
    "github.com/tblang/plugin-myprovider/internal/provider"
    "github.com/tblang/core/pkg/plugin"
)

func main() {
    p := provider.NewMyProvider()
    server := plugin.NewGRPCServer(p)
    server.Serve()
}
```

#### Step 5: Build and Install

```bash
go build -o tblang-provider-myprovider main.go
sudo cp tblang-provider-myprovider /usr/local/lib/tblang/plugins/
```

#### Step 6: Test Your Plugin

Create `test.tbl`:

```tblang
cloud_vendor "myprovider" {
    region = "us-west-1"
}

declare instance_config = {
    name: "my-instance"
    instance_type: "small"
}

declare my_instance = compute_instance("my-instance", instance_config);
```

Run:

```bash
tblang plan test.tbl
tblang apply test.tbl
```

### Plugin Protocol Reference

#### Required Methods

1. **GetSchema**: Return provider and resource schemas
2. **Configure**: Initialize provider with configuration
3. **ApplyResourceChange**: Create, update, or destroy resources
4. **PlanResourceChange**: Plan resource changes
5. **ReadResource**: Read current resource state
6. **ImportResource**: Import existing resources
7. **ValidateResourceConfig**: Validate resource configuration

#### Resource Lifecycle

1. **Create**: `PriorState = nil`, `PlannedState = config`
2. **Update**: `PriorState = old`, `PlannedState = new`
3. **Destroy**: `PriorState = old`, `PlannedState = nil`

### Best Practices

1. **Error Handling**: Return diagnostics instead of errors
2. **Idempotency**: Ensure operations can be safely retried
3. **State Management**: Always return accurate resource state
4. **Validation**: Validate configuration early
5. **Logging**: Use structured logging for debugging
6. **Testing**: Write unit and integration tests

## Project Structure

```
tblang/
├── core/                           # TBLang Core
│   ├── cmd/tblang/                # CLI application
│   │   └── main.go
│   ├── internal/
│   │   ├── compiler/              # Language compiler
│   │   ├── engine/                # Core engine
│   │   ├── state/                 # State management
│   │   ├── ast/                   # AST definitions
│   │   └── graph/                 # Dependency graph
│   ├── pkg/plugin/                # Plugin protocol
│   │   ├── protocol.go
│   │   ├── grpc_server.go
│   │   ├── grpc_client.go
│   │   └── proto/
│   ├── grammar/                   # ANTLR grammar
│   └── go.mod
├── plugin/                        # Provider Plugins
│   └── aws/                       # AWS Provider
│       ├── internal/provider/
│       │   ├── provider.go
│       │   └── client.go
│       ├── main.go
│       └── go.mod
├── tblang-demo/                   # Example projects
├── install-tblang.sh              # Installation script
└── README.md
```

## Contributing

We welcome contributions! Here's how to get started:

### Development Setup

1. Fork the repository
2. Clone your fork
3. Create a feature branch
4. Make your changes
5. Run tests
6. Submit a pull request

### Code Style

- Follow Go conventions
- Use `gofmt` for formatting
- Write clear commit messages
- Add tests for new features
- Update documentation

### Testing

```bash
cd core && go test ./...
cd plugin/aws && go test ./...
```

### Submitting Changes

1. Ensure all tests pass
2. Update documentation
3. Add examples if needed
4. Create a pull request with clear description

## License

MIT License - see LICENSE file for details.

## Documentation

Comprehensive documentation is available in the `docs/` folder:

- **[FEATURES.md](docs/FEATURES.md)** - Complete feature list and capabilities
- **[LOOP_GUIDE.md](docs/LOOP_GUIDE.md)** - Loop syntax and examples
- **[LOOP_IMPLEMENTATION_SUMMARY.md](docs/LOOP_IMPLEMENTATION_SUMMARY.md)** - Technical implementation details
- **[COMPLETION_GUIDE.md](docs/COMPLETION_GUIDE.md)** - Shell completion setup
- **[HOMEBREW_SETUP.md](docs/HOMEBREW_SETUP.md)** - Homebrew installation guide
- **[MAINTAINER_GUIDE.md](docs/MAINTAINER_GUIDE.md)** - Guide for maintainers
- **[RELEASE_GUIDE.md](docs/RELEASE_GUIDE.md)** - Release process documentation

## Scripts

Utility scripts are available in the `scripts/` folder:

- **[install-tblang.sh](scripts/install-tblang.sh)** - Installation script
- **[enable-completion.sh](scripts/enable-completion.sh)** - Enable shell completions
- **[sync-homebrew.sh](scripts/sync-homebrew.sh)** - Sync releases with Homebrew

## Support

- GitHub Issues: Report bugs and request features
- Documentation: Check the `docs/` folder for detailed guides
- Examples: See `tblang-demo/` for example projects

## Roadmap

- [ ] Azure provider plugin
- [ ] GCP provider plugin
- [ ] Kubernetes provider plugin
- [ ] Module system
- [ ] Remote state backend
- [ ] Terraform import compatibility
- [ ] VS Code extension

---

Built with ❤️ using Go and gRPC
