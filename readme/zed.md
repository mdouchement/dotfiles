
This Gist may be more up-to-date of the one below: https://gist.github.com/mdouchement/eb8684be330a09bebe66dbf91af16516
```json
// Zed settings
//
// For information on how to configure Zed, see the Zed
// documentation: https://zed.dev/docs/configuring-zed
// Other guide: https://www.kevnu.com/en/posts/10
//
// To see all of Zed's default settings without changing your
// custom settings, run the `open default settings` command
// from the command palette or from `Zed` application menu.
{
  "ui_font_size": 16,
  "buffer_font_size": 16,
  "buffer_font_family": "SauceCodePro Nerd Font",
  "tab_size": 2,
  "hour_format": "hour24",
  "theme": "monokai Darker Classic",
  "telemetry": {
    "diagnostics": false,
    "metrics": false
  },
  "format_on_save": "off", // Global formating for any type of files breaks a lot of things for existing projects.
  "languages": {
    "Go": {
      "format_on_save": "on",
      "formatter": "language_server",
      "code_actions_on_format": {
        "source.organizeImports": true
      }
    }
  },
  "lsp": {
    "gopls": {
      // https://github.com/golang/tools/blob/master/gopls/doc/settings.md
      "initialization_options": {
        "directoryFilters": ["-**/node_modules", "-**/.git"],
        "formatting.gofumpt": true,
        "symbolScope": "workspace",
        "staticcheck": true,
        "templateExtensions": ["tmpl"],
        "experimentalPostfixCompletions": true,
        "analyses": {
          "nilness": true,
          "unusedparams": true,
          "unusedvariable": true,
          "unusedwrite": true,
          "useany": true,
          "ST1000": false // https://go.dev/gopls/analyzers#st1000-incorrect-or-missing-package-comment
        },
        "codelenses": {
          "gc_details": true
        },
        "hints": {
          "assignVariableTypes": true,
          "compositeLiteralFields": true,
          "compositeLiteralTypes": true,
          "constantValues": true,
          "functionTypeParameters": true,
          "parameterNames": true,
          "rangeVariableTypes": true
        }
      }
    }
  }
}

```
