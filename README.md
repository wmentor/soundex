# soundex

Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English. The goal is for homophones to be encoded to the same representation so that they can be matched despite minor differences in spelling.

This library is a go implementation of the Soundex algorithm.

# Summary

* Require Go version >= 1.14
* Support two algorithm version
* Work only with English words
* No external dependencies
* MIT license

# Install

```plaintext
go get github.com/wmentor/soundex
```

# Usage

```golang
package main

import (
  "fmt"

  "github.com/wmentor/soundex"
)

func main() {

  snd, _ := soundex.New()

  fmt.Println(snd.Code("ammonium"))       // A555
  fmt.Println(snd.Code("implementation")) // I514
  fmt.Println(snd.Code("Miller"))         // M460
  fmt.Println(snd.Code("Muller"))         // M460

  // improved algorithm version

  snd, _ = soundex.New(soundex.AlgoImproved)

  fmt.Println(snd.Code("morphs"))      // M913
  fmt.Println(snd.Code("traderworld")) // T969976
  fmt.Println(snd.Code("aquaman"))     // A588
}
```
