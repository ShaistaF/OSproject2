# Project 2: Shell Builtins

## Description

For this project we'll be adding commands to a simple shell. 

The shell is already written, but you will choose five (5) shell builtins (or shell-adjacent) commands to rewrite into Go, and integrate into the Go shell.

There are many builtins or shell-adjacent commands to pick from: 
[Bourne Shell Builtins](https://www.gnu.org/software/bash/manual/html_node/Bourne-Shell-Builtins.html), 
[Bash Builtins](https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html,), and 
[Built-in csh and tcsh Commands](https://docstore.mik.ua/orelly/linux/lnut/ch08_09.htm).

Feel free to pick from `sh`, `bash`, `csh`, `tcsh`, `ksh` or `zsh` builtins... or if you have something else in mind, ping me and we'll work it out.

As an example, two shell builtins have already been added to the package builtins:

- `cd`
- `env`

## Steps

1. Clone down the example input/output and skeleton `main.go`:

    `git clone https://github.com/jh125486/CSCE4600`
 
2. Copy the `Project2` files to your own git project.

    1. In your `go.mod`, replace "jh125486" in the module line with your GitHub name, e.g.:

      - "module github.com/jh125486/CSCE4600" changes to "module github.com/CoolStudent123/CSCE4600"
  
    2. In the `main.go`, replace "jh125486" in the imports with your package path, e.g.:

      - "github.com/jh125486/CSCE4600/Project2/builtins" changes to "github.com/CoolStudent123/CSCE4600/Project2/builtins"

3. Start editing the `main.go` command switch (lines 57-64) and the package `builtins` with your chosen commands.

---------------------------------------------------------------------------------------------------------------------------------------------

The New Builtins added:
1. Echo: Outputs the given string or strings back to the terminal.
2. Clear: Clears the terminal screen, removing all previous output.
3. Date: Outputs the current date and time based on the system's locale settings.
4. History: Displays a list of previously executed commands in the session.
5. Pwd: Prints the current working directory path to the terminal.
6. Sleep: Pauses the execution of the program for a specified number of seconds.
7. Hostname: Shows the name of the system on which the environment is running.
8. Version: Displays the version of the application.
9. Help: Provides a list and descriptions of all available commands within the application.
10. Whoami: Displays the current system user's username.

   

