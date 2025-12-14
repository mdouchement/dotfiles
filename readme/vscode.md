# VSCode

Extensions:

- Name: Go\
  Id: golang.go\
  Description: Rich Go language support for Visual Studio Code\
  Version: 0.41.4\
  Publisher: golang\
  VS Marketplace Link: https://open-vsx.org/vscode/item?itemName=golang.Go`

- Name: EditorConfig for VS Code\
  Id: EditorConfig.EditorConfig\
  Description: EditorConfig Support for Visual Studio Code\
  Version: 0.15.1\
  Publisher: EditorConfig

- Name: vscode-base64\
  Id: adamhartford.vscode-base64\
  Description: Base64 encode/decode the current selections.\
  Version: 0.1.0\
  Publisher: adamhartford\
  VS Marketplace Link: https://open-vsx.org/vscode/item?itemName=adamhartford.vscode-base64

- Name: file-icons\
  Id: file-icons.file-icons\
  Description: File-specific icons in VSCode for improved visual grepping.\
  Version: 1.1.0\
  Publisher: File Icons\
  VS Marketplace Link: https://open-vsx.org/vscode/item?itemName=file-icons.file-icons

- Name: Prettify JSON\
  Id: mohsen1.prettify-json\
  Description: Visual Studio Code Prettify JSON Extension\
  Version: 0.0.3\
  Publisher: Mohsen Azimi

- Name: Todo Tree\
  Id: Gruntfuggly.todo-tree\
  Description: Show TODO, FIXME, etc. comment tags in a tree view\
  Version: 0.0.215\
  Publisher: Gruntfuggly\
  VS Marketplace Link: https://open-vsx.org/vscode/item?itemName=Gruntfuggly.todo-tree

- Name: Bracket Pair Colorizer 2\
  Id: CoenraadS.bracket-pair-colorizer-2\
  Description: A customizable extension for colorizing matching brackets\
  Version: 0.1.4\
  Publisher: CoenraadS\
  VS Marketplace Link: https://open-vsx.org/vscode/item?itemName=CoenraadS.bracket-pair-colorizer-2\

Settings:

`user/settings.json`

```json
{
    "explorer.autoReveal": false,
    "files.trimTrailingWhitespace": true,
    "go.lintTool": "revive",
    "git.openRepositoryInParentFolders": "never",
    "files.associations": {
        "*.tgo": "tengo"
    },
    "go.coverOnTestPackage": false,
    "go.editorContextMenuCommands": {
        "testAtCursor": false,
        "generateTestForFunction": false,
        "testCoverage": false,
        "playground": false,
        "debugTestAtCursor": false
    },
    "go.playground": {
        "openbrowser": false,
        "share": false,
        "run": false
    },
    "go.showWelcome": false,
    "go.survey.prompt": false,
    "go.formatTool": "gofumpt",
    "workbench.startupEditor": "none",
    "todohighlight.defaultStyle": {},
    "todo-tree.highlights.useColourScheme": true,
    "todo-tree.highlights.backgroundColourScheme": [
        "red",
        "C71585",
        "FF8C00",
        "lime",
        "blue",
        "indigo",
        "violet"
    ],
    "todo-tree.regex.regex": "(//|--|#|<!--|;|/\\*|^|^[ \\t]*(-|\\d+.)).*[^\\w]($TAGS)([^\\p{L}]|$)",
    "editor.fontFamily": "Source Code Pro for Powerline"
}
```

Other variant
```json
{
    "workbench.startupEditor": "newUntitledFile",
    "workbench.colorTheme": "Monokai",
    "editor.fontFamily": "Source Code Pro for Powerline",
    "editor.fontSize": 16,
    "workbench.iconTheme": "file-icons",
    "explorer.confirmDragAndDrop": false,
    "explorer.autoReveal": false,
    "go.editorContextMenuCommands": {
        "toggleTestFile": false,
        "addTags": true,
        "removeTags": false,
        "fillStruct": false,
        "testAtCursor": false,
        "testFile": false,
        "testPackage": false,
        "generateTestForFunction": false,
        "generateTestForFile": false,
        "generateTestForPackage": false,
        "addImport": true,
        "testCoverage": false,
        "playground": false,
        "debugTestAtCursor": false,
        "benchmarkAtCursor": false
    },
    "go.toolsManagement.autoUpdate": true,
    "go.lintOnSave": "package",
    "go.useLanguageServer": true,
    "gopls": {
        "formatting.gofumpt": true,
        "formatting.local": "stash.ovh.net/",
        "analyses": {
            "unusedparams": false
        }
    },
    "go.lintTool": "golangci-lint",
    "go.formatTool": "gofumpt",
    "go.formatFlags": [
        "-w"
    ],
    "security.workspace.trust.untrustedFiles": "open",
    "git.openRepositoryInParentFolders": "never",
    "todo-tree.highlights.useColourScheme": true,
    "todo-tree.highlights.backgroundColourScheme": [
        "red",
        "C71585",
        "FF8C00",
        "32CD32",
        "blue",
        "indigo",
        "violet"
    ],
    "todo-tree.regex.regex": "(//|--|#|<!--|;|/\\*|^|^[ \\t]*(-|\\d+.))[^\\p{L}]*($TAGS)[^\\p{L}]",
    "php.validate.enable": false,
    "php.suggest.basic": false,
    "go.playground": {
        "openbrowser": false,
        "share": false,
        "run": false
    },
    "go.enableCodeLens": {
        "runtest": false
    },
    "[tengo]": {},
    "files.associations": {
        "*.tgo": "tengo"
    },
    "go.coverOnTestPackage": false,
}
```
