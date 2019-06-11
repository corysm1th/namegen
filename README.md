# namegen

namegen generates random names of things nobody would ever want.

[![Build Status](https://travis-ci.org/corysm1th/namegen.svg?branch=travis)](https://travis-ci.org/corysm1th/namegen)
[![codecov](https://codecov.io/gh/corysm1th/namegen/branch/master/graph/badge.svg)](https://codecov.io/gh/corysm1th/namegen)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

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