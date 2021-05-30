# go-study

Go Study Materials

## Installation

**Ubuntu**

```
$ sudo apt install golang-go
```

## Docker

```
$ docker pull golang[:latest]
```

```
$ docker run -t -d \
    -v `pwd`:/go-study \
    golang:latest
```

## `go run` command

Use `go run` command to compile and run a file or the named main Go package.
