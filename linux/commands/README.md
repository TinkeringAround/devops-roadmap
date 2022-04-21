# Commands

A collection of common linux commands, pipes, shortcuts...

## Semantics

- `()` Optional
- `{}` Required
- `{file}` File name
- `{dir}` Directory name
- `{target}` Target directory
- `{content}` Some text Content
- `{number}` Some number
- `{user}` Name of a user
- `{group}` Name of a group
- `{permission}` [Permission code](#Permissions) like `777`
- `{package}` Package name

## Commands

### File/Directory Management

| Command                                | Description                                                                                             |
|----------------------------------------|---------------------------------------------------------------------------------------------------------|
| `pwd`                                  | Print current working directory                                                                         |
| `ls (-al) ({dir/file})`                | List content of current working directory or file (`-al` = with Permissions)                            |
| `cd {dir}`                             | Change to a directory                                                                                   |
| `mkdir {dir}`                          | Create a directory                                                                                      |
| `rmdir {dir}`                          | Remove an empty directory                                                                               |
| `updatedb`                             | Reindex all files (use when not finding files via locate)                                               |
| `locate {file/dir}`                    | Try to find a file/directory in the file system (use updatedb if nothing is found to reindex all files) |
| `mv {file/dir} {target}`               | Move a file/directory to target location (directory) or rename it                                       |
| `rm {file} (-R) (-f)`                  | Removing a file (`-R` = Recursively) (`-f` = force)                                                     |
| `cp {file} {target}`                   | Copy a file to a target directory                                                                       |
| `echo "${ENV}"`                        | Print a environment variable                                                                            |
| `echo "{content}" > {file}`            | Add text to a file overwriting existing content                                                         |
| `echo "{content}" >> {file}`           | Append text to a file not overwriting existing content                                                  |
| `cat {file}`                           | Read files content                                                                                      |
| `cat {file}` &#124; `grep "{content}"` | Search text inside a file                                                                               |

### Administration

| Command                                | Description                                                                                             |
|----------------------------------------|---------------------------------------------------------------------------------------------------------|
| `printenv`                             | Print environment variables                                                                             |
| `passwd`                               | Changing the current user password                                                                      |
| `useradd {user}`                       | Creating a new user                                                                                     |
| `groupadd {group}`                     | Creating a group                                                                                        |
| `usermod {user} -a -G {group}`         | Adding `-a` a user to a group `-G`                                                                      |
| `usermod -a -G sudo {user}`            | Add `-a` a user to the `sudo` group `-G`                                                                |
| `chmod {permission} {file/dir}`        | Change [permission](#Permissions) for a file/directory                                                  |
| `chown {user} {file/dir} (-R)`         | Change Ownership of file/dir (`-R` = Recursively)                                                       |
| `cut -d: -f1 < /etc/passwd`            | List all user accounts                                                                                  |
| `apt install {package}`                | Install a package                                                                                       |
| `apt remove {package}`                 | Remove a package                                                                                        |

### Filesystem + Storage

| Command   | Description                                          |
|-----------|------------------------------------------------------|
| `lsblk`   | List block devices                                   |
| `df (-h)` | Information about partitions (`-h` = Human readable) |

### Utilities

| Command                                | Description                                                                                             |
|----------------------------------------|---------------------------------------------------------------------------------------------------------|
| `clear`                                | Clear console                                                                                           |
| `history (-c)`                         | Show history of commands used (`-c` = clear history)                                                    |
| `history` &#124; `grep "content"`      | Search for a command in the history                                                                     |

## Permissions

When listing permissions of a file/directory via `ls -al {file/dir}` the order of printed permissions
is `User - Group - Everyone` like `-rw-r--r--`.

When setting permissions via `chmod {permission} {file/dir}` use the permission codes like `777`.

| #   | Short | Permissions           |
|-----|-------|-----------------------|
| 0   | `---` | None                  |
| 1   | `--X` | Execute only          |
| 2   | `-WX` | Write only            |
| 3   | `R--` | Write & Execute       |
| 4   | `R-X` | Read only             |
| 5   | `RW-` | Read & Execute        |
| 6   | `RW-` | Read & Write          |
| 7   | `RWX` | Read, Write & Execute |

## Pipes

| Pipe                      | Description                                          |
|---------------------------|------------------------------------------------------|
| `grep {content}`          | Search for string                                    |
| `awk '{print ${number}}'` | Print selected column                                |
| `sort`                    | Sorting                                              |
| `xargs`                   | Transforms column based output into row based output |

## Shortcuts

| Shortcut | Description                          |
|----------|--------------------------------------|
| `Strg+C` | Cancels the current input of the CLI |
