# Dotfiles

[https://github.com/mdouchement/ldt](https://github.com/mdouchement/ldt)

## General usage

```sh
# Global command with YAML configuration

cp dotfiles.sample.yml ~/.config/dotfiles.yml
$ ldt install ~/.config/dotfiles.yml
```

## ZSH

```sh
# Install ZSH
$ chsh -s /bin/zsh
$ curl -L http://install.ohmyz.sh | sh

$ cd dotfiles
$ ldt install ~/.config/dotfiles.yml
```

## Fish

```sh
$ chsh -s /usr/bin/fish

$ cd dotfiles
$ ldt install ~/.config/dotfiles.yml

# https://github.com/catppuccin/fish
$ fisher install catppuccin/fish
$ fish_config theme save "Catppuccin Mocha"
```

## VIM

```sh
$ curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

$ cd dotfiles
$ ldt install ~/.config/dotfiles.yml

$ vim +PlugInstall
```

## Nerd Fonts

- macOS

Just double clic on it and install

- Linux

```sh
cp fonts/*.ttf ~/.local/share/fonts/
fc-cache -vf ~/.local/share/fonts/
```

## Gnome

```sh
gsettings set org.gnome.desktop.default-applications.terminal exec '/usr/bin/konsole'
```

## License

**MIT**
