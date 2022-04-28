# Bash Scripting

BASH - Bourne Again Shell

# Variables

````shell
NAME="Thomas"
AGE=27

echo "Hello, my name is $NAME and I am $AGE years old. Nice to meet you!"
````

## Positional Arguments

Positional arguments are everything written after the bash file name and are automatically assigned to variables
starting by `$1`.

````shell
./printArgs.sh Hello World
````

```shell
echo "First Argument: $1" # Hello
echo "Second Argument: $2" # World
```

## Reading User Input

````shell
read -r NAME
echo "$NAME"
````

## If/Else

With if/else there can be values compared against each other or boolean assertions be made like check whether a path
leads to a directory or a file.

### Comparisons

````shell
if [ $AGE -lt 18 ]
then 
  echo "You are underage still!"
fi
````

| Comparator | Description                     |
|------------|---------------------------------|
| `eq`       | equal                           |
| `ne`       | not equal                       |
| `gt`       | first value is greater          |
| `ge`       | first value is greater or equal |
| `lt`       | first value is less             |
| `le`       | first value is less or equal    |

### Assertions

````shell
FILE="90DaysOfDevOps.txt"
if [ -f "$FILE" ]
then 
  echo "$FILE is a file"
else 
  echo "$FILE is not a file"
fi
````

| Assertion | Description               |
|-----------|---------------------------|
| `-d`      | file is a directory       |
| `-e`      | file exists               |
| `-f`      | provided string is a file |
| `g`       | group id is set on a file |
| `-r`      | file is readable          |
| `-s`      | file has a non-zero size  |

## Loops

- [For-Loops](https://www.cyberciti.biz/faq/bash-for-loop/)
- [While-Loops](https://www.cyberciti.biz/faq/bash-while-loop/)

## Example: Create user Script

Start Ubuntu and execute bash shell.

```shell
make run
```

Create a user

```shell
/bin/bash create_user.sh
```

Switch to user

```shell
su - Thomas
```

Delete a user

```shell
userdel Thomas
```