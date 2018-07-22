ejikit
====================

Generate a useful template for emoji

## Install

```bash
$ go get github.com/fukumone/ejikit
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

### Output oneline output emoji

```go
package main

import (
  "github.com/fukumone/ejikit"
)

func main() {
  ejikit.OnLineEmojiOutPut()
}
```
