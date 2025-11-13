#!/bin/bash

# TBLang Shell Completion Setup Script

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}TBLang Shell Completion Setup${NC}"
echo ""

# Detect shell
SHELL_NAME=$(basename "$SHELL")

case "$SHELL_NAME" in
    bash)
        echo -e "${YELLOW}Setting up Bash completion...${NC}"
        
        # Check if bash-completion is installed
        if ! brew list bash-completion &>/dev/null; then
            echo "Installing bash-completion..."
            brew install bash-completion
        fi
        
        # Generate completion
        tblang completion bash > /usr/local/etc/bash_completion.d/tblang
        
        echo -e "${GREEN}✓ Bash completion installed${NC}"
        echo ""
        echo "Add this to your ~/.bash_profile or ~/.bashrc:"
        echo -e "${BLUE}[[ -r \"/usr/local/etc/profile.d/bash_completion.sh\" ]] && . \"/usr/local/etc/profile.d/bash_completion.sh\"${NC}"
        echo ""
        echo "Then run: source ~/.bash_profile"
        ;;
        
    zsh)
        echo -e "${YELLOW}Setting up Zsh completion...${NC}"
        
        # Create completion directory if it doesn't exist
        mkdir -p ~/.zsh/completion
        
        # Generate completion
        tblang completion zsh > ~/.zsh/completion/_tblang
        
        echo -e "${GREEN}✓ Zsh completion installed${NC}"
        echo ""
        echo "Add this to your ~/.zshrc:"
        echo -e "${BLUE}fpath=(~/.zsh/completion \$fpath)${NC}"
        echo -e "${BLUE}autoload -Uz compinit && compinit${NC}"
        echo ""
        echo "Then run: source ~/.zshrc"
        ;;
        
    fish)
        echo -e "${YELLOW}Setting up Fish completion...${NC}"
        
        # Create completion directory if it doesn't exist
        mkdir -p ~/.config/fish/completions
        
        # Generate completion
        tblang completion fish > ~/.config/fish/completions/tblang.fish
        
        echo -e "${GREEN}✓ Fish completion installed${NC}"
        echo ""
        echo "Completions are automatically loaded in Fish!"
        echo "Restart your shell or run: source ~/.config/fish/config.fish"
        ;;
        
    *)
        echo -e "${YELLOW}Unknown shell: $SHELL_NAME${NC}"
        echo ""
        echo "Generate completion manually:"
        echo "  tblang completion bash > /path/to/completion"
        echo "  tblang completion zsh > /path/to/completion"
        echo "  tblang completion fish > /path/to/completion"
        ;;
esac

echo ""
echo -e "${GREEN}Test completion by typing:${NC}"
echo "  tblang <TAB>"
