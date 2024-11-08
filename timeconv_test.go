package timeconv

import (
	"log"
	"testing"
	"time"
)

func TestStrToMilisec(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
	}{
		{"10s", time.Duration(10 * time.Second)},
		{"1m", time.Duration(1 * time.Minute)},
		{"250ms", time.Duration(250 * time.Millisecond)},
		{"125", time.Duration(125 * time.Millisecond)},
		{"126 ", time.Duration(126 * time.Millisecond)},
		{" 127 ", time.Duration(127 * time.Millisecond)},
		{"128 s", time.Duration(128 * time.Second)},
		{"1.5s", time.Duration(1500 * time.Millisecond)},
		{"1.5h", time.Duration(90 * time.Minute)},
	}

	// Test the conversion from the examples above
	for _, test := range tests {
		result, err := StrToDuration(test.input, 99999)
		if err != nil {
			t.Errorf("Error converting '%s' to milliseconds: %v", test.input, err)
		}

		log.Printf("Converted %s, to %v", test.input, result)
		if result != test.expected {
			t.Errorf("Expected %v, got %d", test.expected, result)
		}
	}

	// Error on a string that cannot be converted
	_, err := StrToDuration("nonsense", 10000)
	if err == nil {
		t.Errorf("Expected error converting 'nonsense' to milliseconds")
	}

	// Test returning default value when the value is blank
	d := time.Duration(10000)
	v, err := StrToDuration("", d)
	if err == nil {
		if v != d {
			t.Errorf("While converting blank, expected %v, got %d", d, v)
		}
	} else {
		t.Errorf("Did not expect an error converting '' to milliseconds")
	}

}
func TestDurationToStr(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{time.Duration(0), "0"},
		{time.Duration(1 * time.Hour), "1h"},
		{time.Duration(2 * time.Hour), "2h"},
		{time.Duration(1 * time.Minute), "1m"},
		{time.Duration(2 * time.Minute), "2m"},
		{time.Duration(1 * time.Second), "1s"},
		{time.Duration(2 * time.Second), "2s"},
		{time.Duration(90 * time.Second), "90s"},
		{time.Duration(60 * time.Second), "1m"},
		{time.Duration(1 * time.Millisecond), "1ms"},
		{time.Duration(2 * time.Millisecond), "2ms"},
		{time.Duration(1500 * time.Millisecond), "1500ms"},
	}

	for _, test := range tests {
		result := DurationToStr(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
