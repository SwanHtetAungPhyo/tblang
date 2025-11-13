#!/bin/bash

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

# Check if version argument is provided
if [ -z "$1" ]; then
    echo -e "${RED}Error: Version number required${NC}"
    echo "Usage: ./sync-homebrew.sh <version>"
    echo "Example: ./sync-homebrew.sh 0.1.2"
    exit 1
fi

VERSION=$1

echo -e "${YELLOW}Syncing TBLang v${VERSION} with Homebrew...${NC}"

# Push to GitHub
echo -e "${YELLOW}1. Pushing to GitHub...${NC}"
git push origin main
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Pushed to GitHub${NC}"
else
    echo "Failed to push to GitHub"
    exit 1
fi

# Create and push tag
echo -e "${YELLOW}2. Creating release tag v${VERSION}...${NC}"
git tag -a "v${VERSION}" -m "Release v${VERSION}"
git push origin "v${VERSION}"
echo -e "${GREEN}✓ Tag created and pushed${NC}"

# Create GitHub release
echo -e "${YELLOW}3. Creating GitHub release...${NC}"
gh release create "v${VERSION}" --title "v${VERSION}" --generate-notes
echo -e "${GREEN}✓ GitHub release created${NC}"

# Get SHA256 of tarball
echo -e "${YELLOW}4. Calculating SHA256...${NC}"
SHA256=$(curl -sL "https://github.com/SwanHtetAungPhyo/tblang/archive/refs/tags/v${VERSION}.tar.gz" | shasum -a 256 | awk '{print $1}')
echo -e "${GREEN}✓ SHA256: ${SHA256}${NC}"

# Update formula
echo -e "${YELLOW}5. Updating Homebrew formula...${NC}"
sed -i '' "s|url \".*\"|url \"https://github.com/SwanHtetAungPhyo/tblang/archive/refs/tags/v${VERSION}.tar.gz\"|" tblang.rb
sed -i '' "s|sha256 \".*\"|sha256 \"${SHA256}\"|" tblang.rb
sed -i '' "s|version \".*\"|version \"${VERSION}\"|" tblang.rb

# Update Homebrew tap
TAP_DIR="/opt/homebrew/Library/Taps/swanhtetaungphyo/homebrew-tblang"

if [ -d "$TAP_DIR" ]; then
    cp tblang.rb "$TAP_DIR/Formula/"
    
    cd "$TAP_DIR"
    git add Formula/tblang.rb
    git commit -m "Update TBLang to v${VERSION}"
    git push origin main
    
    echo -e "${GREEN}✓ Homebrew tap updated${NC}"
else
    echo -e "${YELLOW}⚠ Homebrew tap not found locally${NC}"
fi

# Commit formula changes
cd -
git add tblang.rb
git commit -m "Update formula to v${VERSION}"
git push origin main

echo -e "${GREEN}✓ Sync complete!${NC}"
echo ""
echo "Users can update with:"
echo "  brew update"
echo "  brew upgrade tblang"
