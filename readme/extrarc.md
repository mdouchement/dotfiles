# ~/.extrarc

- Auto `cd` to your favourite directories

```sh
alias acd='set auto_cd ; cdpath=(${HOME}/workspaces/golang/src/github.com/mdouchement ${HOME}/workspaces/gomodules ${HOME}/workspaces/ruby)'

# Usage
$ acd
$ cd workerpool/
```

- `exa` a modern version of ‘ls’.

https://github.com/ogham/exa

```sh
alias ll="exa -al --group-directories-first"
```

- `rg` for golang

https://github.com/BurntSushi/ripgrep

```sh
alias rg='rg -g "!*GODOC.md" -g "!/vendor/"'
```

- Use Hombrew `curl`

```sh
export PATH=/usr/local/opt/curl/bin:${PATH}
```

- Load Rbenv/Anyenv

```sh
function lrb() {
    if [ -z "$RBENV_ROOT" ]; then
        export RBENV_ROOT="${HOME}/.rbenv"
        export PATH="${RBENV_ROOT}/bin:${PATH}"
        eval "$(rbenv init -)"
    fi
}
alias rbenv='lrb && rbenv'
alias ruby='lrb && ruby'
```

- Exclude `todolist` from `history` tracking

```sh
alias tl=todolist
function zshaddhistory() {
    emulate -L zsh
    if [[ $1 = "todolist"* ]] ; then
        return 1
    fi
    if [[ $1 = "tl"* ]] ; then
        return 1
    fi
}
```

- Load python version manager on-demand

```sh
function lpy() {
  export PYENV_ROOT="${HOME}/.pyenv"
  export PATH="${PYENV_ROOT}/bin:$PATH"

  eval "$(pyenv init -)"
  #eval "$(pyenv virtualenv-init -)"
}
```

- Load `node` on-demand

```sh
function ljs(){
  export NVM_DIR="${HOME}/.nvm"
  [ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"  # This loads nvm
}

function node() {
  if [ -z "$NVM_DIR" ]; then
    ljs
  fi
  $(nvm which node) $@
}
alias nodejs=node
```
