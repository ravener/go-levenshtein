# Go Levenshtein
[The Levenshtein Edit Distance](https://en.wikipdia.org/wiki/Levenshtein_distance) algorithm in [Golang](https://golang.org)

This is a straight port of [fast-levenshtein](https://github.com/hiddentao/fast-levenshtein) ([npm](https://npmjs.com/package/fast-levenshtein)) from JavaScript.

- As efficient as the original by not storing the whole matrix.
- Built UTF-8 aware
- Small, the go file is about 1 KB
- Easy to use, exposes only one function.

## Install
```sh
$ go get github.com/ravener/go-levenshtein
```

## Usage
```go
package main

import (
  "fmt"
  "github.com/pollen5/levenshtein"
)

func main() {
  fmt.Println(levenshtein.Get("kitten", "sitting")) // => 3

  // UTF-8 aware
  fmt.Println(levenshtein.Get("Françe", "France")) // => 1
}
```

## License
[MIT](LICENSE)
