# TBLang Syntax Highlighting Guide

This guide explains how to create syntax highlighting for TBLang in various editors and IDEs.

## Table of Contents

1. [VS Code Extension](#vs-code-extension)
2. [TextMate Grammar](#textmate-grammar)
3. [Sublime Text](#sublime-text)
4. [Vim/Neovim](#vimneovim)
5. [IntelliJ IDEA](#intellij-idea)
6. [Language Server Protocol (LSP)](#language-server-protocol)

---

## VS Code Extension

VS Code is the most popular editor, so we'll start here.

### Project Structure

```
tblang-vscode/
├── package.json
├── syntaxes/
│   └── tblang.tmLanguage.json
├── language-configuration.json
├── themes/
│   └── tblang-dark.json
├── snippets/
│   └── tblang.json
└── README.md
```

### Step 1: Create Extension Scaffold

```bash
# Install Yeoman and VS Code Extension Generator
npm install -g yo generator-code

# Generate extension
yo code

# Choose:
# - New Language Support
# - Language id: tblang
# - Language name: TBLang
# - File extensions: .tbl
# - Scope name: source.tblang
```

### Step 2: Create TextMate Grammar

Create `syntaxes/tblang.tmLanguage.json`:

```json
{
  "$schema": "https://raw.githubusercontent.com/martinring/tmlanguage/master/tmlanguage.json",
  "name": "TBLang",
  "scopeName": "source.tblang",
  "patterns": [
    { "include": "#comments" },
    { "include": "#keywords" },
    { "include": "#strings" },
    { "include": "#numbers" },
    { "include": "#booleans" },
    { "include": "#operators" },
    { "include": "#functions" },
    { "include": "#variables" },
    { "include": "#resource-types" },
    { "include": "#punctuation" }
  ],
  "repository": {
    "comments": {
      "patterns": [
        {
          "name": "comment.line.double-slash.tblang",
          "match": "//.*$"
        },
        {
          "name": "comment.block.tblang",
          "begin": "/\\*",
          "end": "\\*/"
        }
      ]
    },
    "keywords": {
      "patterns": [
        {
          "name": "keyword.control.tblang",
          "match": "\\b(for|in|if|else|while|break|continue|return)\\b"
        },
        {
          "name": "keyword.other.tblang",
          "match": "\\b(declare|cloud_vendor)\\b"
        }
      ]
    },
    "strings": {
      "patterns": [
        {
          "name": "string.quoted.double.tblang",
          "begin": "\"",
          "end": "\"",
          "patterns": [
            {
              "name": "constant.character.escape.tblang",
              "match": "\\\\."
            }
          ]
        },
        {
          "name": "string.quoted.single.tblang",
          "begin": "'",
          "end": "'",
          "patterns": [
            {
              "name": "constant.character.escape.tblang",
              "match": "\\\\."
            }
          ]
        }
      ]
    },
    "numbers": {
      "patterns": [
        {
          "name": "constant.numeric.tblang",
          "match": "\\b-?[0-9]+(\\.[0-9]+)?\\b"
        }
      ]
    },
    "booleans": {
      "patterns": [
        {
          "name": "constant.language.boolean.tblang",
          "match": "\\b(true|false)\\b"
        }
      ]
    },
    "operators": {
      "patterns": [
        {
          "name": "keyword.operator.assignment.tblang",
          "match": "="
        },
        {
          "name": "keyword.operator.comparison.tblang",
          "match": "(==|!=|<|>|<=|>=)"
        },
        {
          "name": "keyword.operator.logical.tblang",
          "match": "(&&|\\|\\||!)"
        },
        {
          "name": "keyword.operator.arithmetic.tblang",
          "match": "(\\+|-|\\*|/|%)"
        }
      ]
    },
    "functions": {
      "patterns": [
        {
          "name": "entity.name.function.tblang",
          "match": "\\b([a-zA-Z_][a-zA-Z0-9_]*)\\s*(?=\\()"
        }
      ]
    },
    "variables": {
      "patterns": [
        {
          "name": "variable.other.tblang",
          "match": "\\b[a-zA-Z_][a-zA-Z0-9_]*\\b"
        }
      ]
    },
    "resource-types": {
      "patterns": [
        {
          "name": "support.type.tblang",
          "match": "\\b(vpc|subnet|security_group|ec2|instance|route_table|internet_gateway)\\b"
        }
      ]
    },
    "punctuation": {
      "patterns": [
        {
          "name": "punctuation.separator.tblang",
          "match": "[,;:]"
        },
        {
          "name": "punctuation.section.braces.tblang",
          "match": "[{}]"
        },
        {
          "name": "punctuation.section.brackets.tblang",
          "match": "[\\[\\]]"
        },
        {
          "name": "punctuation.section.parens.tblang",
          "match": "[()]"
        },
        {
          "name": "punctuation.accessor.tblang",
          "match": "\\."
        }
      ]
    }
  }
}
```

### Step 3: Language Configuration

Create `language-configuration.json`:

```json
{
  "comments": {
    "lineComment": "//",
    "blockComment": ["/*", "*/"]
  },
  "brackets": [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  "autoClosingPairs": [
    { "open": "{", "close": "}" },
    { "open": "[", "close": "]" },
    { "open": "(", "close": ")" },
    { "open": "\"", "close": "\"" },
    { "open": "'", "close": "'" }
  ],
  "surroundingPairs": [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["\"", "\""],
    ["'", "'"]
  ],
  "folding": {
    "markers": {
      "start": "^\\s*//\\s*#?region\\b",
      "end": "^\\s*//\\s*#?endregion\\b"
    }
  },
  "wordPattern": "(-?\\d*\\.\\d\\w*)|([^\\`\\~\\!\\@\\#\\%\\^\\&\\*\\(\\)\\-\\=\\+\\[\\{\\]\\}\\\\\\|\\;\\:\\'\\\"\\,\\.\\<\\>\\/\\?\\s]+)",
  "indentationRules": {
    "increaseIndentPattern": "^((?!\\/\\/).)*(\\{[^}\"'`]*|\\([^)\"'`]*|\\[[^\\]\"'`]*)$",
    "decreaseIndentPattern": "^((?!.*?\\/\\*).*\\*/)?\\s*[\\)\\}\\]].*$"
  }
}
```

### Step 4: Code Snippets

Create `snippets/tblang.json`:

```json
{
  "Cloud Vendor": {
    "prefix": "cloud",
    "body": [
      "cloud_vendor \"${1:aws}\" {",
      "\tregion = \"${2:us-east-1}\"",
      "\tprofile = \"${3:default}\"",
      "}"
    ],
    "description": "Cloud vendor configuration"
  },
  "Declare Variable": {
    "prefix": "declare",
    "body": [
      "declare ${1:variable_name} = ${2:value};"
    ],
    "description": "Variable declaration"
  },
  "VPC Resource": {
    "prefix": "vpc",
    "body": [
      "declare ${1:vpc_name} = vpc(\"${2:vpc-name}\", {",
      "\tcidr_block: \"${3:10.0.0.0/16}\"",
      "\tenable_dns_hostnames: ${4:true}",
      "\tenable_dns_support: ${5:true}",
      "});"
    ],
    "description": "VPC resource"
  },
  "Subnet Resource": {
    "prefix": "subnet",
    "body": [
      "declare ${1:subnet_name} = subnet(\"${2:subnet-name}\", {",
      "\tvpc_id: ${3:vpc_reference}",
      "\tcidr_block: \"${4:10.0.1.0/24}\"",
      "\tavailability_zone: \"${5:us-east-1a}\"",
      "});"
    ],
    "description": "Subnet resource"
  },
  "For Loop": {
    "prefix": "for",
    "body": [
      "for ${1:item} in ${2:collection} {",
      "\t${3:// statements}",
      "}"
    ],
    "description": "For loop"
  },
  "Object Literal": {
    "prefix": "obj",
    "body": [
      "{",
      "\t${1:key}: ${2:value}",
      "}"
    ],
    "description": "Object literal"
  },
  "Array Literal": {
    "prefix": "arr",
    "body": [
      "[${1:items}]"
    ],
    "description": "Array literal"
  }
}
```

### Step 5: Package Configuration

Create `package.json`:

```json
{
  "name": "tblang",
  "displayName": "TBLang",
  "description": "TBLang language support for VS Code",
  "version": "1.0.0",
  "publisher": "tblang",
  "engines": {
    "vscode": "^1.75.0"
  },
  "categories": [
    "Programming Languages"
  ],
  "keywords": [
    "tblang",
    "infrastructure",
    "iac",
    "cloud"
  ],
  "contributes": {
    "languages": [
      {
        "id": "tblang",
        "aliases": ["TBLang", "tblang"],
        "extensions": [".tbl"],
        "configuration": "./language-configuration.json",
        "icon": {
          "light": "./icons/tblang-light.png",
          "dark": "./icons/tblang-dark.png"
        }
      }
    ],
    "grammars": [
      {
        "language": "tblang",
        "scopeName": "source.tblang",
        "path": "./syntaxes/tblang.tmLanguage.json"
      }
    ],
    "snippets": [
      {
        "language": "tblang",
        "path": "./snippets/tblang.json"
      }
    ],
    "themes": [
      {
        "label": "TBLang Dark",
        "uiTheme": "vs-dark",
        "path": "./themes/tblang-dark.json"
      }
    ]
  },
  "repository": {
    "type": "git",
    "url": "https://github.com/SwanHtetAungPhyo/tblang-vscode"
  }
}
```

### Step 6: Testing

```bash
# Install dependencies
npm install

# Open in VS Code
code .

# Press F5 to launch Extension Development Host
# Open a .tbl file to test syntax highlighting
```

### Step 7: Publishing

```bash
# Install vsce
npm install -g @vscode/vsce

# Package extension
vsce package

# Publish to marketplace
vsce publish
```

---

## TextMate Grammar

TextMate grammars work in VS Code, Sublime Text, and other editors.

### Complete Grammar File

Create `tblang.tmLanguage`:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>fileTypes</key>
    <array>
        <string>tbl</string>
    </array>
    <key>name</key>
    <string>TBLang</string>
    <key>scopeName</key>
    <string>source.tblang</string>
    <key>patterns</key>
    <array>
        <dict>
            <key>include</key>
            <string>#comments</string>
        </dict>
        <dict>
            <key>include</key>
            <string>#keywords</string>
        </dict>
        <dict>
            <key>include</key>
            <string>#strings</string>
        </dict>
        <dict>
            <key>include</key>
            <string>#numbers</string>
        </dict>
    </array>
    <key>repository</key>
    <dict>
        <key>comments</key>
        <dict>
            <key>patterns</key>
            <array>
                <dict>
                    <key>name</key>
                    <string>comment.line.double-slash.tblang</string>
                    <key>match</key>
                    <string>//.*$</string>
                </dict>
            </array>
        </dict>
        <key>keywords</key>
        <dict>
            <key>patterns</key>
            <array>
                <dict>
                    <key>name</key>
                    <string>keyword.control.tblang</string>
                    <key>match</key>
                    <string>\b(for|in|declare|cloud_vendor)\b</string>
                </dict>
            </array>
        </dict>
        <key>strings</key>
        <dict>
            <key>patterns</key>
            <array>
                <dict>
                    <key>name</key>
                    <string>string.quoted.double.tblang</string>
                    <key>begin</key>
                    <string>"</string>
                    <key>end</key>
                    <string>"</string>
                </dict>
            </array>
        </dict>
        <key>numbers</key>
        <dict>
            <key>patterns</key>
            <array>
                <dict>
                    <key>name</key>
                    <string>constant.numeric.tblang</string>
                    <key>match</key>
                    <string>\b-?[0-9]+(\.[0-9]+)?\b</string>
                </dict>
            </array>
        </dict>
    </dict>
</dict>
</plist>
```

---

## Sublime Text

### Installation

1. Create package directory:
```bash
mkdir -p ~/Library/Application\ Support/Sublime\ Text/Packages/TBLang
```

2. Create `TBLang.sublime-syntax`:

```yaml
%YAML 1.2
---
name: TBLang
file_extensions: [tbl]
scope: source.tblang

contexts:
  main:
    - include: comments
    - include: keywords
    - include: strings
    - include: numbers
    - include: booleans
    - include: operators
    - include: functions
    - include: resource-types

  comments:
    - match: '//'
      scope: punctuation.definition.comment.tblang
      push:
        - meta_scope: comment.line.double-slash.tblang
        - match: $
          pop: true
    
    - match: '/\*'
      scope: punctuation.definition.comment.begin.tblang
      push:
        - meta_scope: comment.block.tblang
        - match: '\*/'
          scope: punctuation.definition.comment.end.tblang
          pop: true

  keywords:
    - match: '\b(for|in|if|else|while|break|continue|return)\b'
      scope: keyword.control.tblang
    - match: '\b(declare|cloud_vendor)\b'
      scope: keyword.other.tblang

  strings:
    - match: '"'
      scope: punctuation.definition.string.begin.tblang
      push:
        - meta_scope: string.quoted.double.tblang
        - match: '\\.'
          scope: constant.character.escape.tblang
        - match: '"'
          scope: punctuation.definition.string.end.tblang
          pop: true

  numbers:
    - match: '\b-?[0-9]+(\.[0-9]+)?\b'
      scope: constant.numeric.tblang

  booleans:
    - match: '\b(true|false)\b'
      scope: constant.language.boolean.tblang

  operators:
    - match: '='
      scope: keyword.operator.assignment.tblang
    - match: '(==|!=|<|>|<=|>=)'
      scope: keyword.operator.comparison.tblang

  functions:
    - match: '\b([a-zA-Z_][a-zA-Z0-9_]*)\s*(?=\()'
      scope: entity.name.function.tblang

  resource-types:
    - match: '\b(vpc|subnet|security_group|ec2|instance)\b'
      scope: support.type.tblang
```

3. Create `TBLang.sublime-completions`:

```json
{
  "scope": "source.tblang",
  "completions": [
    {
      "trigger": "cloud_vendor",
      "contents": "cloud_vendor \"${1:aws}\" {\n\tregion = \"${2:us-east-1}\"\n\tprofile = \"${3:default}\"\n}"
    },
    {
      "trigger": "declare",
      "contents": "declare ${1:name} = ${2:value};"
    },
    {
      "trigger": "for",
      "contents": "for ${1:item} in ${2:collection} {\n\t${3}\n}"
    },
    {
      "trigger": "vpc",
      "contents": "vpc(\"${1:name}\", {\n\tcidr_block: \"${2:10.0.0.0/16}\"\n})"
    }
  ]
}
```

---

## Vim/Neovim

### Installation

Create `~/.vim/syntax/tblang.vim` or `~/.config/nvim/syntax/tblang.vim`:

```vim
" Vim syntax file
" Language: TBLang
" Maintainer: TBLang Team
" Latest Revision: 2024-11-14

if exists("b:current_syntax")
  finish
endif

" Keywords
syn keyword tblangKeyword declare cloud_vendor for in
syn keyword tblangControl if else while break continue return
syn keyword tblangBoolean true false

" Comments
syn match tblangComment "//.*$"
syn region tblangBlockComment start="/\*" end="\*/"

" Strings
syn region tblangString start='"' end='"' skip='\\"'
syn region tblangString start="'" end="'" skip="\\'"

" Numbers
syn match tblangNumber '\v<-?\d+(\.\d+)?>'

" Functions
syn match tblangFunction '\v<\w+\ze\('

" Resource Types
syn keyword tblangResourceType vpc subnet security_group ec2 instance

" Operators
syn match tblangOperator "="
syn match tblangOperator "=="
syn match tblangOperator "!="
syn match tblangOperator "<"
syn match tblangOperator ">"
syn match tblangOperator "<="
syn match tblangOperator ">="

" Highlighting
hi def link tblangKeyword Keyword
hi def link tblangControl Conditional
hi def link tblangBoolean Boolean
hi def link tblangComment Comment
hi def link tblangBlockComment Comment
hi def link tblangString String
hi def link tblangNumber Number
hi def link tblangFunction Function
hi def link tblangResourceType Type
hi def link tblangOperator Operator

let b:current_syntax = "tblang"
```

Create `~/.vim/ftdetect/tblang.vim`:

```vim
au BufRead,BufNewFile *.tbl set filetype=tblang
```

---

## IntelliJ IDEA

### Using Grammar-Kit

1. Install Grammar-Kit plugin
2. Create BNF grammar file
3. Generate parser and lexer
4. Create syntax highlighter

This is more complex and requires Java knowledge. Consider using TextMate bundle support instead.

---

## Language Server Protocol (LSP)

For advanced features like autocomplete, go-to-definition, and diagnostics, implement an LSP server.

### Basic LSP Server Structure

```go
package main

import (
    "context"
    "github.com/sourcegraph/go-lsp"
    "github.com/sourcegraph/jsonrpc2"
)

type Server struct{}

func (s *Server) Initialize(ctx context.Context, params *lsp.InitializeParams) (*lsp.InitializeResult, error) {
    return &lsp.InitializeResult{
        Capabilities: lsp.ServerCapabilities{
            TextDocumentSync: &lsp.TextDocumentSyncOptionsOrKind{
                Options: &lsp.TextDocumentSyncOptions{
                    OpenClose: true,
                    Change:    lsp.TDSKFull,
                },
            },
            CompletionProvider: &lsp.CompletionOptions{
                TriggerCharacters: []string{"."},
            },
            HoverProvider: true,
        },
    }, nil
}

func main() {
    handler := jsonrpc2.HandlerWithError((&Server{}).handle)
    <-jsonrpc2.NewConn(
        context.Background(),
        jsonrpc2.NewBufferedStream(stdrwc{}, jsonrpc2.VSCodeObjectCodec{}),
        handler,
    ).DisconnectNotify()
}
```

---

## Quick Start: VS Code Extension

The fastest way to get syntax highlighting working:

```bash
# 1. Create extension directory
mkdir tblang-vscode && cd tblang-vscode

# 2. Create package.json (see above)

# 3. Create syntaxes/tblang.tmLanguage.json (see above)

# 4. Create language-configuration.json (see above)

# 5. Test locally
code .
# Press F5

# 6. Package and share
npm install -g @vscode/vsce
vsce package
# Share the .vsix file
```

---

## Resources

- [VS Code Language Extensions](https://code.visualstudio.com/api/language-extensions/overview)
- [TextMate Grammar](https://macromates.com/manual/en/language_grammars)
- [Sublime Text Syntax](https://www.sublimetext.com/docs/syntax.html)
- [Vim Syntax Highlighting](https://vim.fandom.com/wiki/Creating_your_own_syntax_files)
- [Language Server Protocol](https://microsoft.github.io/language-server-protocol/)

---

## Next Steps

1. **Start with VS Code** - Easiest and most popular
2. **Test thoroughly** - Use real TBLang files
3. **Add snippets** - Improve developer experience
4. **Publish** - Share with the community
5. **Add LSP** - For advanced features
6. **Support other editors** - Vim, Sublime, etc.

For questions or contributions, visit the [TBLang repository](https://github.com/SwanHtetAungPhyo/tblang).
