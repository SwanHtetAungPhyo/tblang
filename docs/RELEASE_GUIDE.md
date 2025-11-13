# Release Guide

## Quick Release Process

When you want to release a new version:

```bash
# 1. Make your changes and test
make build
make install
tblang --version

# 2. Commit and push
git add .
git commit -m "Your changes"
git push

# 3. Create release (replace 0.1.2 with your version)
./sync-homebrew.sh 0.1.2
```

That's it! Users can now update with:
```bash
brew update
brew upgrade tblang
```

## Version Numbering

Follow semantic versioning:
- **Major** (1.0.0): Breaking changes
- **Minor** (0.1.0): New features, backwards compatible
- **Patch** (0.0.1): Bug fixes

Examples:
- `0.1.1` → `0.1.2`: Bug fix
- `0.1.2` → `0.2.0`: New feature
- `0.9.0` → `1.0.0`: Major release

## What Happens During Release

The `sync-homebrew.sh` script:

1. ✅ Pushes code to GitHub
2. ✅ Creates git tag (e.g., `v0.1.2`)
3. ✅ Creates GitHub release with notes
4. ✅ Calculates SHA256 of tarball
5. ✅ Updates `tblang.rb` formula
6. ✅ Pushes to Homebrew tap

## End User Experience

### First Install
```bash
brew tap swanhtetaungphyo/tblang
brew install tblang
```

### Updates
```bash
brew update
brew upgrade tblang
```

Homebrew will automatically:
- Detect new version (0.1.0 → 0.1.1)
- Download new tarball
- Build from source
- Install new version
- Remove old version

## Troubleshooting

### If release fails:

```bash
# Delete tag locally and remotely
git tag -d v0.1.2
git push origin :refs/tags/v0.1.2

# Delete GitHub release
gh release delete v0.1.2

# Try again
./sync-homebrew.sh 0.1.2
```

### If users don't see update:

```bash
# Clear Homebrew cache
brew update-reset

# Force update
brew upgrade tblang
```

## Current Version

Check current version:
```bash
git describe --tags --abbrev=0
```

Latest: **v0.1.1**
