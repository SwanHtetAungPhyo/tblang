# Shell Completion Guide

TBLang supports shell completion for Bash, Zsh, Fish, and PowerShell.

## Automatic Installation (Homebrew)

If you installed via Homebrew, completions are automatically installed:

```bash
brew install tblang
```

Completions are installed to:
- **Bash**: `/opt/homebrew/etc/bash_completion.d/tblang`
- **Zsh**: `/opt/homebrew/share/zsh/site-functions/_tblang`
- **Fish**: `/opt/homebrew/share/fish/vendor_completions.d/tblang.fish`

## Manual Setup

### Bash

```bash
# Generate completion script
tblang completion bash > /usr/local/etc/bash_completion.d/tblang

# Or add to your ~/.bashrc
echo 'source <(tblang completion bash)' >> ~/.bashrc
source ~/.bashrc
```

### Zsh

```bash
# Generate completion script
tblang completion zsh > "${fpath[1]}/_tblang"

# Or add to your ~/.zshrc
echo 'source <(tblang completion zsh)' >> ~/.zshrc
source ~/.zshrc
```

### Fish

```bash
# Generate completion script
tblang completion fish > ~/.config/fish/completions/tblang.fish

# Reload completions
source ~/.config/fish/completions/tblang.fish
```

### PowerShell

```powershell
# Add to your PowerShell profile
tblang completion powershell | Out-String | Invoke-Expression

# Or save to profile
tblang completion powershell >> $PROFILE
```

## Testing Completions

After setup, test by typing:

```bash
tblang <TAB>
```

You should see:
```
apply       -- Apply infrastructure changes
completion  -- Generate the autocompletion script
destroy     -- Destroy infrastructure
graph       -- Show dependency graph
help        -- Help about any command
plan        -- Show what infrastructure changes will be made
plugins     -- Plugin management commands
show        -- Show current infrastructure state
```

## Completion Features

- **Commands**: All TBLang commands
- **Subcommands**: Plugin subcommands
- **Flags**: Global and command-specific flags
- **File paths**: Automatic `.tbl` file completion

## Troubleshooting

### Bash: Completions not working

```bash
# Check if bash-completion is installed
brew install bash-completion@2

# Add to ~/.bash_profile
export BASH_COMPLETION_COMPAT_DIR="/opt/homebrew/etc/bash_completion.d"
[[ -r "/opt/homebrew/etc/profile.d/bash_completion.sh" ]] && . "/opt/homebrew/etc/profile.d/bash_completion.sh"
```

### Zsh: Completions not working

```bash
# Check fpath
echo $fpath

# Ensure Homebrew completions are in fpath (add to ~/.zshrc)
if type brew &>/dev/null; then
  FPATH=$(brew --prefix)/share/zsh/site-functions:$FPATH
  autoload -Uz compinit
  compinit
fi
```

### Fish: Completions not working

```bash
# Reload Fish completions
fish_update_completions
```

## Uninstall Completions

### Bash
```bash
rm /usr/local/etc/bash_completion.d/tblang
```

### Zsh
```bash
rm "${fpath[1]}/_tblang"
```

### Fish
```bash
rm ~/.config/fish/completions/tblang.fish
```
