# Homebrew Setup for TBLang

## For End Users

### Install via Homebrew (once published)

```bash
brew tap yourusername/tblang
brew install tblang
```

### Verify Installation

```bash
tblang --version
tblang plugins list
```

## For Maintainers

### 1. Create Homebrew Tap Repository

```bash
# Create a new repository named homebrew-tblang
# GitHub URL: https://github.com/yourusername/homebrew-tblang
```

### 2. Add Formula

Copy `tblang.rb` to the tap repository:

```bash
git clone https://github.com/yourusername/homebrew-tblang
cp tblang.rb homebrew-tblang/Formula/
cd homebrew-tblang
git add Formula/tblang.rb
git commit -m "Add TBLang formula"
git push
```

### 3. Create Release

```bash
# Tag and create release
git tag v0.1.0
git push origin v0.1.0

# Build release
make release

# Upload dist/*.tar.gz to GitHub releases
```

### 4. Update Formula SHA256

```bash
# Calculate SHA256
shasum -a 256 dist/tblang-0.1.0-darwin-amd64.tar.gz

# Update tblang.rb with the SHA256
```

### 5. Test Formula

```bash
brew install --build-from-source tblang
brew test tblang
brew audit --strict tblang
```

## Distribution Methods

### Method 1: Homebrew Tap (Recommended)

Users install via:
```bash
brew tap yourusername/tblang
brew install tblang
```

### Method 2: Direct Installation Script

Users install via:
```bash
curl -fsSL https://raw.githubusercontent.com/yourusername/tblang/main/install-tblang.sh | bash
```

### Method 3: Manual Installation

Users download and install manually:
```bash
# Download release
wget https://github.com/yourusername/tblang/releases/download/v0.1.0/tblang-0.1.0-darwin-amd64.tar.gz

# Extract
tar -xzf tblang-0.1.0-darwin-amd64.tar.gz

# Install
sudo cp tblang-darwin-amd64 /usr/local/bin/tblang
sudo mkdir -p /usr/local/lib/tblang/plugins
sudo cp tblang-provider-aws-darwin-amd64 /usr/local/lib/tblang/plugins/tblang-provider-aws
```
