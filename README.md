# Introduction

Go's idiomatic error handling is primarily an error return pattern. For certain scenarios, however, a try-catch-finally pattern is more appropriate, and sometimes completely necessary.

Go has the `panic()` and `recover()` functions and the `defer` keyword, which vaguely map to the `throw` and `try...catch...finally` concepts in many common languages.

Implementing this pattern correctly, and handling arbitrary panic types the way many are familiar with in other languages, ends up being a bit messy. So here is a function that captures that pattern in one place.

# API

The `Try` function accepts a body, a list of `Recover` blocks, and a `finally` block. The body is executed first, the `Recover` block which matches the panic type, if any, is executed next, and lastly the `finally` block is executed.

NOTE: The recover block matching is by iterating the blocks until the panic type is assignable to the type specified in the block.

# Usage

```go
package main

import (
    . "github.com/NickStrupat/go-trf"
)

func main() {
    Try(
        func() {
            println("attempt to do things here")

            panic("a string containing an error message")
        },
        Recovers{
            Recover(func(err string) {
                println("caught string: ", err)
            }),
            Recover(func(err int) {
                println("caught int:", err)
            }),
            Recover(func(err any) {
                println("caught any", err)
            }),
        },
        func() {
            println("finally")
        },
    )
}

```