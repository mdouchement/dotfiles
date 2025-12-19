status is-interactive && command -sq zoxide
or exit 0

function z
  cd (command zoxide query --exclude (builtin pwd) -- $argv)
end

function zadd
  command zoxide add $PWD
end

function zdadd
  for x in *
    if test -d $x
      echo "zoxide add $x"
      command zoxide add "$(builtin pwd)/$x"
    end
  end
end
