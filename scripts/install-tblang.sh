#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║     TBLang Installation Script        ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""

check_dependencies() {
    echo -e "${YELLOW}Checking dependencies...${NC}"
    
    if ! command -v go &> /dev/null; then
        echo -e "${RED}✗ Go is not installed${NC}"
        echo "  Please install Go 1.21+ from https://golang.org/dl/"
        exit 1
    fi
    
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    echo -e "${GREEN}✓ Go ${GO_VERSION} found${NC}"
    
    if ! command -v aws &> /dev/null; then
        echo -e "${YELLOW}⚠ AWS CLI not found (optional for AWS provider)${NC}"
    else
        echo -e "${GREEN}✓ AWS CLI found${NC}"
    fi
    echo ""
}

build_core() {
    echo -e "${YELLOW}Building TBLang core...${NC}"
    cd core
    go build -ldflags="-s -w" -o tblang ./cmd/tblang
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Core built successfully${NC}"
    else
        echo -e "${RED}✗ Core build failed${NC}"
        exit 1
    fi
    cd ..
    echo ""
}

build_plugins() {
    echo -e "${YELLOW}Building AWS provider plugin...${NC}"
    cd plugin/aws
    go build -o tblang-provider-aws main.go
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ AWS plugin built successfully${NC}"
    else
        echo -e "${RED}✗ AWS plugin build failed${NC}"
        exit 1
    fi
    cd ../..
    echo ""
}

install_binaries() {
    echo -e "${YELLOW}Installing binaries...${NC}"
    
    sudo cp core/tblang /usr/local/bin/
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Core installed to /usr/local/bin/tblang${NC}"
    else
        echo -e "${RED}✗ Failed to install core${NC}"
        exit 1
    fi
    
    sudo mkdir -p /usr/local/lib/tblang/plugins
    sudo cp plugin/aws/tblang-provider-aws /usr/local/lib/tblang/plugins/
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ AWS plugin installed to /usr/local/lib/tblang/plugins/${NC}"
    else
        echo -e "${RED}✗ Failed to install AWS plugin${NC}"
        exit 1
    fi
    echo ""
}

verify_installation() {
    echo -e "${YELLOW}Verifying installation...${NC}"
    
    if command -v tblang &> /dev/null; then
        VERSION=$(tblang --version 2>&1 || echo "unknown")
        echo -e "${GREEN}✓ TBLang is installed: ${VERSION}${NC}"
    else
        echo -e "${RED}✗ TBLang command not found${NC}"
        exit 1
    fi
    
    if [ -f "/usr/local/lib/tblang/plugins/tblang-provider-aws" ]; then
        echo -e "${GREEN}✓ AWS provider plugin is installed${NC}"
    else
        echo -e "${RED}✗ AWS provider plugin not found${NC}"
        exit 1
    fi
    echo ""
}

show_next_steps() {
    echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
    echo -e "${GREEN}║   Installation Complete! 🎉           ║${NC}"
    echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
    echo ""
    echo -e "${BLUE}Next steps:${NC}"
    echo ""
    echo -e "  1. Configure AWS credentials:"
    echo -e "     ${YELLOW}aws configure --profile your-profile${NC}"
    echo ""
    echo -e "  2. Create your first infrastructure file:"
    echo -e "     ${YELLOW}cat > infrastructure.tbl << 'EOF'
cloud_vendor \"aws\" {
    region = \"us-east-1\"
    profile = \"your-profile\"
}

declare vpc_config = {
    cidr_block: \"10.0.0.0/16\"
    tags: { Name: \"my-vpc\" }
}

declare my_vpc = vpc(\"my-vpc\", vpc_config);
EOF${NC}"
    echo ""
    echo -e "  3. Run TBLang commands:"
    echo -e "     ${YELLOW}tblang plan infrastructure.tbl${NC}"
    echo -e "     ${YELLOW}tblang apply infrastructure.tbl${NC}"
    echo -e "     ${YELLOW}tblang show${NC}"
    echo ""
    echo -e "  4. Get help:"
    echo -e "     ${YELLOW}tblang --help${NC}"
    echo ""
    echo -e "${BLUE}Examples available in:${NC} tblang-demo/"
    echo ""
}

main() {
    check_dependencies
    build_core
    build_plugins
    install_binaries
    verify_installation
    show_next_steps
}

main
