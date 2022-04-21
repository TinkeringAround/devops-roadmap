# VIM

Possibly the most common text editor around? A sibling of the UNIX text editor vi from 1976 we get a lot of
functionality with vim.

- Pretty much supported on every single Linux distribution.
- Incredibly powerful!

## Installation

```shell
apt install vim
```

## Usage

First step is to open the file via `vim text.txt` altogether.

### Only Reading

1. Read the file until done
2. Press `escape` to return to normal mode
3. Hit `:q` and press `enter` to exit vim without any changes

### Editing

1. Press `i` to switch to edit mode
2. Make changes, saving in between with `:w`
3. Press `escape` to return to normal mode
4. Hit `:wq` and press `enter` to exit vim

### Copy and Paste

1. Press `yy` to copy a line
2. Press `p` to paste to a new line
3. Press `dd` to delete a line
4. Hit `:wq` and press `enter` to exit vim

### Searching

1. Press `/word` and then `enter`
2. Navigate with `n` to all the results
3. Hit `:q` and press `enter` to exit vim without any changes

### Discarding Changes
1. Hit `:q!` and press `enter` to exit vim and to discard changes
