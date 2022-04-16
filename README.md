# github.com/pascal71/tstmod

## Usage

### Initialize your module

```
$ go mod init example.com/tstmod
```

### Get the recursivefunc module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/pascal71/tstmod@v0.1.0
```

```go
package main

import (
    "fmt"

    "github.com/pascal71/tstmod"
)

func main() {
    fmt.Println(recursivefunc.Fac(3))
}
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v0.1.0
$ git push origin --tags
```

