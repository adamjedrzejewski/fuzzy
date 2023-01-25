#include "fuzzy.h"

#include <time.h>
#include <stdio.h>

#define BUFFER_SIZE 256

int main(void) {
	char buffer[BUFFER_SIZE];

	time_t now = time(NULL);
	struct tm *now_tm = localtime(&now);
	int hour = now_tm->tm_hour;
	int minute = now_tm->tm_min;

	fputs(get_fuzzy_time(buffer, BUFFER_SIZE, hour, minute), stdout);
}
