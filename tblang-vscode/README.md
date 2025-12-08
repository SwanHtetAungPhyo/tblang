# TBLang VS Code Extension

Language support for TBLang - a modern Infrastructure as Code language.

## Features

### Syntax Highlighting

Full syntax highlighting for:
- **Keywords**: `declare`, `cloud_vendor`, `for`, `in`
- **Resource Types**: `vpc`, `subnet`, `security_group`, `ec2`, `internet_gateway`, `nat_gateway`, `eip`, `route_table`
- **Data Sources**: `data_ami`, `data_vpc`, `data_subnet`, `data_availability_zones`, `data_caller_identity`
- **Built-in Functions**: `print`, `output`
- **Strings, Numbers, Booleans**
- **Comments**: Single-line (`//`) and multi-line (`/* */`)
- **Operators and Punctuation**

### Code Snippets

Quick snippets for common patterns:

| Prefix | Description |
|--------|-------------|
| `cloud` | Cloud vendor configuration |
| `vpc` | VPC resource |
| `subnet` | Subnet resource |
| `sg` | Security group |
| `ec2` | EC2 instance |
| `igw` | Internet Gateway |
| `nat` | NAT Gateway |
| `eip` | Elastic IP |
| `rt` | Route Table |
| `dataami` | Data source for AMI |
| `datavpc` | Data source for VPC |
| `dataaz` | Data source for Availability Zones |
| `for` | For loop |
| `print` | Print statement |
| `output` | Output statement |
| `vpcstack` | Complete VPC stack |

### Bracket Matching & Auto-Completion

- Automatic bracket matching for `{}`, `[]`, `()`
- Auto-closing pairs for brackets and quotes
- Smart indentation

## Installation

### From VSIX

1. Download the `.vsix` file
2. Open VS Code
3. Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on Mac)
4. Type "Install from VSIX"
5. Select the downloaded file

### From Source

```bash
cd tblang-vscode
npm install
npm run package
```

## Usage

Create a file with `.tbl` extension and start writing TBLang code:

```tblang
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
    ami: "ami-0123456789abcdef0"
    instance_type: "t2.micro"
    subnet_id: main_vpc
    tags: {
        Name: "web-server"
    }
});

print("Infrastructure defined!");
output("VPC ID", main_vpc);
```

## Supported Resource Types

### Network Resources
- `vpc` - Virtual Private Cloud
- `subnet` - Subnet
- `internet_gateway` - Internet Gateway
- `nat_gateway` - NAT Gateway
- `route_table` - Route Table
- `eip` - Elastic IP

### Compute Resources
- `ec2` - EC2 Instance
- `security_group` - Security Group

### Data Sources
- `data_ami` - Find AMI images
- `data_vpc` - Query VPC information
- `data_subnet` - Query subnet information
- `data_availability_zones` - List availability zones
- `data_caller_identity` - Get AWS caller identity

## Changelog

### 1.1.0
- Added syntax highlighting for new resource types (EC2, Internet Gateway, NAT Gateway, EIP, Route Table)
- Added syntax highlighting for data sources
- Added syntax highlighting for built-in functions (print, output)
- Added 20+ new code snippets
- Improved bracket matching and auto-indentation
- Added string interpolation highlighting

### 1.0.0
- Initial release
- Basic syntax highlighting
- Core snippets for VPC, Subnet, Security Group

## Contributing

Contributions are welcome! Please visit the [GitHub repository](https://github.com/SwanHtetAungPhyo/tblang).

## License

MIT License
