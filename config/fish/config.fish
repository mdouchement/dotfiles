set fish_greeting ""
# set fish_private_mode 1 # Start in private mode
set -gx EDITOR vim
set -gx LANG en_US.utf8

fish_add_path ~/bin



if command -sq tozai
  tozai -hook | source
end

function fish_prompt_loading_indicator -a last_prompt
  echo -n "$last_prompt" | sed -r 's/\x1B\[[0-9;]*[JKmsu]//g' | read -zl uncolored_last_prompt
  echo -n (set_color brblack)"$uncolored_last_prompt"(set_color normal)
end

function fish_right_prompt_loading_indicator -a last_prompt
  echo -n "$last_prompt" | sed -r 's/\x1B\[[0-9;]*[JKmsu]//g' | read -zl uncolored_last_prompt
  echo -n (set_color brblack)"$uncolored_last_prompt"(set_color normal)
end



if test -e ~/.hook.fish
  source ~/.hook.fish
end

if status is-interactive && command -sq direnv
  direnv hook fish | source
end


#alias gpg='LANG=en_US gpg'
alias ll='eza -al --group-directories-first'
alias tree='eza --tree'
alias rg='rg -g "!*GODOC.md" -g "!/vendor/"'

if test -e ~/.extra.fish
  source ~/.extra.fish
end
