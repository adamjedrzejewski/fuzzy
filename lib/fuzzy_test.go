package fuzzy

import (
	"regexp"
	"testing"
)

func testFuzzyTimeEquals(t *testing.T, hour, minute int, expectedString string, expectedError error) {
	actual, err := SayTime(hour, minute)
	if expectedString != actual || err != expectedError {
		t.Fatalf("SayTime(%d, %d) = %q, %v, expected %#q, %#v", hour, minute, actual, err, expectedString, expectedError)
	}
}

func testFuzzyTimeMatches(t *testing.T, hour, minute int, expectedPattern string, expectedError error) {
	actual, err := SayTime(hour, minute)
	expectedRe := regexp.MustCompile(expectedPattern)
	if !expectedRe.MatchString(actual) || err != expectedError {
		t.Fatalf("SayTime(%d, %d) = %q, %v, expected match for %#q, %#v", hour, minute, actual, err, expectedPattern, expectedError)
	}
}

func TestFullHours(t *testing.T) {
	oClockHours := map[int]string{
		0:  "twelve o'clock",
		12: "twelve o'clock",
		1:  "one o'clock",
		13: "one o'clock",
		2:  "two o'clock",
		14: "two o'clock",
		3:  "three o'clock",
		15: "three o'clock",
		4:  "four o'clock",
		16: "four o'clock",
		5:  "five o'clock",
		17: "five o'clock",
		6:  "six o'clock",
		18: "six o'clock",
		7:  "seven o'clock",
		19: "seven o'clock",
		8:  "eight o'clock",
		20: "eight o'clock",
		9:  "nine o'clock",
		21: "nine o'clock",
		10: "ten o'clock",
		22: "ten o'clock",
		11: "eleven o'clock",
		23: "eleven o'clock",
	}

	for hour, expected := range oClockHours {
		testFuzzyTimeEquals(t, hour, 0, expected, nil)
	}
}

func TestOClock(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 58, "o'clock$", nil)
	testFuzzyTimeMatches(t, 0, 59, "o'clock$", nil)
	testFuzzyTimeMatches(t, 0, 0, "o'clock$", nil)
	testFuzzyTimeMatches(t, 0, 1, "o'clock$", nil)
	testFuzzyTimeMatches(t, 0, 2, "o'clock$", nil)
}

func TestFivePast(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 3, "^five past", nil)
	testFuzzyTimeMatches(t, 0, 4, "^five past", nil)
	testFuzzyTimeMatches(t, 0, 5, "^five past", nil)
	testFuzzyTimeMatches(t, 0, 6, "^five past", nil)
	testFuzzyTimeMatches(t, 0, 7, "^five past", nil)
}

func TestTenPast(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 8, "^ten past", nil)
	testFuzzyTimeMatches(t, 0, 9, "^ten past", nil)
	testFuzzyTimeMatches(t, 0, 10, "^ten past", nil)
	testFuzzyTimeMatches(t, 0, 11, "^ten past", nil)
	testFuzzyTimeMatches(t, 0, 12, "^ten past", nil)
}

func TestQuarterPast(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 13, "^quarter past", nil)
	testFuzzyTimeMatches(t, 0, 14, "^quarter past", nil)
	testFuzzyTimeMatches(t, 0, 15, "^quarter past", nil)
	testFuzzyTimeMatches(t, 0, 16, "^quarter past", nil)
	testFuzzyTimeMatches(t, 0, 17, "^quarter past", nil)
}

func TestTwentyPast(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 18, "^twenty past", nil)
	testFuzzyTimeMatches(t, 0, 19, "^twenty past", nil)
	testFuzzyTimeMatches(t, 0, 20, "^twenty past", nil)
	testFuzzyTimeMatches(t, 0, 21, "^twenty past", nil)
	testFuzzyTimeMatches(t, 0, 22, "^twenty past", nil)
}

func TestTwentyFivePast(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 23, "^twenty five past", nil)
	testFuzzyTimeMatches(t, 0, 24, "^twenty five past", nil)
	testFuzzyTimeMatches(t, 0, 25, "^twenty five past", nil)
	testFuzzyTimeMatches(t, 0, 26, "^twenty five past", nil)
	testFuzzyTimeMatches(t, 0, 27, "^twenty five past", nil)
}

func TestHalfPast(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 28, "^half past", nil)
	testFuzzyTimeMatches(t, 0, 29, "^half past", nil)
	testFuzzyTimeMatches(t, 0, 30, "^half past", nil)
	testFuzzyTimeMatches(t, 0, 31, "^half past", nil)
	testFuzzyTimeMatches(t, 0, 32, "^half past", nil)
}

func TestTwentyFiveTo(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 33, "^twenty five to", nil)
	testFuzzyTimeMatches(t, 0, 34, "^twenty five to", nil)
	testFuzzyTimeMatches(t, 0, 35, "^twenty five to", nil)
	testFuzzyTimeMatches(t, 0, 36, "^twenty five to", nil)
	testFuzzyTimeMatches(t, 0, 37, "^twenty five to", nil)
}

func TestTwentyTo(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 38, "^twenty to", nil)
	testFuzzyTimeMatches(t, 0, 39, "^twenty to", nil)
	testFuzzyTimeMatches(t, 0, 40, "^twenty to", nil)
	testFuzzyTimeMatches(t, 0, 41, "^twenty to", nil)
	testFuzzyTimeMatches(t, 0, 42, "^twenty to", nil)
}

func TestQuarterTo(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 43, "^quarter to", nil)
	testFuzzyTimeMatches(t, 0, 44, "^quarter to", nil)
	testFuzzyTimeMatches(t, 0, 45, "^quarter to", nil)
	testFuzzyTimeMatches(t, 0, 46, "^quarter to", nil)
	testFuzzyTimeMatches(t, 0, 47, "^quarter to", nil)
}

func TestTenTo(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 48, "^ten to", nil)
	testFuzzyTimeMatches(t, 0, 49, "^ten to", nil)
	testFuzzyTimeMatches(t, 0, 50, "^ten to", nil)
	testFuzzyTimeMatches(t, 0, 51, "^ten to", nil)
	testFuzzyTimeMatches(t, 0, 52, "^ten to", nil)
}

func TestFiveTo(t *testing.T) {
	testFuzzyTimeMatches(t, 0, 53, "^five to", nil)
	testFuzzyTimeMatches(t, 0, 54, "^five to", nil)
	testFuzzyTimeMatches(t, 0, 55, "^five to", nil)
	testFuzzyTimeMatches(t, 0, 56, "^five to", nil)
	testFuzzyTimeMatches(t, 0, 57, "^five to", nil)
}
