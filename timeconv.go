package timeconv

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var Regex_StrToMs_all = regexp.MustCompile(`^\d+(?:\.\d+)?(h|m|ms|s)?$`)
var Regex_StrToMs_with_unit = regexp.MustCompile(`^\d+(?:\.\d+)?(h|m|ms|s)$`)

// Convert string to milliseconds
// 10s -> 10000
// 1m -> 60000
// 250ms -> 250
// 125 -> 125
// defvalue is returned if the string is blank (defvalue is in milliseconds)
func StrToDuration(val string, defvalue time.Duration) (time.Duration, error) {

	if val == "" {
		return time.Duration(defvalue), nil
	}

	val = strings.ReplaceAll(val, " ", "")
	val = strings.ToLower(val)

	if Regex_StrToMs_all.MatchString(val) {

		// if no unit is provided, assume milliseconds
		if !Regex_StrToMs_with_unit.MatchString(val) {
			val = val + "ms"
		}

		// Find breakingpoint between number and unit
		var pos int
		var c rune
		for pos, c = range val {
			if !(c >= '0' && c <= '9' || c == '.') {
				break
			}
		}

		// Get the numeric value
		num, err := strconv.ParseFloat(val[:pos], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid time value: %s", val)
		}

		// Get the unit
		unit := val[pos:]

		// Convert the value and unit to milliseconds
		switch unit {
		case "h":
			return time.Duration(num * float64(time.Hour)), nil
		case "m":
			return time.Duration(num * float64(time.Minute)), nil
		case "s":
			return time.Duration(num * float64(time.Second)), nil
		case "ms", "":
			return time.Duration(num * float64(time.Millisecond)), nil
		default:
			return 0, fmt.Errorf("invalid time unit: %s", unit)
		}
	}

	// Whatever we got was invalid
	return 0, fmt.Errorf("invalid time: %s", val)
}

// MustStrToDuration is like StrToDuration but panics if the string is invalid
func MustStrToDuration(val string, defvalue time.Duration) time.Duration {
	result, err := StrToDuration(val, defvalue)
	if err != nil {
		panic(err)
	}
	return result
}

// MustStrToDuration2 is like StrToDuration with the following charasteristics:
//   - Panics if the string is invalid
//   - No default needed, defaults to 100ms if invalid
func MustStrToDuration2(val string) time.Duration {
	result, err := StrToDuration(val, 100*time.Millisecond)
	if err != nil {
		panic(err)
	}
	return result
}
