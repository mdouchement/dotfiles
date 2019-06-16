# Git

- Prune local branches

```sh
git remote prune origin
```

- Disable `less` style on `git branch` command

```sh
git config --global pager.branch false

# NB: You can change the "branch" in the previous line by any other git commands
```

- Reuse recorded resolution of conflicted merges

```sh
git config --global rerere.enabled true
```
