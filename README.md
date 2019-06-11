# namegen

namegen generates random names of things nobody would ever want.

[![Build Status](https://travis-ci.com/corysm1th/namegen.svg?branch=master)](https://travis-ci.com/corysm1th/namegen)

## Usage

### Command Example

```sh
$ namegen 5

blue_radioactive_ET_fingers
purple_firery_balzaks
green_pulverizing_magic_mitts
purple_radioactive_blue_blockers
green_electrocuting_pogo_sticks
```

### Code Example

```go
package main

import (
	"fmt"

	namegen "github.com/corysm1th/namegen/pkg"
)

func main() {
	textCase := namegen.Case(namegen.Snake)

	fmt.Println(namegen.Things(textCase))
}
```