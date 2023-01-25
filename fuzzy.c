#include "fuzzy.h"

#include <stdio.h>
#include <stddef.h>

const char* get_hour_name(int hour);

char* get_fuzzy_time(char* buffer, size_t size, int hour, int minute) {
        const char *format_string = NULL;

	if (minute < 3) {
		format_string = "%s o'clock";
	} else if (minute < 8) {
		format_string = "five past %s";
	} else if (minute < 13) {
		format_string = "ten past %s";
	} else if (minute < 18) {
		format_string = "quater past %s";
	} else if (minute < 23) {
		format_string = "twenty past %s";
	} else if (minute < 28) {
		format_string = "twenty five past %s";
	} else if (minute < 33) {
		format_string = "half past %s";
	} else if (minute < 38) {
		format_string = "twenty five to %s";
		hour += 1;
	} else if (minute < 43) {
		format_string = "twenty to %s";
		hour += 1;
	} else if (minute < 48) {
		format_string = "quarter to %s";
		hour += 1;
	} else if (minute < 53) {
		format_string = "ten to %s";
		hour += 1;
	} else if (minute < 58) {
		format_string = "five to %s";
		hour += 1;
	} else {
		format_string = "%s o'clock";
		hour += 1;
	}

	if (format_string == NULL) {
		return NULL;
	}

	snprintf(buffer, size, format_string, get_hour_name(hour));
	return buffer;
}

const char* get_hour_name(int hour) {
	switch (hour) {
	case 0:
	case 12:
		return "twelve";
	case 1:
	case 13:
		return "one";
	case 2:
	case 14:
		return "two";
	case 3:
	case 15:
		return "three";
	case 4:
	case 16:
		return "four";
	case 5:
	case 17:
		return "five";
	case 6:
	case 18:
		return "six";
	case 7:
	case 19:
		return "seven";
	case 8:
	case 20:
		return "eight";
	case 9:
	case 21:
		return "nine";
	case 10:
	case 22:
		return "ten";
	case 11:
	case 23:
		return "eleven";
	default:
		return "";
	}
}
