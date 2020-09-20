# Gomon

Gomon is a utility for monitoring changes in your Go source code and automatically restart your program.

This is particularly useful for development of Go servers.

## Install

    go get github.com/thomd/gomon

## Usage

Instead of a regular

    go run server.go

run

    gomon server.go
