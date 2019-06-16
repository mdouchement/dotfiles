# Atom

```sh
apm list --installed --bare > ~/.atom/package.list
apm install --packages-file ~/.atom/package.list
```

## Packages

```sh
apm install language-docker
apm install language-fish-shell
apm install language-viml
apm install language-vue

apm install monokai
apm install file-icons

apm install atom-ide-ui
apm install ide-go
```

## Unwanted packages

```sh
apm disable metrics
```

## Cleanup

```sh
apm prune
```
