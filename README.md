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

## VIM

```sh
$ curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

$ cd dotfiles
$ ldt install ~/.config/dotfiles.yml

$ vim +PlugInstall
```

## Powerline Fonts

https://github.com/powerline/fonts/tree/master/SourceCodePro

- macOS

Just double clic on it and install

- Linux

```sh
cp fonts/Source_Code_Pro_Medium_for_Powerline.otf ~/.local/share/fonts/
fc-cache -vf ~/.local/share/fonts/
```

## License

**MIT**
