# ~/.hook.fish

- Homebrew

```fish
set -gx HOMEBREW_NO_ANALYTICS 1

# Hombrew & Hombrecask in the $HOME folder
set -gx HOMEBREW_CASK_OPTS "--appdir=~/Applications"
set -gx PATH $HOME/homebrew/bin $PATH

alias rbrew='env HOMEBREW_CASK_OPTS="--appdir=/Applications" brew'
```

https://www.scivision.dev/macos-homebrew-non-sudo/
=> `mkdir homebrew && curl -L https://github.com/Homebrew/brew/tarball/master | tar xz --strip 1 -C homebrew`


- Golang version manager

https://github.com/kevincobain2000/gobrew/releases
```sh
set -gx GOROOT ~/.gobrew/current/go
fish_add_path ~/.gobrew/current/bin
fish_add_path ~/.gobrew/bin
```
