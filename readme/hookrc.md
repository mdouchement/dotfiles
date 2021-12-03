# ~/.hookrc

- Homebrew

```sh
export HOMEBREW_NO_ANALYTICS=1

# https://docs.brew.sh/Shell-Completion
if type brew &>/dev/null; then
  FPATH=$(brew --prefix)/share/zsh/site-functions:$FPATH
fi

# Hombrew & Hombrecask in the $HOME folder
export HOMEBREW_CASK_OPTS="--appdir=~/Applications"
export PATH=${HOME}/homebrew/bin:${PATH}

alias rbrew='HOMEBREW_CASK_OPTS="--appdir=/Applications" brew'
```

https://www.scivision.dev/macos-homebrew-non-sudo/
=> `mkdir homebrew && curl -L https://github.com/Homebrew/brew/tarball/master | tar xz --strip 1 -C homebrew`


- ZSH (completion)

```sh
# https://gist.github.com/ctechols/ca1035271ad134841284
autoload -Uz compinit
if [ $(date +'%j') != $(stat -f '%Sm' -t '%j' ~/.zcompdump) ]; then
    compinit;
else
    compinit -C;
fi;
```

- Golang version manager

```sh
# cd ~/bin
# curl -LO https://raw.githubusercontent.com/stefanmaric/g/master/bin/g
# chmod +x g

export GOROOT=${HOME}/.go
function g_go() {
  GOPATH=${HOME}/.gopath ${HOME}/bin/g $@
}
# Override git alias
alias g='g_go'
```

- Direnv loader

```sh
eval "$(direnv hook zsh)"
```
