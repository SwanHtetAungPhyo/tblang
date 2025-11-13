# TBLang Scripts

This directory contains utility scripts for TBLang installation, configuration, and release management.

## Available Scripts

### Installation

#### install-tblang.sh
Installs TBLang and the AWS provider plugin system-wide.

```bash
./scripts/install-tblang.sh
```

**What it does:**
- Builds TBLang core from source
- Builds AWS provider plugin
- Installs binaries to `/usr/local/bin/`
- Installs plugins to `/usr/local/lib/tblang/plugins/`
- Requires sudo privileges

**Requirements:**
- Go 1.21 or later
- Git
- Make

### Shell Completion

#### enable-completion.sh
Enables shell completions for Bash, Zsh, or Fish.

```bash
./scripts/enable-completion.sh
```

**What it does:**
- Detects your current shell
- Generates completion scripts
- Installs completions to appropriate directories
- Provides instructions for enabling completions

**Supported Shells:**
- Bash
- Zsh
- Fish

### Release Management

#### sync-homebrew.sh
Automates the release process and syncs with Homebrew tap.

```bash
./scripts/sync-homebrew.sh <version>
```

**Example:**
```bash
./scripts/sync-homebrew.sh 1.1.1
```

**What it does:**
1. Pushes changes to GitHub
2. Creates and pushes a git tag
3. Creates a GitHub release
4. Calculates SHA256 of the release tarball
5. Updates the Homebrew formula
6. Pushes to the Homebrew tap repository
7. Commits formula changes

**Requirements:**
- GitHub CLI (`gh`) installed and authenticated
- Write access to the main repository
- Write access to the Homebrew tap repository
- Clean git working directory

**Environment:**
- Expects Homebrew tap at: `/opt/homebrew/Library/Taps/swanhtetaungphyo/homebrew-tblang`
- Updates `tblang.rb` in the project root

## Usage Examples

### Fresh Installation

```bash
# Clone the repository
git clone https://github.com/SwanHtetAungPhyo/tblang.git
cd tblang

# Install TBLang
./scripts/install-tblang.sh

# Enable shell completions
./scripts/enable-completion.sh

# Verify installation
tblang --version
```

### Creating a New Release

```bash
# Make your changes
git add .
git commit -m "Your changes"

# Update version in core/cmd/tblang/main.go
# Then sync with Homebrew
./scripts/sync-homebrew.sh 1.2.0
```

### Manual Installation Steps

If you prefer manual installation:

```bash
# Build
make build

# Install
make install

# Or use the script
./scripts/install-tblang.sh
```

## Script Maintenance

### Adding New Scripts

When adding new scripts to this directory:

1. Make the script executable: `chmod +x scripts/your-script.sh`
2. Add a shebang line: `#!/bin/bash`
3. Include error handling: `set -e`
4. Add usage instructions in comments
5. Update this README with documentation

### Best Practices

- Use `set -e` to exit on errors
- Provide clear error messages
- Include usage examples
- Check for required dependencies
- Use colored output for better UX
- Test on multiple platforms

## Troubleshooting

### install-tblang.sh fails

**Issue:** Permission denied
```bash
sudo ./scripts/install-tblang.sh
```

**Issue:** Go not found
```bash
# Install Go first
brew install go
# or download from https://golang.org/dl/
```

### enable-completion.sh doesn't work

**Issue:** Completions not loading
- Restart your shell or source your shell config
- Check that the completion directory is in your shell's path

### sync-homebrew.sh fails

**Issue:** GitHub CLI not authenticated
```bash
gh auth login
```

**Issue:** Homebrew tap not found
- Clone the tap repository to the expected location
- Or update the `TAP_DIR` variable in the script

## Related Documentation

- [MAINTAINER_GUIDE.md](../docs/MAINTAINER_GUIDE.md) - Maintainer workflows
- [RELEASE_GUIDE.md](../docs/RELEASE_GUIDE.md) - Release process details
- [HOMEBREW_SETUP.md](../docs/HOMEBREW_SETUP.md) - Homebrew installation guide

## Contributing

When modifying scripts:

1. Test on macOS and Linux if possible
2. Update this README
3. Add comments explaining complex logic
4. Follow shell scripting best practices
5. Test error conditions
