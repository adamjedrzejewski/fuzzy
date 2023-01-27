package fuzzy

import "fmt"

func SayTime(hour, minute int) (string, error) {
	if hour < 0 || hour > 23 {
		return "", fmt.Errorf("Hour: %d out of bounds.", hour)
	}
	if minute < 0 || minute > 59 {
		return "", fmt.Errorf("Minute: %d out of bounds.", minute)
	}

	if minute < 3 {
		return fmt.Sprintf("%s o'clock", sayHour(hour)), nil
	}
	if minute < 8 {
		return fmt.Sprintf("five past %s", sayHour(hour)), nil
	}
	if minute < 13 {
		return fmt.Sprintf("ten past %s", sayHour(hour)), nil
	}
	if minute < 18 {
		return fmt.Sprintf("quarter past %s", sayHour(hour)), nil
	}
	if minute < 23 {
		return fmt.Sprintf("twenty past %s", sayHour(hour)), nil
	}
	if minute < 28 {
		return fmt.Sprintf("twenty five past %s", sayHour(hour)), nil
	}
	if minute < 33 {
		return fmt.Sprintf("half past past %s", sayHour(hour)), nil
	}
	if minute < 38 {
		return fmt.Sprintf("twenty five to %s", sayHour(hour+1)), nil
	}
	if minute < 43 {
		return fmt.Sprintf("twenty to %s", sayHour(hour+1)), nil
	}
	if minute < 48 {
		return fmt.Sprintf("quarter to %s", sayHour(hour+1)), nil
	}
	if minute < 53 {
		return fmt.Sprintf("ten to %s", sayHour(hour+1)), nil
	}
	if minute < 58 {
		return fmt.Sprintf("five to %s", sayHour(hour+1)), nil
	}

	return fmt.Sprintf("%s o'clock", sayHour(hour+1)), nil
}

func sayHour(hour int) string {
	switch hour {
	case 1:
		fallthrough
	case 13:
		return "one"
	case 2:
		fallthrough
	case 14:
		return "two"
	case 3:
		fallthrough
	case 15:
		return "three"
	case 4:
		fallthrough
	case 16:
		return "four"
	case 5:
		fallthrough
	case 17:
		return "five"
	case 6:
		fallthrough
	case 18:
		return "six"
	case 7:
		fallthrough
	case 19:
		return "seven"
	case 8:
		fallthrough
	case 20:
		return "eight"
	case 9:
		fallthrough
	case 21:
		return "nine"
	case 10:
		fallthrough
	case 22:
		return "ten"
	case 11:
		fallthrough
	case 23:
		return "eleven"
	case 0:
		fallthrough
	case 12:
		return "twelve"
	}

	return ""
}
