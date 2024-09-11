## timeconv: A Go Package for Converting Strings to Time Durations

### Introduction

The `timeconv` package provides a convenient way to convert strings representing time durations into `time.Duration` values in Go. It supports various time units (hours, minutes, seconds, milliseconds) and handles common formatting variations.

### Usage

**1. Importing the Package:**

```bash
go get github.com/maeck70/timeconv
```

```go
import "github.com/maeck70/timeconv"
```

**2. Converting Strings to Durations:**

* `StrToDuration(val string, defvalue time.Duration) (time.Duration, error)`:
  - Converts the input string `val` into a `time.Duration` value.
  - Supports various time units (e.g., "10s", "1m", "250ms").
  - Returns a `time.Duration` value and an error if the conversion fails.
  - If the input string is empty, the `defvalue` is returned.

* `MustStrToDuration(val string, defvalue time.Duration) time.Duration`:
  - Similar to `StrToDuration` but panics if the conversion fails instead of returning an error.

### Examples

```go
package main

import (
    "fmt"
    "time"
    "github.com/maeck70/timeconv"
)

func main() {
    // Convert various string formats to durations
    duration1, err := timeconv.StrToDuration("10s", 0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(duration1) // Output: 10s
    }

    duration2 := timeconv.MustStrToDuration("1m", 0)
    fmt.Println(duration2) // Output: 1m

    // Use defvalue for empty strings
    duration3, _ := timeconv.StrToDuration("", time.Minute)
    fmt.Println(duration3) // Output: 1m
}
```

### Supported Formats

The package supports the following string formats:

- **Numeric values:**
  - `10` (assumed to be in milliseconds)
  - `1234` (assumed to be in milliseconds)
- **Numeric values with units:**
  - `10s` (seconds)
  - `1m` (minutes)
  - `250ms` (milliseconds)
  - `1h` (hours)

### Additional Notes

- The package handles whitespace and case insensitivity.
- If a unit is not provided, it is assumed to be milliseconds.
- The `MustStrToDuration` function is useful for situations where errors should be handled by panicking.
- The `MustStrToDuration2` function builds on MustStrToDuration and no default is needed (defaults to 100ms in case of error).

By using the `timeconv` package, you can easily and reliably convert strings representing time durations into `time.Duration` values in your Go applications.
