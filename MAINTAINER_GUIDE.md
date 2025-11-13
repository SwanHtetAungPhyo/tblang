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

### 4. Create a New Release

```bash
# Sync with Homebrew (creates tag, release, and updates formula)
./sync-homebrew.sh 0.1.2  # Replace with your version number
```

This script will:
- Push your changes to GitHub
- Create a git tag (e.g., v0.1.2)
- Create a GitHub release
- Calculate SHA256 of the tarball
- Update the Homebrew formula
- Push to the Homebrew tap

**Manual process (if needed):**

```bash
# 1. Create and push tag
git tag -a v0.1.2 -m "Release v0.1.2"
git push origin v0.1.2

# 2. Create GitHub release
gh release create v0.1.2 --title "v0.1.2" --generate-notes

# 3. Get SHA256
curl -sL https://github.com/SwanHtetAungPhyo/tblang/archive/refs/tags/v0.1.2.tar.gz | shasum -a 256

# 4. Update tblang.rb with new version, URL, and SHA256

# 5. Copy to tap and push
cp tblang.rb /opt/homebrew/Library/Taps/swanhtetaungphyo/homebrew-tblang/Formula/
cd /opt/homebrew/Library/Taps/swanhtetaungphyo/homebrew-tblang
git add Formula/tblang.rb
git commit -m "Update TBLang to v0.1.2"
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
