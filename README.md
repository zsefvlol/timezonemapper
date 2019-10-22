# timezonemapper
Golang lib that converts from GPS coordinate latitude/longitude to timezone string. 

Translated from java version, see https://github.com/drtimcooper/LatLongToTimezone .

# Install

```shell script
go get -u github.com/zsefvlol/timezonemapper
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/zsefvlol/timezonemapper"
	"time"
)

func main() {
	// Get timezone string from lat/long
	timezone := timezonemapper.LatLngToTimezoneString(39.9254474,116.3870752)
	// Should print "Timezone: Asia/Shanghai"
	fmt.Printf("Timezone: %s\n", timezone)
	// Load location from timezone
	loc, _ := time.LoadLocation(timezone)
	// Parse time string with location
	t, _ :=time.ParseInLocation("2006-01-02 15:04:05", "2010-01-01 00:00:00", loc)
	// Should print
	// 2010-01-01 00:00:00 +0800 CST
	// 2009-12-31 16:00:00 +0000 UTC
	fmt.Println(t)
	fmt.Println(t.UTC())
}
```

# Map update

Last updated at `2019.10.22`

Will update (I guess) if the original project [drtimcooper/LatLongToTimezone](https://github.com/drtimcooper/LatLongToTimezone) updates. 
