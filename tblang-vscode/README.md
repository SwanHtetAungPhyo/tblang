# TBLang for Visual Studio Code

Syntax highlighting and code snippets for TBLang - Infrastructure as Code Language.

## Features

- **Syntax Highlighting**: Full syntax highlighting for `.tbl` files
- **Code Snippets**: Quick snippets for common TBLang patterns
- **Auto-completion**: Bracket and quote auto-closing
- **Comment Support**: Line and block comments

## Installation

### From VSIX (Local)

1. Download the `.vsix` file
2. Open VS Code
3. Go to Extensions (Cmd+Shift+X)
4. Click "..." menu → "Install from VSIX..."
5. Select the downloaded file

### From Source

```bash
# Clone the repository
git clone https://github.com/SwanHtetAungPhyo/tblang.git
cd tblang/tblang-vscode

# Install vsce
npm install -g @vscode/vsce

# Package the extension
vsce package

# Install the generated .vsix file
code --install-extension tblang-1.0.0.vsix
```

## Snippets

Type these prefixes and press Tab:

- `cloud` - Cloud vendor configuration
- `declare` - Variable declaration
- `vpc` - VPC resource
- `subnet` - Subnet resource
- `for` - For loop

## Example

```tblang
// Cloud configuration
cloud_vendor "aws" {
    region = "us-east-1"
    profile = "default"
}

// Create VPC
declare main_vpc = vpc("my-vpc", {
    cidr_block: "10.0.0.0/16"
    enable_dns_hostnames: true
});

// Create subnets using loop
declare subnet_configs = [
    { name: "subnet-1", cidr: "10.0.1.0/24", az: "us-east-1a" },
    { name: "subnet-2", cidr: "10.0.2.0/24", az: "us-east-1b" }
];

for config in subnet_configs {
    declare sub = subnet(config.name, {
        cidr_block: config.cidr,
        vpc_id: main_vpc,
        availability_zone: config.az
    });
}
```

## Supported Syntax

- Keywords: `declare`, `cloud_vendor`, `for`, `in`
- Resource types: `vpc`, `subnet`, `security_group`, `ec2`, `instance`
- Comments: `//` and `/* */`
- Strings: `"double"` and `'single'`
- Numbers: `42`, `3.14`, `-10`
- Booleans: `true`, `false`
- Operators: `=`, `==`, `!=`, `<`, `>`, etc.

## Links

- [TBLang Repository](https://github.com/SwanHtetAungPhyo/tblang)
- [Documentation](https://github.com/SwanHtetAungPhyo/tblang/tree/main/docs)
- [Report Issues](https://github.com/SwanHtetAungPhyo/tblang/issues)

## License

MIT
