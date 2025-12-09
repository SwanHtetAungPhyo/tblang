# AWS Plugin Refactoring

## Overview

The AWS plugin has been refactored to improve code organization and maintainability. Large monolithic files have been split into smaller, focused files with single responsibilities.

## Changes Made

### Provider Package (`plugin/aws/internal/provider/`)

**Before:** 2 large files (2702 lines total)
- `provider.go`: 1794 lines
- `client.go`: 908 lines

**After:** 36 focused files

## File Organization

### Core Provider Files

| File | Lines | Purpose |
|------|-------|---------|
| `provider.go` | 32 | Package documentation and entry point |
| `types.go` | 14 | AWSProvider type definition |
| `constructor.go` | 6 | Provider constructor |
| `configure.go` | 54 | Provider configuration |
| `interface.go` | 38 | Provider interface methods (Plan, Read, Import, Validate) |
| `apply.go` | 115 | ApplyResourceChange dispatcher |
| `helpers.go` | 19 | Helper functions (extractTags) |

### Schema Files

| File | Purpose |
|------|---------|
| `schema.go` | Schema entry point and dispatcher |
| `schema_resources.go` | Resource schema definitions (VPC, Subnet, SG, EC2, IGW, RT, EIP, NAT) |
| `schema_datasources.go` | Data source schema definitions (AMI, VPC, Subnet, AZs, CallerIdentity) |

### Resource Operation Files

| File | Purpose |
|------|---------|
| `resource_vpc.go` | VPC apply and destroy operations |
| `resource_subnet.go` | Subnet apply and destroy operations |
| `resource_security_group.go` | Security Group apply and destroy operations |
| `resource_ec2.go` | EC2 instance apply and destroy operations |
| `resource_internet_gateway.go` | Internet Gateway apply and destroy operations |
| `resource_route_table.go` | Route Table apply and destroy operations |
| `resource_eip.go` | Elastic IP apply and destroy operations |
| `resource_nat_gateway.go` | NAT Gateway apply and destroy operations |

### Data Source Files

| File | Purpose |
|------|---------|
| `datasource_ami.go` | AMI data source read operation |
| `datasource_vpc.go` | VPC data source read operation |
| `datasource_subnet.go` | Subnet data source read operation |
| `datasource_availability_zones.go` | Availability Zones data source read operation |
| `datasource_caller_identity.go` | Caller Identity data source read operation |

### Client Files

| File | Purpose |
|------|---------|
| `client.go` | Package documentation and entry point |
| `client_types.go` | Type definitions for AWSClient and result types |
| `client_constructor.go` | Client constructor |
| `client_helpers.go` | Helper functions (buildTags) |
| `client_vpc.go` | VPC client operations |
| `client_subnet.go` | Subnet client operations |
| `client_security_group.go` | Security Group client operations |
| `client_ec2.go` | EC2 instance client operations |
| `client_internet_gateway.go` | Internet Gateway client operations |
| `client_route_table.go` | Route Table client operations |
| `client_eip.go` | Elastic IP client operations |
| `client_nat_gateway.go` | NAT Gateway client operations |
| `client_datasources.go` | Data source client operations |

## Benefits

1. **Better Organization**: Each file has a single, clear responsibility
2. **Easier Navigation**: Developers can quickly find relevant code
3. **Improved Maintainability**: Smaller files are easier to understand and modify
4. **Better Testing**: Focused files make unit testing more straightforward
5. **Reduced Merge Conflicts**: Smaller files reduce the likelihood of conflicts
6. **Clearer Dependencies**: File organization makes dependencies more obvious
7. **Parallel Development**: Multiple developers can work on different resources without conflicts

## File Naming Convention

- `types.go` - Type definitions and constants
- `constructor.go` - Constructor functions
- `schema.go` / `schema_*.go` - Schema definitions
- `resource_<name>.go` - Resource-specific operations (apply/destroy)
- `datasource_<name>.go` - Data source-specific operations (read)
- `client_<name>.go` - AWS SDK client operations for specific resources
- `helpers.go` - Shared helper functions

## Code Structure

### Provider Layer
```
provider.go (entry point)
├── types.go (AWSProvider definition)
├── constructor.go (NewAWSProvider)
├── configure.go (Configure method)
├── interface.go (Plan, Read, Import, Validate)
├── apply.go (ApplyResourceChange dispatcher)
├── schema.go (GetSchema dispatcher)
│   ├── schema_resources.go (resource schemas)
│   └── schema_datasources.go (data source schemas)
├── resource_*.go (resource operations)
├── datasource_*.go (data source operations)
└── helpers.go (shared utilities)
```

### Client Layer
```
client.go (entry point)
├── client_types.go (AWSClient, result types)
├── client_constructor.go (NewAWSClient)
├── client_helpers.go (buildTags, etc.)
├── client_vpc.go (VPC SDK operations)
├── client_subnet.go (Subnet SDK operations)
├── client_security_group.go (SG SDK operations)
├── client_ec2.go (EC2 SDK operations)
├── client_internet_gateway.go (IGW SDK operations)
├── client_route_table.go (RT SDK operations)
├── client_eip.go (EIP SDK operations)
├── client_nat_gateway.go (NAT SDK operations)
└── client_datasources.go (Data source SDK operations)
```

## Backward Compatibility

All refactoring maintains 100% backward compatibility. No public APIs were changed, only internal organization was improved.

## Testing

All code compiles successfully:
```bash
cd plugin/aws && go build .
```

The plugin works correctly with the core:
```bash
tblang plan ./tblang-demo/ec2-example.tbl
```

Successfully tested with:
- VPC creation
- Subnet creation
- Security Group creation
- EC2 instance creation
- Internet Gateway creation
- Data sources (AMI, VPC, Subnet, AZs, CallerIdentity)
- Print and output functions
- Dependency graph building

## Comparison

### Before Refactoring
- 2 files
- 2702 total lines
- Largest file: 1794 lines
- Hard to navigate
- Difficult to find specific functionality
- High risk of merge conflicts

### After Refactoring
- 36 files
- ~2700 total lines (same functionality)
- Largest file: ~115 lines
- Easy to navigate
- Quick to find specific functionality
- Low risk of merge conflicts
- Clear separation of concerns

## Future Improvements

Consider:
- Adding unit tests for each resource file
- Adding integration tests
- Implementing resource update operations (currently only create/destroy)
- Adding more AWS resource types (RDS, S3, Lambda, etc.)
- Implementing resource import functionality
- Adding configuration validation
