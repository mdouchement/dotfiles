" Nerd Tree
" The NERD tree allows you to explore your filesystem and to open files and directories.
"---------------------------------------------------------------------------------------

" Display/hide NERDTree
map <C-n> :NERDTreeTabsToggle<CR>

" Ctrl+x & arrow Navigation
nmap <silent> <C-x><up> :wincmd k<CR>
nmap <silent> <C-x><down> :wincmd j<CR>
nmap <silent> <C-x><left> :wincmd h<CR>
nmap <silent> <C-x><right> :wincmd l<CR>
