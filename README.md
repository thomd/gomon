# Gomon

Gomon is a utility for monitoring changes in your Go source code and automatically restart your program.

This is particularly useful for development of Go servers.

## Install

    go get github.com/thomd/gomon

## Examples of usage

Run `go run .` as

    gomon .

Ignore folders with `-i` flag:

    gomon -i .git -i test .
