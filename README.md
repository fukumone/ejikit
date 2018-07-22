ejiKit
====================

Generate a useful template for emoji

## Install

```bash
$ go get github.com/fukumone/generateEmojiCodeMap
```

## Usage

### Create a mapping for emoji

```go
package main

import (
  "github.com/fukumone/ejikit"
)

func main() {
  ejikit.CreateMap()
}
```
