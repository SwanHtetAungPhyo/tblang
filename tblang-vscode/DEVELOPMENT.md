# TBLang VS Code Extension Development

## Quick Start

### Testing Locally

1. **Open in VS Code**:
```bash
cd tblang-vscode
code .
```

2. **Press F5** to launch Extension Development Host
   - A new VS Code window will open with the extension loaded
   - Open any `.tbl` file to see syntax highlighting

3. **Test Features**:
   - Create a new file: `test.tbl`
   - Type `cloud` and press Tab (snippet)
   - Type `for` and press Tab (loop snippet)
   - Add comments with `//` or `/* */`
   - Test auto-closing brackets `{`, `[`, `(`

### Building the Extension

```bash
# Install vsce (VS Code Extension Manager)
npm install -g @vscode/vsce

# Package the extension
vsce package

# This creates: tblang-1.0.0.vsix
```

### Installing Locally

```bash
# Install the packaged extension
code --install-extension tblang-1.0.0.vsix

# Or in VS Code:
# 1. Open Extensions (Cmd+Shift+X)
# 2. Click "..." menu
# 3. Select "Install from VSIX..."
# 4. Choose the .vsix file
```

### Testing the Extension

Create a test file `test.tbl`:

```tblang
// This is a comment
cloud_vendor "aws" {
    region = "us-east-1"
    profile = "default"
}

declare numbers = [1, 2, 3];

for num in numbers {
    declare vpc = vpc("test-vpc", {
        cidr_block: "10.0.0.0/16"
    });
}
```

**Check**:
- ✅ Keywords are highlighted (`declare`, `cloud_vendor`, `for`, `in`)
- ✅ Strings are colored (`"aws"`, `"us-east-1"`)
- ✅ Numbers are highlighted (`1`, `2`, `3`)
- ✅ Comments are grayed out
- ✅ Function names are highlighted (`vpc`)
- ✅ Brackets auto-close

## File Structure

```
tblang-vscode/
├── package.json                    # Extension manifest
├── language-configuration.json     # Language config (brackets, comments)
├── syntaxes/
│   └── tblang.tmLanguage.json     # TextMate grammar
├── snippets/
│   └── tblang.json                # Code snippets
├── README.md                       # User documentation
└── DEVELOPMENT.md                  # This file
```

## Modifying the Grammar

### Adding New Keywords

Edit `syntaxes/tblang.tmLanguage.json`:

```json
{
  "keywords": {
    "patterns": [
      {
        "name": "keyword.control.tblang",
        "match": "\\b(for|in|while|if|else|YOUR_NEW_KEYWORD)\\b"
      }
    ]
  }
}
```

### Adding New Resource Types

```json
{
  "resource-types": {
    "patterns": [
      {
        "name": "support.type.tblang",
        "match": "\\b(vpc|subnet|YOUR_NEW_RESOURCE)\\b"
      }
    ]
  }
}
```

### Adding New Snippets

Edit `snippets/tblang.json`:

```json
{
  "Your Snippet Name": {
    "prefix": "trigger",
    "body": [
      "line 1 with ${1:placeholder}",
      "line 2 with ${2:another}"
    ],
    "description": "Description shown in autocomplete"
  }
}
```

## Testing Changes

After modifying files:

1. **Reload Extension**: Press `Cmd+R` in Extension Development Host
2. **Or**: Close and press F5 again
3. **Test**: Open a `.tbl` file and verify changes

## Publishing

### Prerequisites

1. Create a [Visual Studio Marketplace](https://marketplace.visualstudio.com/) account
2. Create a [Personal Access Token](https://code.visualstudio.com/api/working-with-extensions/publishing-extension#get-a-personal-access-token)
3. Create a publisher:
```bash
vsce create-publisher your-publisher-name
```

### Publish Steps

```bash
# Login
vsce login your-publisher-name

# Publish
vsce publish

# Or publish with version bump
vsce publish patch  # 1.0.0 -> 1.0.1
vsce publish minor  # 1.0.0 -> 1.1.0
vsce publish major  # 1.0.0 -> 2.0.0
```

### Update Extension

```bash
# Make changes
# Update version in package.json
# Then publish
vsce publish
```

## Debugging

### Extension Not Loading

1. Check `package.json` syntax
2. Verify file paths in `contributes` section
3. Check VS Code Developer Tools: `Help > Toggle Developer Tools`

### Syntax Highlighting Not Working

1. Verify `.tmLanguage.json` syntax
2. Check regex patterns (escape backslashes: `\\b`)
3. Test patterns at [regex101.com](https://regex101.com/)
4. Check scope names match theme

### Snippets Not Appearing

1. Verify `snippets/tblang.json` syntax
2. Check `package.json` contributes.snippets path
3. Restart VS Code
4. Type prefix and press `Ctrl+Space`

## Useful Commands

```bash
# Package extension
vsce package

# List files that will be packaged
vsce ls

# Publish extension
vsce publish

# Unpublish extension
vsce unpublish

# Show extension info
vsce show your-publisher.tblang
```

## Resources

- [VS Code Extension API](https://code.visualstudio.com/api)
- [Language Extensions Guide](https://code.visualstudio.com/api/language-extensions/overview)
- [TextMate Grammar](https://macromates.com/manual/en/language_grammars)
- [Snippet Guide](https://code.visualstudio.com/docs/editor/userdefinedsnippets)
- [Publishing Extensions](https://code.visualstudio.com/api/working-with-extensions/publishing-extension)

## Tips

1. **Use Scope Inspector**: `Cmd+Shift+P` → "Developer: Inspect Editor Tokens and Scopes"
2. **Test with Different Themes**: Some themes may not support all scopes
3. **Keep Grammar Simple**: Start simple, add complexity gradually
4. **Test Edge Cases**: Empty files, large files, nested structures
5. **Get Feedback**: Share with users and iterate

## Next Steps

1. ✅ Basic syntax highlighting
2. ✅ Code snippets
3. ✅ Auto-closing pairs
4. 🔲 Semantic highlighting (requires LSP)
5. 🔲 Go-to-definition (requires LSP)
6. 🔲 Auto-completion (requires LSP)
7. 🔲 Error diagnostics (requires LSP)
8. 🔲 Formatting (requires formatter)

For advanced features, consider implementing a [Language Server](../docs/SYNTAX_HIGHLIGHTING_GUIDE.md#language-server-protocol-lsp).

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## Support

- [GitHub Issues](https://github.com/SwanHtetAungPhyo/tblang/issues)
- [Documentation](https://github.com/SwanHtetAungPhyo/tblang/tree/main/docs)
