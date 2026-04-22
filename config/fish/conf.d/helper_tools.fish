function spacer --argument-names 'n'
  if test -z $n
    set n 20
  end

  set i 0
  while test $i -lt $n
    echo ""
    set i (math $i + 1)
  end
end
