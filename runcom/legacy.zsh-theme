if [ -d "${HOME}/.rbenv" ]; then
  function rbenv_prompt_info() {
    echo "‹$(rbenv version-name)›"
  }
else
  function rbenv_prompt_info() {
    echo "‹NA›"
  }
fi

# The prompt

PROMPT='%{$fg[cyan]%}${ZSH_THEME_PREFIX_BLOCK}[%c]>> %{$reset_color%}'

# The right-hand prompt

RPROMPT='${time} %{$fg[yellow]%}$(rbenv_prompt_info)%{$reset_color%} %{$fg[cyan]%}$(git_prompt_info)%{$reset_color%}'

# local time, color coded by last return code
time_enabled="%(?.%{$fg[green]%}.%{$fg[red]%})%*%{$reset_color%}"
time_disabled="%{$fg[green]%}%*%{$reset_color%}"
time=$time_enabled

ZSH_THEME_GIT_PROMPT_PREFIX="%{$fg[red]%}"
ZSH_THEME_GIT_PROMPT_SUFFIX="%{$reset_color%}"
