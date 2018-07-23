ejikit
====================

Generate a useful template for [eji](https://github.com/fukumone/eji)

## Install

```bash
$ go get github.com/fukumone/ejikit
```

## Usage

### Create template for emoji

```go
package main

import (
  "github.com/fukumone/ejikit"
)

func main() {
  ejikit.CreateTemplate()
}
```

### Output emoji list

```go
package main

import (
  "github.com/fukumone/ejikit"
)

func main() {
  ejikit.EmojiOutput()
}
```

### Output oneline emoji

```go
package main

import (
  "github.com/fukumone/ejikit"
)

func main() {
  ejikit.OnLineEmojiOutPut()
}
```
