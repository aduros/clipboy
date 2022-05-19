# Clipboy

Clipboy is a cross-platform, minimal clipboard manager.

## Features

- Works on Linux and Mac, maybe even Windows(?).
- Records clipboard history and exposes it to scripts.
- Doesn't do much else.

## Usage

1. Add `clipboy watch &` to your startup script to start the daemon.

2. Run `clipboy list` to get the clipboard history.

## Example

```shell
# Show a menu to update the clipboard from history
clipboy list | fzf --read0 --print0 | clipboy paste
```
