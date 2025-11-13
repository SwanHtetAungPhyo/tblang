#!/bin/bash

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${YELLOW}Syncing TBLang with Homebrew...${NC}"

# Push to GitHub
echo -e "${YELLOW}1. Pushing to GitHub...${NC}"
git push origin main
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Pushed to GitHub${NC}"
else
    echo "Failed to push to GitHub"
    exit 1
fi

# Update Homebrew tap
echo -e "${YELLOW}2. Updating Homebrew tap...${NC}"
TAP_DIR="/opt/homebrew/Library/Taps/swanhtetaungphyo/homebrew-tblang"

if [ -d "$TAP_DIR" ]; then
    # Copy updated formula
    cp tblang.rb "$TAP_DIR/Formula/"
    
    # Commit and push tap changes
    cd "$TAP_DIR"
    git add Formula/tblang.rb
    git commit -m "Update TBLang formula" || true
    git push origin main || true
    
    echo -e "${GREEN}✓ Homebrew tap updated${NC}"
else
    echo -e "${YELLOW}⚠ Homebrew tap not found locally${NC}"
    echo "Users can update with: brew update && brew upgrade tblang"
fi

# Update Homebrew
echo -e "${YELLOW}3. Updating local Homebrew...${NC}"
brew update
brew upgrade tblang 2>/dev/null || echo "TBLang not installed locally"

echo -e "${GREEN}✓ Sync complete!${NC}"
echo ""
echo "Users can update with:"
echo "  brew update"
echo "  brew upgrade tblang"
