# ~/.extra.fish

- `eza` a modern version of 'ls'.

https://github.com/eza-community/eza

```fish
alias ll="eza -al --group-directories-first"
```

- `rg` for golang

https://github.com/BurntSushi/ripgrep

```fish
alias rg='~/bin/rg -g "!*GODOC.md" -g "!/vendor/" -g "!/node_modules/"'
```

- Yaegi

```fish
alias yaegi='rlwrap yaegi'
```

- GNU/Linux memory usage

```fish
function mem-usage
  ps -eo size,pid,user,command --sort -size | awk '{ hr=$1/1024 ; printf("%13.2f Mb ",hr) } { for ( x=4 ; x<=NF ; x++ ) { printf("%s ",$x) } print "" }'
end
```

- Golang

```fish
go build -ldflags "-s -w" -o ~/bin/(basename (pwd)) .
```

```fish
function gofinn --argument-names dst src --description 'gofinn /tmp/(basename $PWD) .'
  env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $dst $src
  finn $dst
  rm $dst
end
```

- cURL + TLS check

```fish
function curt
	curl --insecure -v $@ 2>&1 | awk 'BEGIN { cert=0 } /^\* SSL connection/ { cert=1 } /^\*/ { if (cert) print }'
end
```

- Distrobox

```fish
# Icons
# https://starship.rs/presets/nerd-font

distrobox create --name Arch --hostname (hostname) --additional-flags "--env TOZAI_MACHINE_NAME='' --env SHELL=/usr/bin/fish" --image docker.io/library/archlinux:latest --additional-packages "zed"

distrobox create --name f43 --hostname (hostname) --additional-flags "--env TOZAI_MACHINE_NAME='\' --env SHELL=/usr/bin/fish" --image quay.io/fedora/fedora:43
```

```fish
alias dua='distrobox upgrade --all'
alias arch='distrobox enter Arch'
alias zed='distrobox-enter -n Arch -- zeditor'
```

- Taskfile

```fish
function task
  if test -z $argv
    command task -l
    return 1
  end
  command  task $argv
end
```

- Grift

```fish
function grift
  if test -z $argv
    command grift list
    return 1
  end
  command grift $argv
end
```
