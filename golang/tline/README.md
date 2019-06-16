# tline

Simple tmux status bar inspired by [limeline](https://github.com/mcartmell/limeline).

## Features

- Load average (load1 of CPU)
- Current date and time

## Useful commands

```sh
kill -9 $(cat /tmp/tline.pid)
kill -9 $(cat $TMPDIR/tline.pid)
```