# Please, don't use this

It is really just a meme. Don't use this module in your projects.

# Introduction

This package is a joke inspired by `java.util.Optional` class.

# Examples

```go
package main

import "github.com/igoracmelo/optional"

func main() {
  opt := optional.Of(123)

  opt.Equals(123)                                  // true
  opt.IsPresent()                                  // true
  opt.IsEmpty()                                    // false

  opt.IfPresent(func () {                          // func gets called
    DoSomething()             
  })

  x := optional.OfNullable(nil).OrElse(10)         // x = 10
}
```
