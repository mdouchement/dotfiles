" General configurations
"------------------------

set nocompatible                " Don't maintain compat with Vi.
set hidden                      " Allow buffer change w/o saving
set autoread                    " Load file from disk, ie for git reset
set nocompatible		            " Not concerned with vi compatibility
set lazyredraw                  " Don't update while executing macros
set backspace=indent,eol,start	" Sane backspace behavior
set history=1000  		          " Remember last 1000 commands
set scrolloff=4                 " Keep at least 4 lines below cursor
set expandtab                   " Convert <tab> to spaces (2 or 4)
set tabstop=2                   " Two spaces per tab as default
set shiftwidth=2                "     then override with per filteype
set softtabstop=2               "     specific settings via autocmd
set encoding=utf-8

" Use relativenumber, but provide easy toggle
set relativenumber
nnoremap <leader>nu :call ToggleRelativeNumber()<cr>
nnoremap <leader>rn :call ToggleRelativeNumber()<cr>
function! ToggleRelativeNumber()
  if &relativenumber
    set number
  else
    set relativenumber
  endif
endfunction

" Disable sound/visual bell on errors
" May need additional config about audible bell
set t_vb=

" Set modeline to 1 to allow rcfiles to be recognized as vim files
set modelines=1

" I save constantly and hate swap files in my prject dirs
set nobackup
set nowritebackup
set noswapfile

" Replace the <cedit> command to pop into the 'command-line window'
set cedit=<C-Y>

" Automatically remove trailing whitespace
autocmd BufWritePre * :%s/\s\+$//e

" vim:ft=vim
