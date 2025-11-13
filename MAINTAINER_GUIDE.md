# Maintainer Guide

## Updating TBLang

### 1. Make Your Changes

Edit code in `core/` or `plugin/` directories.

### 2. Test Locally

```bash
make build
make install
tblang --version
```

### 3. Push to GitHub

```bash
git add .
git commit -m "Your commit message"
git push
```

### 4. Update Homebrew Tap

```bash
./sync-homebrew.sh
```

Or manually:

```bash
# Copy formula to tap
cp tblang.rb /opt/homebrew/Library/Taps/swanhtetaungphyo/homebrew-tblang/Formula/

# Push tap changes
cd /opt/homebrew/Library/Taps/swanhtetaungphyo/homebrew-tblang
git add Formula/tblang.rb
git commit -m "Update TBLang formula"
git push
```

## End User Installation

Users install with:

```bash
brew tap swanhtetaungphyo/tblang
brew install tblang
```

## End User Updates

Users update with:

```bash
brew update
brew upgrade tblang
```

## Repositories

- **Main**: https://github.com/SwanHtetAungPhyo/tblang
- **Homebrew Tap**: https://github.com/SwanHtetAungPhyo/homebrew-tblang

## Quick Commands

```bash
# Build
make build

# Install locally
make install

# Uninstall
make uninstall

# Clean
make clean

# Test
make test

# Sync with Homebrew
./sync-homebrew.sh
```
