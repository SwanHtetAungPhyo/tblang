# TBLang Features

## 🚀 Core Features

### Plugin Architecture
- **Extensible**: Add new cloud providers via plugins
- **gRPC Communication**: Fast, efficient plugin protocol
- **Hot-swappable**: Plugins can be updated independently

### Language Features
- **Simple Syntax**: Easy-to-read `.tbl` files
- **Resource References**: Automatic dependency resolution
- **Type Safety**: Strong typing for resource attributes
- **Comments**: Single-line (`//`) and multi-line (`/* */`)

### State Management
- **Local State**: `.tbstate` files track infrastructure
- **Atomic Operations**: Safe concurrent operations
- **State Locking**: Prevent conflicting changes

### Dependency Management
- **Automatic Resolution**: Detects resource dependencies
- **Visual Graph**: See deployment order
- **Parallel Execution**: Independent resources in parallel

## 🎨 User Experience

### Colorful CLI
- **Green**: Success messages
- **Red**: Errors
- **Yellow**: Warnings
- **Cyan**: Information
- **Magenta**: Headers
- **Resource-specific**: Different colors for different types

### Shell Completion
- **Bash**: Full command completion
- **Zsh**: Full command completion
- **Fish**: Full command completion
- **PowerShell**: Full command completion

### Commands
- `tblang plan` - Preview changes
- `tblang apply` - Create/update infrastructure
- `tblang destroy` - Remove infrastructure
- `tblang show` - Display current state
- `tblang graph` - Visualize dependencies
- `tblang plugins list` - List available plugins

## ☁️ Cloud Providers

### AWS (Available)
- **VPC**: Virtual Private Cloud
- **Subnet**: Public and private subnets
- **Security Group**: Firewall rules
- **Profile Support**: AWS profile configuration

### Coming Soon
- Azure
- Google Cloud Platform
- Kubernetes

## 📦 Installation

### Homebrew (Recommended)
```bash
brew tap swanhtetaungphyo/tblang
brew install tblang
```

### Manual
```bash
git clone https://github.com/SwanHtetAungPhyo/tblang.git
cd tblang
./install-tblang.sh
```

## 🔄 Updates

### Automatic Updates
```bash
brew update
brew upgrade tblang
```

### Version Management
- Semantic versioning (0.1.0, 0.2.0, 1.0.0)
- GitHub releases with changelogs
- Automatic Homebrew formula updates

## 🛠️ Development

### Built With
- **Go 1.21+**: Core language
- **Cobra**: CLI framework
- **gRPC**: Plugin communication
- **AWS SDK v2**: Cloud operations
- **fatih/color**: Terminal colors

### Architecture
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

## 📚 Documentation

- [README.md](README.md) - Getting started
- [COMPLETION_GUIDE.md](COMPLETION_GUIDE.md) - Shell completion
- [MAINTAINER_GUIDE.md](MAINTAINER_GUIDE.md) - For maintainers
- [RELEASE_GUIDE.md](RELEASE_GUIDE.md) - Release process
- [HOMEBREW_SETUP.md](HOMEBREW_SETUP.md) - Homebrew setup

## 👥 Credits

Developed with ❤️ by:
- 🚀 Swan Htet Aung Phyo
- 🚀 Aung Zayar Moe

## 📄 License

MIT License - see LICENSE file for details

## 🌟 Highlights

- ✅ **Simple**: Easy-to-learn syntax
- ✅ **Fast**: Built in Go for performance
- ✅ **Extensible**: Plugin architecture
- ✅ **Beautiful**: Colorful CLI output
- ✅ **Safe**: State management and locking
- ✅ **Modern**: Latest cloud SDKs
- ✅ **Open Source**: MIT licensed
