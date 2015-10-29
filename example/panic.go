package main

import (
    "github.com/gotool"
)

func main() {
    defer gotool.Crash()
    panic("test")
}
