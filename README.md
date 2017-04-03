# cjam

A cli tool to create boilerplate code when participating in Google's Code Jam.

## Usage

```bash
$ go get github.com/fenrirunbound/cjam
$ cjam
NAME:
   CJam - A Google Code Jam boilerplate generator

USAGE:
   cjam [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     init, i  initialize the current folder for a new Code Jam problem set
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Golang (Go)

To create the boilerplate for a go-based solution, run the following command

```bash
$ cjam init go
```

### NodeJS 6.x

To create the boilerplate for a Node6-based solution, run the following command:

```bash
$ cjam init nodejs
```

### Python 2.7

To create the boilerplate for a python2.7-based solution, run the following command:

```bash
$ cjam init python
```
